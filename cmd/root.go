package cmd

import (
	"docker-cli/services"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var displayOptionCommand bool = false

var rootCmd = &cobra.Command{
	Use:   "ns",
	Short: "Un CLI pour exécuter et automatiser des commandes docker dans tous vos projets.",
	Long: `NS est un outil CLI conçu pour remplacer les fichiers Makefile et centraliser
l’exécution des commandes courantes docker sur tous vos projets.`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func RetrieveAllArgumentAfterTheCommand() []string {
	return os.Args[1:]
}

func init() {
    // desactive l'affichage auto des erreurs par cobra
    rootCmd.SilenceErrors = true 
    rootCmd.SilenceUsage = false 

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
        services.DisplayWithSpaceUpDown(func() {
            services.DisplayWithSpaceUpDown(func() {
                fmt.Println(services.CYAN, "🚀 Commandes disponibles :", services.RESET)
            })

            services.DisplayCommands(cmd)

            if displayOptionCommand {
                services.DisplayMessageForHelpCommand()
            }
    
            // Afficher les options spécifiques au bon contexte
            // On affiche toujours les options de `rootCmd` (ns)
            if len(args) == 0 {
                services.DisplayWithSpaceUpDown(func() {
                    fmt.Printf("ℹ️  %sOptions spécifiques à `ns` :%s\n", services.CYAN, services.RESET)
                })

                if cmd.Flags().HasFlags() {
                    services.DisplayFlagForCommand(cmd)
                } else {
                    fmt.Println("⚠️  Aucune option disponible.")
                }
            } else {
                // On vérifie si une sous-commande est demandée
                subCmd, _, _ := cmd.Root().Find(args)
                if subCmd != nil {
                    services.DisplayWithSpaceUpDown(func() {
                        fmt.Printf("%sℹ️  Options pour la commande `%s` :%s\n", services.CYAN, subCmd.Use, services.RESET)
                    })
    
                    if subCmd.Flags().HasFlags() || subCmd.PersistentFlags().HasFlags() {
                        services.DisplayFlagForCommand(subCmd)
                        services.DisplayFlagPersitForCommand(subCmd)
                    } else {
                        services.DisplayWithSpaceUpDown(func() {
                            fmt.Println("⚠️  Cette commande n'a pas d'options disponibles.")
                        })
                    }
                } else {
                    services.DisplayWithSpaceUpDown(func() {
                        fmt.Println("⚠️  Commande inconnue ou sans options.")
                    })
                }
            }
        })
    })
}

func Execute() {
	args := RetrieveAllArgumentAfterTheCommand()

	if len(args) == 0 || (len(args) == 1 && (args[0] == "--help" || args[0] == "-h")) {
        displayOptionCommand = true
		rootCmd.Help()
		return
	}

	// recupere la commande en cours
	subCmd, _, err := rootCmd.Find(args)

	if err != nil || subCmd == nil {
        services.DisplayWithSpaceUpDown(func() {
            fmt.Printf("%s❌ Erreur : Commande inconnue `%s`%s\n", services.RED, args[0], services.RESET)
            services.DisplayMessageForCommandHelp()
        })
		os.Exit(1)
	}
	
    // verifie que le premier argument est une demande d'aide
	if len(args) > 1 && (args[1] == "--help" || args[1] == "-h") {
        // on verifie que help ne fait pas partie des options de la commande
		if subCmd.Flags().Lookup("help") != nil || subCmd.PersistentFlags().Lookup("help") != nil {
            subCmd.Help()
        } else {
            fmt.Printf("%s❌ Erreur : L'option `%s` n'est pas reconnue pour `%s`. %s\n", services.RED, args[1], subCmd.Use, services.RESET)
            os.Exit(1)
        }
        return
	}

	// si une erreur persite on l'affiche
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%s❌ Erreur : %s%s\n", services.RED, err, services.RESET)
		os.Exit(1)
	}
}

// 📁 cmd/custom.go
package cmd

import (
	"docker-cli/internal/commandFile"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var customCmd = &cobra.Command{
	Use:   "custom [commande] [arg=value...]",
	Short: "Exécuter une commande définie dans commands.yaml",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ Merci de spécifier une commande personnalisée.")
			return
		}

		commandsFile, err := commandFile.LoadCommands("commands.yaml")
		if err != nil {
			log.Fatalf("Erreur chargement YAML: %v", err)
		}

		commandName := args[0]
		cmdConf, found := commandsFile.Commands[commandName]
		if !found {
			log.Fatalf("Commande '%s' introuvable", commandName)
		}

		requiredVars, err := commandFile.ExtractTemplateVars(cmdConf.Command)
		if err != nil {
			log.Fatalf("Erreur parsing template: %v", err)
		}

		providedArgs := make(map[string]string)
		for _, arg := range args[1:] {
			if strings.Contains(arg, "=") {
				parts := strings.SplitN(arg, "=", 2)
				key := parts[0]
				value := parts[1]
				providedArgs[key] = value
			}
		}

		if err := commandFile.ValidateParams(requiredVars, providedArgs); err != nil {
			log.Fatalf("Erreur: %v", err)
		}

		finalCmd, err := commandFile.RenderCommand(cmdConf.Command, providedArgs)
		if err != nil {
			log.Fatalf("Erreur génération commande: %v", err)
		}

		fmt.Printf("\n>> Exécution: %s\n\n", finalCmd)
		if err := commandFile.RunShellCommand(finalCmd); err != nil {
			log.Fatalf("Erreur exécution: %v", err)
		}
	},
}

var customListCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste les commandes disponibles depuis commands.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		commandsfile, err := commandFile.LoadCommands("commands.yaml")
		if err != nil {
			log.Fatalf("Erreur chargement YAML: %v", err)
		}

		fmt.Println("\n📦 Commandes disponibles :")
		for name, conf := range commandsfile.Commands {
			fmt.Printf("- %s : %s\n", name, conf.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(customCmd)
	customCmd.AddCommand(customListCmd)
}

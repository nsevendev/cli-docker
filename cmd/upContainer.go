package cmd

import (
	"docker-cli/internal/composeFile"
	"docker-cli/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

var nodetach bool
var envup string

func detachOrNot() string {
	if nodetach {
		return ""
	}

	return "-d"
}

var upContainerCmd = &cobra.Command{
	Use: "up",
	Short: "🐳 Lance les conteneurs (mode détaché par defaut)",
	Long: "En fonction des services dans le docker compose, la commande lance les conteneurs",
	Run: func(cmd *cobra.Command, args []string)  {
		var listCommands []string

		namesFile, err := composeFile.DetectAllComposeFile(envup)
		if err != nil {
			fmt.Printf("%s❌ Aucun fichier `compose` trouvé ! erreur : %v%s", services.RED, err, services.RESET)
		}

		fmt.Printf("%s🐳 Lecture du fichier %v%s\n", services.CYAN, namesFile, services.RESET)

		var fileStringToExecute string
		for i, nameFile := range namesFile {
			if i == len(nameFile)-1 {
				fileStringToExecute += fmt.Sprintf(" -f %v ", nameFile)
			} else {
				fileStringToExecute += fmt.Sprintf(" -f %v", nameFile)
			}
		}

		fmt.Printf("Les fichiers suivant vont etre executer pour monter les conteneurs : %v\n", fileStringToExecute)

		command := fmt.Sprintf("docker compose%v up %v", fileStringToExecute, detachOrNot())
		listCommands = append(listCommands, command)
		services.DisplayCommandsForExecute(&listCommands)

		errorService := services.ExecuteShellCommand(listCommands[0])
		if errorService != nil {
			fmt.Printf("%s❌ Erreur lors du démarrage des conteneurs : %s%s\n", services.RED, errorService, services.RESET)
			return
		}

		fmt.Println(services.GREEN + "✅ Conteneur(s) démarré(s) avec succès !" + services.RESET)
	},
}

func init() {
	upContainerCmd.Flags().BoolVarP(&nodetach, "nodetach", "n", false, "no detach (default detach)")
	upContainerCmd.Flags().StringVarP(&envup, "env", "e", "dev", "Environnement cible (`dev` ou `prod`, dev default)")

	rootCmd.AddCommand(upContainerCmd)
}
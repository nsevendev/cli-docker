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

		nameFile, err := composeFile.DetectComposeFile(envup)
		if err != nil {
			fmt.Printf("%s❌ Aucun fichier `compose` trouvé ! erreur : %v%s", services.RED, err, services.RESET)
		}

		fmt.Printf("%s🐳 Lecture du fichier %s%s\n", services.CYAN, nameFile, services.RESET)

		command := fmt.Sprintf("docker compose -f %v up %v", nameFile, detachOrNot())
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
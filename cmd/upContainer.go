package cmd

import (
	"docker-cli/internal/services"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var nodetach bool

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

		command := fmt.Sprintf("docker compose up %v", detachOrNot())
		listCommands = append(listCommands, command)
		services.DisplayCommandsForExecute(&listCommands)

		response := services.QuestionStartCommand("Voulez-vous lancer les conteneurs ?")
		if strings.ToLower(response) != "y" {
			services.DisplayWithSpaceUpDown(func() {
				fmt.Println(services.YELLOW + "🚫 Chargement des conteneurs annulé." + services.RESET)
			})
			return
		}

		services.ExecuteShellCommand(listCommands[0])
	},
}

func init() {
	upContainerCmd.Flags().BoolVarP(&nodetach, "nodetach", "n", false, "no detach (default detach)")

	rootCmd.AddCommand(upContainerCmd)
}
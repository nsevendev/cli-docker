package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var downContainerNameCmd = &cobra.Command{
	Use:   "d <container_name> ...",
	Short: "🐳 Arrête et supprime les conteneur désignés et donnés par nom à la commande",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Vérifie si Docker est installé
		if _, err := exec.LookPath("docker"); err != nil {
			fmt.Println("❌ Erreur : Docker n'est pas installé ou accessible dans le PATH.")
			os.Exit(1)
		}

		var listErrorMessage []string

		for i, a := range args {
			containerName := strings.TrimSpace(a)
			
			if containerName == "" {
				listErrorMessage = append(listErrorMessage, fmt.Sprintf("❌ Erreur : index %v vide aucun nom de conteneur", i))  
				continue
			}

			fmt.Printf("🔧 Suppression du conteneur \"%s\"...\n", containerName)

			dockerCmd := exec.Command("docker", "rm", "-f", containerName)
			dockerCmd.Stdout = os.Stdout
			dockerCmd.Stderr = os.Stderr

			if err := dockerCmd.Run(); err != nil {
				listErrorMessage = append(listErrorMessage, fmt.Sprintf("❌ Erreur : impossible de supprimer le conteneur \"%s\".\n", containerName))
				continue
			}

			fmt.Printf("✅ Conteneur \"%s\" supprimé avec succès.\n", containerName)
		}

		if len(listErrorMessage) > 0 {
			for _, messageError := range listErrorMessage {
				fmt.Println(messageError)
			}
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(downContainerNameCmd)
}
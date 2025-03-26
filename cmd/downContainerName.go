package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var downContainerNameCmd = &cobra.Command{
	Use:   "d <container_name>",
	Short: "Arrête et supprime le conteneur désigné et donné par nom à la commande",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerName := strings.TrimSpace(args[0])

		if containerName == "" {
			fmt.Println("❌ Erreur : le nom du conteneur ne peut pas être vide.")
			os.Exit(1)
		}

		// Vérifie si Docker est installé
		if _, err := exec.LookPath("docker"); err != nil {
			fmt.Println("❌ Erreur : Docker n'est pas installé ou accessible dans le PATH.")
			os.Exit(1)
		}

		fmt.Printf("🔧 Suppression du conteneur \"%s\"...\n", containerName)

		dockerCmd := exec.Command("docker", "rm", "-f", containerName)
		dockerCmd.Stdout = os.Stdout
		dockerCmd.Stderr = os.Stderr

		if err := dockerCmd.Run(); err != nil {
			fmt.Printf("❌ Erreur : impossible de supprimer le conteneur \"%s\".\n", containerName)
			os.Exit(1)
		}

		fmt.Printf("✅ Conteneur \"%s\" supprimé avec succès.\n", containerName)
	},
}

func init() {
	rootCmd.AddCommand(downContainerNameCmd)
}
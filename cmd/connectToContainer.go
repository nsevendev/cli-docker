package cmd

import (
    "fmt"
	"os" // imports basic
    "os/exec"
	"strings"

	"github.com/spf13/cobra" // imports externe
)

var connectToContainerCmd = &cobra.Command{
	Use:   "exec <container_name>", // command à utiliser
	Short: "📡 Connecte au conteneur.",
	Long: `Cette commande permet de se connecter à un conteneur.`,
    Args:  cobra.MinimumNArgs(1), // nombre d'arguments attendu (minimum 1 argument)
    Run: func(cmd *cobra.Command, args []string) {
		containerName := strings.TrimSpace(args[0]) // récupere le premier argument dans la commande et le stock dans la variable containerName

        // retourne une erreur si on donne + de 1 argument
        if len(args) > 1 {
            fmt.Println("❌ Erreur : trop d'arguments fournis. Argument attendu : 1 (nom du container).")
            os.Exit(1)
        }

		fmt.Printf("📡 Connexion au conteneur \"%s\" avec bash ...\n", containerName)

		dockerCmd := exec.Command("docker", "exec", "-it", containerName, "/bin/bash") // crée une command qui execute docker exec -it <containerName> "/bin/bash"
		dockerCmd.Stdout = os.Stdout // on envoie la sortie standard de la commande docker vers la sortie standard du programme go
		dockerCmd.Stderr = os.Stderr // pareil mais pour la sortie d'erreur
		dockerCmd.Stdin = os.Stdin // sortie pour le terminal

		if err := dockerCmd.Run(); err != nil {
			fmt.Println("🚨 échec de la connection avec `bash`, tentative avec `sh`...")

			dockerCmd = exec.Command("docker", "exec", "-it", containerName, "/bin/sh") // crée une command qui execute docker exec -it <containerName> "/bin/sh"
			dockerCmd.Stdout = os.Stdout  // on envoie la sortie standard de la commande docker vers la sortie standard du programme go
			dockerCmd.Stderr = os.Stderr  // pareil mais pour la sortie d'erreur
			dockerCmd.Stdin = os.Stdin // sortie pour le terminal

			if err := dockerCmd.Run(); err != nil {
				fmt.Printf("❌ Erreur : échec de la connection au container \"%s\" avec `bash` et `sh`.\n", containerName)
				os.Exit(1)
			}
		}
		
	},
}

func init() {
	rootCmd.AddCommand(connectToContainerCmd)
}
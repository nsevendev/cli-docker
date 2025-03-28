package cmd

import (
    "fmt"
	"os" // imports basic
    "os/exec"
	"strings"

	"github.com/spf13/cobra" // imports externe
)

var logsContainerCmd = &cobra.Command{
	Use:   "log <container_name>", // command à utiliser
	Short: "📜 Affiche les logs du conteneur.",
	Long: `Cette commande permet d'afficher les logs d'un conteneur.`,
    Args:  cobra.MinimumNArgs(1), // nombre d'arguments attendu (minimum 1 argument)
    Run: func(cmd *cobra.Command, args []string) {
		containerName := strings.TrimSpace(args[0]) // récupere le premier argument dans la commande et le stock dans la variable containerName

        // retourne une erreur si on donne + de 1 argument
        if len(args) > 1 {
            fmt.Println("❌ Erreur : trop d'arguments fournis. Argument attendu : 1 (nom du container).")
            os.Exit(1)
        }

		fmt.Printf("📜 Logs du conteneur \"%s\"...\n", containerName)

		dockerCmd := exec.Command("docker", "logs", "-f", containerName) // crée une command qui execute docker logs -f <containerName>
		dockerCmd.Stdout = os.Stdout // on envoie la sortie standard de la commande docker vers la sortie standard du programme go
		dockerCmd.Stderr = os.Stderr // pareil mais pour la sortie d'erreur

        if err := dockerCmd.Run(); err != nil { // execute la command que j'ai crée juste avant, et récupere les érreurs et les stock dans err si il y en a, puis rentre dans le if si err != de null
			fmt.Printf("❌ Erreur : impossible d'afficher les logs du conteneur \"%s\".\n", containerName)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(logsContainerCmd)
}
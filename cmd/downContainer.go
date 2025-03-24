package cmd

import (
	"docker-cli/internal/composeFile"
	"docker-cli/internal/services"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// executeShellCommand exécute une commande shell et redirige la sortie.
func executeShellCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var downContainerCmd = &cobra.Command{
	Use:   "down",
	Short: "🐳 Arrête et supprime les conteneurs Docker.",
	Long: `Cette commande tente d'arrêter et de supprimer les conteneurs 
créés par un fichier docker compose, ou par un Dockerfile si présent.`,
	Run: func(cmd *cobra.Command, args []string) {
		services.DisplayWithSpaceUpDown(func() {
			fmt.Println(services.CYAN + "🔽 Tentative d'arrêt et de suppression des conteneurs..." + services.RESET)
		})
		// 1) Vérifier la présence d'un fichier compose
		cf, err := composeFile.DetectComposeFile(env) // Appel via le package composeFile
		if err == nil {
			fmt.Println("Fichier compose trouvé :", cf)
			// On peut éventuellement afficher les conteneurs, mais on n'empêche pas l'exécution
			psCmd := exec.Command("sh", "-c", "docker compose ps -q")
			output, errPs := psCmd.Output()
			if errPs != nil {
				fmt.Printf("%s❌ Erreur lors de la vérification des conteneurs : %s%s\n", services.RED, errPs, services.RESET)
				return
			} else {
				containers := strings.TrimSpace(string(output))
				if containers == "" {
					fmt.Println(services.YELLOW + "⚠️ Aucun conteneur en cours d'exécution pour ce compose, mais on va quand même exécuter 'docker compose down'." + services.RESET)
				}
			}
			// Exécuter "docker compose down" même s'il n'y a aucun conteneur listé
			downCmdStr := "docker compose down"
			fmt.Printf("%s🚀 Exécution : %s%s\n", services.CYAN, downCmdStr, services.RESET)
			errDown := executeShellCommand(downCmdStr)
			if errDown != nil {
				fmt.Printf("%s❌ Erreur lors de docker compose down : %s%s\n", services.RED, errDown, services.RESET)
			} else {
				fmt.Println(services.GREEN + "✅ Conteneurs arrêtés et supprimés avec succès !" + services.RESET)
			}
		} else {
			// Pas de fichier compose => vérifier s'il existe un Dockerfile
			dockerfile := "Dockerfile"
			if _, errFile := os.Stat(dockerfile); errFile == nil {
				imageName := "mon-image:latest"
				// Lister tous les conteneurs (même arrêtés) pour cette image
				checkCmdStr := fmt.Sprintf("docker ps -a -q --filter ancestor=%s", imageName)
				out, errCheck := exec.Command("sh", "-c", checkCmdStr).Output()
				if errCheck != nil {
					fmt.Printf("%s❌ Erreur lors de la vérification des conteneurs : %s%s\n", services.RED, errCheck, services.RESET)
					return
				}
				containerIDs := strings.TrimSpace(string(out))
				if containerIDs == "" {
					fmt.Println(services.YELLOW + "⚠️ Aucun conteneur trouvé pour l'image " + imageName + services.RESET)
					return
				}
				stopRmCmdStr := fmt.Sprintf("docker stop %s && docker rm %s", containerIDs, containerIDs)
				fmt.Printf("%s🚀 Exécution : %s%s\n", services.CYAN, stopRmCmdStr, services.RESET)
				errStopRm := executeShellCommand(stopRmCmdStr)
				if errStopRm != nil {
					fmt.Printf("%s❌ Erreur lors de l'arrêt/suppression : %s%s\n", services.RED, errStopRm, services.RESET)
				} else {
					fmt.Println(services.GREEN + "✅ Conteneur(s) arrêté(s) et supprimé(s) avec succès !" + services.RESET)
				}
			} else {
				fmt.Println(services.RED + "❌ Aucun fichier compose ou Dockerfile trouvé !" + services.RESET)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(downContainerCmd)
}

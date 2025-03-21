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

// Options de la commande
var tag string
var dockerfile string
var env string

// Génération du nom de l'image si `image:` est absent
func generateImageName(serviceName string, env string) string {
	tag := "latest"
	
	if env == "prod" {
		tag = "prod"
	}

	return fmt.Sprintf("%s:%s", serviceName, tag)
}

// Fonction pour générer la commande `docker build`
func generateBuildCommands(composeService map[string]composeFile.ComposeService, env string) []string {
	var commands []string

	for name, service := range composeService {
		if service.Build.Context == "" {
			service.Build.Context = "." // Par défaut, build dans le dossier actuel
		}

		// Définition du nom de l’image
		imageName := service.Image

		if imageName == "" {
			imageName = generateImageName(name, env)
		}

		// Construction de la commande
		cmd := fmt.Sprintf("docker build -t %s", imageName)
		
		if service.Build.Target != "" {
			cmd += fmt.Sprintf(" --target %s", service.Build.Target)
		}
		
		if service.Build.Dockerfile != "" {
			cmd += fmt.Sprintf(" -f %s", service.Build.Dockerfile)
		}
		
		cmd += fmt.Sprintf(" %s", service.Build.Context)
		commands = append(commands, cmd)
	}

	if len(commands) == 0 {
		fmt.Println(services.RED + "❌ Aucun service à builder trouvé dans `compose.yml` !" + services.RESET)
		os.Exit(1)
	}

	return commands
}

// Fonction utilitaire pour exécuter une commande shell
func executeShellCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Exécution des builds
func executeBuild(commands []string) {
	for _, cmd := range commands {
		fmt.Printf("%s🚀 Exécution : %s%s\n", services.CYAN, cmd, services.RESET)
		err := executeShellCommand(cmd)
		
		if err != nil {
			fmt.Printf("%s❌ Erreur lors du build : %s%s\n", services.RED, err, services.RESET)
		} else {
			fmt.Printf("%s✅ Build terminé avec succès !%s\n", services.GREEN, services.RESET)
		}
	}
}

func questionStartCommand() (string) {
	fmt.Print("\n✅ Démarrer le build ? (y/N) : ")
	var response string
	fmt.Scanln(&response)
	return response
}

func displayCommandsForBuild(commands *[]string) {
	fmt.Println(services.CYAN + "📌 Commandes à exécuter :" + services.RESET)
	for _, cmd := range *commands {
		fmt.Println("  " + cmd)
	}
}

// buildImageCmd represents the buildImage command
var buildImageCmd = &cobra.Command{
	Use:   "bi",
	Short: "🐳 Construit Les images Docker pour le projet.",
	Long: `🚀 Cette commande permet de générer une image Docker à partir du Dockerfile du projet.`,
	Run: func(cmd *cobra.Command, args []string) {
		services.DisplayWithSpaceUpDown(func() {
			fmt.Println(services.CYAN + "🐳 Détection des images à builder..." + services.RESET)

			file, err := composeFile.DetectComposeFile(env)

			if err != nil {
				fmt.Println(services.RED + err.Error() + services.RESET)

				// Si aucun fichier compose, vérifier si un Dockerfile est spécifié
				if dockerfile == "" {
					dockerfile = "Dockerfile"
				}

				_, err := os.Stat(dockerfile)

				if os.IsNotExist(err) {
					fmt.Println(services.RED + "❌ Aucun `compose.yml` ni `Dockerfile` trouvé !" + services.RESET)
					os.Exit(1)
				}

				// Build direct avec Dockerfile seul
				imageName := "mon-image:latest"
				if tag != "" {
					imageName = tag
				}

				buildCmd := fmt.Sprintf("docker build -t %s -f %s .", imageName, dockerfile)
				fmt.Printf("%s🚀 Commande exécutée : %s%s\n", services.CYAN, buildCmd, services.RESET)

				// Confirmation
				response := questionStartCommand()
				if strings.ToLower(response) != "y" {
					fmt.Println(services.YELLOW + "🚫 Build annulé." + services.RESET)
					return
				}
				
				executeShellCommand(buildCmd)
				
				return
			}

			composeService := composeFile.ReadAndParseComposeFile(file)
			commands := generateBuildCommands(composeService, env)
			displayCommandsForBuild(&commands)

			// Confirmation
			response := questionStartCommand()
			if strings.ToLower(response) != "y" {
				fmt.Println(services.YELLOW + "🚫 Build annulé." + services.RESET)
				return
			}

			executeBuild(commands)
		})
	},
}

func init() {
	buildImageCmd.Flags().StringVarP(&tag, "tag", "t", "", "Nom et tag de l’image (ex: mon-image:v1.0)")
	buildImageCmd.Flags().StringVarP(&dockerfile, "file", "f", "", "Chemin du fichier Dockerfile (optionnel)")
	buildImageCmd.Flags().StringVarP(&env, "env", "e", "dev", "Environnement cible (`dev` ou `prod`)")

	rootCmd.AddCommand(buildImageCmd)
}

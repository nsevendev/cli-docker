package cmd

import (
	"docker-cli/internal/commandFile"
	"docker-cli/internal/services"
	"fmt"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func replaceNSVars(cmd string, env map[string]string) string {
	for key, value := range env {
		placeholder := "*" + key + "*"
		if strings.Contains(cmd, placeholder) {
			log.Printf("%s %v %s====>%s %v %s\n", services.CYAN, placeholder, services.RESET, services.YELLOW, value, services.RESET)
			cmd = strings.ReplaceAll(cmd, placeholder, value)
		}
	}

	return cmd
}

func displayVarsEnv(env map[string]string) {
	if len(env) > 0 {
		log.Printf("%sℹ️  Affichage des variables d'environement NS :%s\n", services.CYAN, services.RESET)
		found := false
		
		for key, value := range env {
			if strings.HasPrefix(key, "NSC_") {
				log.Printf("%s %v%s => %s%v%s\n",services.YELLOW, key, services.RESET, services.GREEN, value, services.RESET)
				found = true
			}
		}

		if !found {
			log.Println(services.YELLOW + "⚠️  Aucune variable d'environement commençant par 'NS_'" + services.RESET)
		}

		return
	}

	log.Printf("%s⚠️  Aucune variable d'environement:%s\n", services.YELLOW, services.RESET)
}

var customCmd = &cobra.Command{
	Use:   "c [commande] [arg=value...]",
	Short: "Exécuter une commande définie dans commands.yaml",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ Merci de spécifier une commande personnalisée.")
			return
		}

		log.Println(services.CYAN + "ℹ️  Lecture du fichier .env" + services.RESET)
		envVars, err := godotenv.Read(".env")
		if err != nil {
			log.Printf("%s⚠️  Aucun fichier .env trouvé ou erreur de lecture: %v%s\n", services.YELLOW, err, services.RESET)
			envVars = make(map[string]string)
		}
		displayVarsEnv(envVars)

		log.Println(services.CYAN + "ℹ️  Lecture du fichier commands.yaml" + services.RESET)
		commandsFile, err := commandFile.LoadCommands("commands.yaml")
		if err != nil {
			log.Fatalf("❌ Erreur chargement YAML: %v", err)
		}

		commandName := args[0]
		cmdConf, found := commandsFile.Commands[commandName]
		if !found {
			log.Fatalf("❌ Commande '%s' introuvable", commandName)
		}

		log.Println(services.CYAN + "ℹ️  Exctraction des variables de template" + services.RESET)
		requiredVars, err := commandFile.ExtractTemplateVars(cmdConf.Command)
		if err != nil {
			log.Fatalf("❌ Erreur parsing template: %v", err)
		}

		// args=value parse
		log.Println(services.CYAN + "ℹ️  Parcing des variables de template" + services.RESET)
		providedArgs := make(map[string]string)
		for _, arg := range args[1:] {
			if strings.Contains(arg, "=") {
				parts := strings.SplitN(arg, "=", 2)
				key := parts[0]
				value := parts[1]
				providedArgs[key] = value
			}
		}
		log.Printf("%s✅ Parsing des variables de template : %v%s\n", services.GREEN, providedArgs, services.RESET)

		if err := commandFile.ValidateParams(requiredVars, providedArgs); err != nil {
			log.Fatalf("❌ Erreur validation : %v", err)
		}
		log.Println(services.GREEN + "✅ Les args ont été validés" + services.RESET)

		finalCmd, err := commandFile.RenderCommand(cmdConf.Command, providedArgs)
		if err != nil {
			log.Fatalf("❌ Erreur génération commande: %v", err)
		}
		log.Printf(services.GREEN + "✅ Intégration de arguments dans la commande : %s%v\n", finalCmd, services.RESET)

		finalCmd = replaceNSVars(finalCmd, envVars)
		log.Printf(services.GREEN + "✅ Remplacement des variables d'environement dans la commande : %s%v\n", finalCmd, services.RESET)

		fmt.Printf("\n>> Exécution: %s\n\n", finalCmd)
		if err := commandFile.RunShellCommand(finalCmd); err != nil {
			log.Fatalf("❌ Erreur exécution: %v", err)
		}
	},
}

var customListCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste les commandes disponibles depuis commands.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		commandsfile, err := commandFile.LoadCommands("commands.yaml")
		if err != nil {
			log.Fatalf("❌ Erreur chargement YAML: %v", err)
		}

		fmt.Println("\n📦 Commandes disponibles :")
		for name, conf := range commandsfile.Commands {
			fmt.Printf("- %s%-25s%s : %s%s%s\n",services.GREEN, name, services.RESET, services.YELLOW, conf.Description, services.RESET)
		}
	},
}

func init() {
	rootCmd.AddCommand(customCmd)
	customCmd.AddCommand(customListCmd)
}

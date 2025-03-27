package commandFile

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// Charge le fichier YAML
func LoadCommands(filename string) (*CommandsFile, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cmdFile CommandsFile
	if err := yaml.Unmarshal(data, &cmdFile); err != nil {
		return nil, err
	}
	return &cmdFile, nil
}

// Extrait les variables {{var}} de la commande
func ExtractTemplateVars(templateStr string) ([]string, error) {
	re := regexp.MustCompile(`{{\s*([a-zA-Z0-9_]+)\s*}}`)
	matches := re.FindAllStringSubmatch(templateStr, -1)
	if matches == nil {
		return nil, errors.New("❌ aucune variable détectée dans la commande")
	}
	var vars []string
	for _, match := range matches {
		vars = append(vars, match[1])
	}
	return vars, nil
}

// Vérifie que tous les paramètres sont présents
func ValidateParams(required []string, given map[string]string) error {
	var missing []string
	for _, key := range required {
		if _, ok := given[key]; !ok {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("❌ paramètres manquants: %s", strings.Join(missing, ", "))
	}
	return nil
}

// Fait un remplacement basique {{var}} -> valeur
func RenderCommand(templateStr string, args map[string]string) (string, error) {
	log.Println("ℹ️  Template de la commande :\n", templateStr)
	log.Println("ℹ️  Liste arguments :\n", args)

	out := templateStr
	for key, value := range args {
		placeholder := fmt.Sprintf("{{%s}}", key)
		out = strings.ReplaceAll(out, placeholder, value)
		// aussi gérer {{ var }} avec espace
		placeholderSpaced := fmt.Sprintf("{{ %s }}", key)
		out = strings.ReplaceAll(out, placeholderSpaced, value)
	}
	// Vérifie qu'il ne reste pas de {{var}}
	leftover, _ := ExtractTemplateVars(out)
	if len(leftover) > 0 {
		return "", fmt.Errorf("❌ valeurs manquantes pour: %s", strings.Join(leftover, ", "))
	}
	return out, nil
}

// Exécute la commande shell
func RunShellCommand(cmdStr string) error {
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

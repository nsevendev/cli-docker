package composeFile

import (
	"docker-cli/internal/services"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func DetectComposeFile(env string) (string, error) {
	files := []string{
		"compose.yml", 
		"docker-compose.yml", 
		"compose.yaml", 
		"docker-compose.yaml",
	}

	if env == "prod" {
		files = []string{
			"compose.prod.yml", 
			"docker-compose.prod.yml", 
			"compose.prod.yaml", 
			"docker-compose.prod.yaml",
		}
	}

	for _, file := range files {
		_, err := os.Stat(file)
		
		if err == nil {
			return file, nil
		}
	}

	return "", fmt.Errorf("❌ Aucun fichier `compose` trouvé")
}

func ReadComposeFile(filePath *string) ([]byte) {
	data, errorFile := os.ReadFile(*filePath)
	
	if errorFile != nil {
		fmt.Println(services.RED + "❌ Erreur lors de la lecture du fichier compose :" + services.RESET)
		fmt.Println(errorFile)
		os.Exit(1)
	}

	return data
}

func ParseComposeYml(data []byte, composeData *ComposeFile) {
	err := yaml.Unmarshal(data, &composeData)
	if err != nil {
		fmt.Println(services.RED + "❌ Erreur lors de la lecture du fichier compose :" + services.RESET)
		fmt.Println(err)
		os.Exit(1)
	}
}

func ReadAndParseComposeFile(filePath string) (map[string]ComposeService) {
	data := ReadComposeFile(&filePath)
	var composeData ComposeFile   
	ParseComposeYml(data, &composeData)
	return composeData.Services
}
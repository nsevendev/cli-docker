package composeFile

import (
	"docker-cli/internal/services"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func nameFilesCompose(env string) []string {
	files := []string{
		"compose.yml", 
		"compose.override.yml", 
		"docker-compose.yml", 
		"docker-compose.override.yml", 
		"compose.yaml", 
		"compose.override.yaml", 
		"docker-compose.yaml",
		"docker-compose.override.yaml",
	}

	if env == "prod" {
		files = []string{
			"compose.yml", 
			"docker-compose.yml", 
			"compose.prod.yml", 
			"docker-compose.prod.yml", 
			"compose.yaml", 
			"docker-compose.yaml", 
			"compose.prod.yaml", 
			"docker-compose.prod.yaml",
		}
	}

	return files
}

func DetectComposeFile(env string) (string, error) {
	files := nameFilesCompose(env)

	for _, file := range files {
		_, err := os.Stat(file)
		
		if err == nil {
			return file, nil
		}
	}

	return "", fmt.Errorf("❌ Aucun fichier `compose` trouvé")
}

func DetectAllComposeFile(env string) ([]string, error) {
	files := nameFilesCompose(env)
	var found []string

	for _, file := range files {
		_, err := os.Stat(file)
		
		if err == nil {
			found = append(found, file)
		}
	}

	if len(found) == 0 {
		return nil, fmt.Errorf("❌ 0 fichier `compose` trouvé")
	}

	return found, nil
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
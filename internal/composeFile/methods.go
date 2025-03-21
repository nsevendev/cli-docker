package composeFile

import (
	"fmt"
	"os"
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
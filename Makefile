DOCKER_COMP = docker compose

# Misc
.DEFAULT_GOAL = help
.PHONY        : help build build-md build-mm build-l

## —— 🐳 Makefile project CLI-DOCKER 🐳 ——————————————————————————————————

help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9\./_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

build: ## build docker for this project
	@echo "🚀 build de image -------------> START"
	@$(DOCKER_COMP) build --pull --no-cache
	@echo "✅ build de l'image dev -------------> END"

dev: ## Start the docker mode dev (no logs)
	@echo "🚀 Demarrage des conteneurs dev -------------> START"
	@$(DOCKER_COMP) up --detach
	@echo "✅ Demarrage des conteneurs dev -------------> END"

stop: ## Stop the docker
	@echo "🚀 Arret des conteneurs -------------> START"
	@$(DOCKER_COMP) down --remove-orphans
	@echo "✅ Arret des conteneurs -------------> END"

## —— 🐳 Build 🐳 ——————————————————————————————————

ns: ## build binaire for this project 
	@echo "🚀 start buid -------------> START"
	go build -o ns

ns-md: ## build binaire for mac amd64
	GOOS=darwin GOARCH=amd64 go build -o ./build-mac-amd/ns

ns-mm: ## build binaire for mac arm64
	GOOS=darwin GOARCH=arm64 go build -o ./build-mac-arm/ns

ns-l: ## build binaire for linux
	GOOS=linux GOARCH=amd64 go build -o ./build-linux/ns
DOCKER_COMP = docker compose
VERSION = v0.0.1
APP_NAME = ns
FOLDER_BUILD_LOCAL = build-local

BUILD_MAC_AMD = GOOS=darwin GOARCH=amd64
BUILD_MAC_ARM = GOOS=darwin GOARCH=arm64
BUILD_LINUX = GOOS=linux GOARCH=amd64

COMMAND_BUILD_GO = go build -buildvcs=false -o

# Misc
.DEFAULT_GOAL = help
.PHONY        : help build build-md build-mm build-l

%:
	@:

## —— 🐳 Commande pour le container CLI-DOCKER 🐳 ——————————————————————————————————

help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9\./_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

build: ## build image docker for this project
	@echo "🚀 build de image -------------> START"
	@$(DOCKER_COMP) build --pull --no-cache
	@echo "✅ build de l'image dev -------------> END"

u: ## Start the container docker mode dev (no logs)
	@echo "🚀 Demarrage des conteneurs dev -------------> START"
	@$(DOCKER_COMP) up --detach
	@echo "✅ Demarrage des conteneurs dev -------------> END"

d: ## Stop the docker
	@echo "🚀 Arret des conteneurs -------------> START"
	@$(DOCKER_COMP) down --remove-orphans
	@echo "✅ Arret des conteneurs -------------> END"

## —— 🐳 Build + Install en local du CLI 🐳 ——————————————————————————————————

cc: # Clean le dossier build local
	@$(eval c ?=)
	@echo "🚀 clean -------------> START"
	rm -rf $(FOLDER_BUILD_LOCAL)/$(APP_NAME)
	@echo "✅ clean -------------> END"

bl: # build executable et l'ajoute en local au bin
	@$(eval c ?=)
	@$(MAKE) cc c=$(c)
	@echo "🚀 buid ns -------------> START"
	docker exec -i cli-docker sh -c "$(if $(filter $(c),d),$(BUILD_MAC_AMD),$(if $(filter $(c),m),$(BUILD_MAC_ARM),$(if $(filter $(c),l),$(BUILD_LINUX)))) $(COMMAND_BUILD_GO) $(FOLDER_BUILD_LOCAL)/$(APP_NAME)"
	@echo "✅ buid ns -------------> END"

il: ## 🖥️ Installer le binaire localement dans /usr/local/bin (c="l" pour linux, "d" pour macOsAmd, "m" pour macOsArm)
	@$(eval c ?=d)
	@$(MAKE) bl c=$(c)
	@echo "🚀 install ns -------------> START"
	sudo chmod +x $(FOLDER_BUILD_LOCAL)/$(APP_NAME)
	sudo mv $(FOLDER_BUILD_LOCAL)/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	@echo "✅ install ns -------------> END"

cli: ## 🚀 execute cli ns dev local
	@echo "🚀 exec cli -------------> START"
	docker exec -i cli-docker sh -c "tmp/ns $(wordlist 2, 99, $(MAKECMDGOALS))"
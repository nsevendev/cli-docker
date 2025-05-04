DOCKER_COMP = docker compose
COMPOSE_FILES := -f docker/compose.yaml

VERSION = v0.0.1
APP_NAME = ns
FOLDER_BUILD_LOCAL = build-local

BUILD_MAC_AMD = GOOS=darwin GOARCH=amd64
BUILD_MAC_ARM = GOOS=darwin GOARCH=arm64
BUILD_LINUX = GOOS=linux GOARCH=amd64

COMMAND_BUILD_GO = go build -buildvcs=false -o

# Misc
.PHONY        : help build build-md build-mm build-l
.DEFAULT_GOAL = help

## —— 🐳 ALL 🐳 ——————————————————————————————————
help: ## Afficher l'aide
	@grep -E '(^[a-zA-Z0-9\./_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

install: ## Instruction pour installer le projet
	@cat doc/install.md

install-ns: ## instruction pour installer le binaire localement
	@cat doc/install-build-ns.md

## —— 🐳 CONTAINER 🐳 ——————————————————————————————————

build: ## build image docker
	@echo "🚀 build de image -------------> START"
	@$(DOCKER_COMP) build --pull --no-cache
	@echo "✅ build de l'image dev -------------> END"

up: ## Démarre l'environnement de développement
	@echo "🚀 Demarrage des conteneurs dev -------------> START"
	@$(DOCKER_COMP) $(COMPOSE_FILES) up --detach
	@echo "✅ Demarrage des conteneurs dev -------------> END"

down: ## Arrête les conteneurs de développement
	@echo "🚀 Arret des conteneurs -------------> START"
	@$(DOCKER_COMP) $(COMPOSE_FILES) down --remove-orphans
	@echo "✅ Arret des conteneurs -------------> END"

## —— 🐳 TOOL 🐳 ——————————————————————————————————

s: ## Ouvre un shell dans le conteneur app
	@echo "🚀 Ouvrir un shell dans le conteneur -------------> START"
	@docker exec -it cli-docker sh
	@echo "✅ Ouvrir un shell dans le conteneur -------------> END"

l: ## Affiche les logs du conteneur app
	@echo "🚀 Affiche logs du conteneur -------------> START"
	@docker logs -f cli-docker

## —— 🐳 NS INSTALL LOCAL 🐳 ——————————————————————————————————

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

install-local: ## Installer le binaire localement (c="l" linux, "d" (defaut) macOsAmd, "m" macOsArm)
	@$(eval c ?=d)
	@$(MAKE) bl c=$(c)
	@echo "🚀 install ns -------------> START"
	sudo chmod +x $(FOLDER_BUILD_LOCAL)/$(APP_NAME)
	sudo mv $(FOLDER_BUILD_LOCAL)/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	@echo "✅ install ns -------------> END"

## —— 🐳 NS BUILD/EXEC LOCAL 🐳 ——————————————————————————————————

ns-exec: ## execute cli ns dev local (c="NAME COMMAND ns)
	@echo "🚀 exec cli -------------> START"
	docker exec -i cli-docker sh -c "tmp/ns $(wordlist 2, 99, $(MAKECMDGOALS))"
	@echo "🚀 exec cli -------------> END"

ns-build-mv: ## build et deplace le binaire dans le dossier cible (c="PATH cible")
	@$(eval c ?=)
	@echo "🛠️  Build du binaire Linux dans le conteneur..."
	@docker exec -i cli-docker sh -c "$(BUILD_LINUX) $(COMMAND_BUILD_GO) $(FOLDER_BUILD_LOCAL)/$(APP_NAME)"
	@echo "📁 Déplacement du binaire vers : $(c)"
	@mv $(FOLDER_BUILD_LOCAL)/$(APP_NAME) $(c)/$(APP_NAME)
	@chmod +x $(c)/$(APP_NAME)
	@echo "✅ Binaire copié dans $(c)/$(APP_NAME)"
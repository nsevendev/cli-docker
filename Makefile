DOCKER_COMP = docker compose
VERSION = v0.0.1
APP_NAME = ns
FOLDER_BUILD_LOCAL = build-local

BUILD_MAC_AMD = GOOS=darwin GOARCH=amd64
BUILD_MAC_ARM = GOOS=darwin GOARCH=arm64
BUILD_LINUX = GOOS=linux GOARCH=amd64

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

up: ## Start the docker mode dev (no logs)
	@echo "🚀 Demarrage des conteneurs dev -------------> START"
	@$(DOCKER_COMP) up --detach
	@echo "✅ Demarrage des conteneurs dev -------------> END"

down: ## Stop the docker
	@echo "🚀 Arret des conteneurs -------------> START"
	@$(DOCKER_COMP) down --remove-orphans
	@echo "✅ Arret des conteneurs -------------> END"

## —— 🐳 Build + Install en local 🐳 ——————————————————————————————————

clean:
	@$(eval c ?=)
	@echo "🚀 clean -------------> START"
	rm -rf $(FOLDER_BUILD_LOCAL)/$(APP_NAME)
	@echo "✅ clean -------------> END"

ns: 
	@$(eval c ?=)
	@$(MAKE) clean c=$(c)
	@echo "🚀 buid ns -------------> START"
	docker exec -i cli-docker sh -c "$(if $(filter $(c),d),$(BUILD_MAC_AMD),$(if $(filter $(c),m),$(BUILD_MAC_ARM),$(if $(filter $(c),l),$(BUILD_LINUX)))) go build -buildvcs=false -o $(FOLDER_BUILD_LOCAL)/$(APP_NAME)"
	@echo "✅ buid ns -------------> END"

il: ## 🖥️ Installer le binaire localement dans /usr/local/bin (c= l pour linux, d pour macOsAmd, m pour macOsArm)
	@$(eval c ?=d)
	@$(MAKE) ns c=$(c)
	@echo "🚀 install ns -------------> START"
	chmod +x $(FOLDER_BUILD_LOCAL)/$(APP_NAME)
	sudo mv $(FOLDER_BUILD_LOCAL)/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	@echo "✅ install ns -------------> END"

ns-md: ## build binaire for mac amd64
	GOOS=darwin GOARCH=amd64 go build -o ./build-mac-amd/ns

ns-mm: ## build binaire for mac arm64
	GOOS=darwin GOARCH=arm64 go build -o ./build-mac-arm/ns

ns-l: ## build binaire for linux
	GOOS=linux GOARCH=amd64 go build -o ./build-linux/ns

release-tag:
	git tag $(VERSION)
	git push origin $(VERSION)
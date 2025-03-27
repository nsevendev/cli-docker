# cli-docker

*Note : taper `make` pour voir les commande makefile disponible*

## Lancer le projet en mode dev

```bash
# build l'image docker du projet
make build

# lance le conteneur docker
make u

# pour arrêter le conteneur
make d
```

*Note : les commande `u`et `d` représentent les commandes `up` et `down`  
elles on été modifiées car elle rentre en conflit pour les tests de commande `ns up`et `ns down`*

## Utilisation

- une fois les containers demarrés, ouvrir vscode dans le container  
go est disponible dans le terminal, vous pouvez utiliser tout les commandes go que vous voulez  

- le Makefile, sert à 
    - dockerizer ce projet et faire tourner go pour le developpement
    - lancer les build du cli pour toute platforme ou le build en local du projet pour tester
tous le reste se fera avec les commandes go classique

## Utiliser le cli

une fois le repo cloner, lancer le conteneur en mode dev et lancer la commande suivante qui vous corresponds:  

```bash
# pour macOs amd64
make il

# pour macOs arm64
make il c="m"

# pour linux amd64
make il c="l"
```

*Note : il vous demandera votre mot de passe pour installer l'executable `ns`*  

- si tout c'est bien passé utiliser la commande `ns` n'importe où pour voir les commandes disponible  
et utiliser les dans les projets que vous souhaitez  

## Contruction d'une commande personnalisé dans un projet

- pour passer des variables d'environement il vous faut un `.env`
- creer un fichier `commands.yaml` 
- suivez le format suivant  
```yaml
commands:
  cm: # nom de la commande
    description: # description de celle ci
    command: # mettre ça commande shell ici
```
- vous pouvez ajouter des variables qui seront relier à des variables d'environement avec ce format:
```bash
# variable d'environement dans le .env 
*NSC_NAME*

# exemple
NSC=nseven
```
- vous pouvez ajouter des variables à la volé de la commande avec ce format:
```bash
# variable que l'on affectera depuis la commande
{{nameFile}}

# exemple
ns c mycommand nameFile=nseven
```

- exemple de commande et d'execution de la commande:  
```bash
# commande
commands:
  cm: # nom de la commande
    description: # description de celle ci
    command: docker exec -i *NSC_NAME* goose create -s {{nameFile}} sql --dir ./migrations

# execution
ns c cm nameFile=nseven
```



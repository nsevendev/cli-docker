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


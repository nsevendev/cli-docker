# cli-docker

## Dev

```bash
# build l'image docker du projet
make build

# lance le container docker
make up

# pour arrêter le container
make down
```

## Utilisation

- une fois les containers demarrés, ouvrir vscode dans le container
go est disponible dans le terminal, vous pouvez utiliser tout les commandes go que vous voulez

- le Makefile, sert à 
    - dockerizer ce projet et faire tourner go pour le developpement
    - lancer les build du cli pour toute platforme
tous le reste se fera avec les commandes go classique


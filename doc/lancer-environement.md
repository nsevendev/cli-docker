# Lancement de l'environement

```bash
# build l'image docker du projet
make build

# lance le conteneur docker
make up

# pour arrêter le conteneur
make down
```

- le Makefile, sert à
    - dockerizer ce projet et faire tourner go pour le developpement
    - lancer les build du cli pour toute platforme ou le build en local du projet pour tester  
      tous le reste se fera avec les commandes go classique  

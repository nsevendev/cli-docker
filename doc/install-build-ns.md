# Install ns en local

## Installation

une fois le repo cloner, lancer le conteneur en mode dev et lancer la commande suivante qui vous corresponds:

```bash
# pour macOs amd64
make install-local

# pour macOs arm64
make install-local c="m"

# pour linux amd64
make install-local c="l"
```

## Indications  

*Note : il vous demandera votre mot de passe pour installer l'executable `ns`*

- si tout c'est bien passé utiliser la commande `ns` n'importe où pour voir les commandes disponible  
  et utiliser les dans les projets que vous souhaitez  

# Utiliser l'image officielle de Go
FROM golang:1.21

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Exposer le port sur lequel l'application s'exécutera
EXPOSE 5000

# Démarrer un shell interactif
#CMD [ "bash" ]
CMD ["tail", "-f", "/dev/null"]
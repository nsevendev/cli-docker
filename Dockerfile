FROM golang:1.21
WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

#CMD ["tail", "-f", "/dev/null"]
CMD ["CompileDaemon", "--build=go build -o tmp/ns .", "--command=tmp/ns"]
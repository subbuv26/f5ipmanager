FROM golang

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/f5ipmanager ./cmd/f5ipmanager/

EXPOSE 8080

ENTRYPOINT ["/bin/sh"]

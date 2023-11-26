# GoDB
GoDB is an in-memory, key-value store that allows users to set, get, and delete key-value pairs. The database server supports TCP connections over port 6532. This repository holds the executable for the database server as well as the code for interfacing with the server from a client application.

## Requirements
* Go 1.21

## Running The Server Locally
```
git clone https://github.com/MattLaidlaw/godb
cd ./godb
go build ./cmd/server
./server
```

## Testing The Server And Client Code
```
git clone https://github.com/MattLaidlaw/godb
cd ./godb
go build -v ./...
go test -v ./...
```

## Running The Server Via Docker
```
git clone https://github.com/MattLaidlaw/godb
docker build -t godb
docker run -p 6532:6532 godb
```

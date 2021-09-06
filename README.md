# GoDB
GoDB is an in-memory, key-value store that allows users to set, get, and delete key-value pairs. The database server supports TCP connections over port 6342. Currently, Go is the only language with an available [client](https://github.com/MattLaidlaw/godb-go-driver).

## Requirements
* Go 1.17

## Install / Run
```
git clone https://github.com/MattLaidlaw/godb
cd ./godb
go build ./cmd/godb
./godb
```

## Docker
[laidlawm/godb](https://hub.docker.com/repository/docker/laidlawm/godb)
```
docker run -p 6342:6342 laidlawm/godb:latest
```

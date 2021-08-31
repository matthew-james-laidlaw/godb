# GoDB
GoDB is an in-memory, key-value store. GoDB supports client connections over TCP via port 6342. Currently, the only available client is for the Go language and can be found [here](https://github.com/MattLaidlaw/godb-go-driver). GoDB allows users to set, get, and delete key value pairs.

# Requirements
* Go 1.16

# Install / Run
```
git clone https://github.com/MattLaidlaw/godb
cd ./godb
go build ./cmd/godb
./godb
```

# Docker
[laidlawm/godb](https://hub.docker.com/repository/docker/laidlawm/godb)\
```docker run -p 6342:6342 laidlawm/godb:latest```

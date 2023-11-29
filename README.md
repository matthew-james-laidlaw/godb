# GoDB
GoDB is an in-memory, key-value store that allows users to set, get, and delete key-value pairs. The database server supports TCP connections over port 8000. This repository holds the executable for the database server as well as the code for interfacing with the server from a client application.

# Purpose
This application does not have a lot of capabilities. As a database, it is not very useful, nor is it very robust. It was written as an exercise in learning networking in GoLang. I also explored using a custom, very lightweight, JSON-RPC-esque method for communicating between client and server. This tool was also an exercise in packaging an application with Docker.

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
docker build . -t godb
docker run -p 8000:8000 godb

# and to test an example client after spinning up the server
go run ./cmd/example
```

## Usage

### Creating A Client
```
  client, err := godb.NewClient(":8000")
  if err != nil {
    log.Fatalln(err)
  }
```

### CRUD Operations
```
  # insert a key-value pair
	res, err := client.Set("key", "value")
	if err != nil {
		log.Fatalln(err)
	}

  # retrieve a value from the given key
	res, err := client.Get("key")
	if err != nil {
		log.Fatalln(err)
	}

  # delete a key-value pair
	res, err := client.Del("key")
	if err != nil {
		log.Fatalln(err)
	}
```

### Response Objects
`Response` objects wrap the response from the database to the client. For `Set` and `Del` operations, the `Result` field of the object will contain the number of records inserted or deleted. For `Get` operations, the `Result` field will contain the value that matches the key used in the operation. If no such key-value pair exists, the record will be empty.

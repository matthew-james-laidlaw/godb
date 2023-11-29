package main

import (
	"GoDB/pkg/godb"
	"fmt"
	"log"
)

func main() {

	// assumes a server running on port 8000
	client, err := godb.NewClient(":8000")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Set("key", "value")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Number of records inserted:", res.Result)

	res, err = client.Get("key")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Value of `key`:", res.Result)

}

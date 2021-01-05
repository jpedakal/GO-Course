package main

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	URI := "mongodb://127.0.0.1:27017/?gssapiServiceName=mongodb"
	clientOptions := options.Client().ApplyURI(URI)
}

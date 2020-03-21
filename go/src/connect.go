// go get go.mongodb.org/mongo-driver/mongo
// go run connect.go

package main

import (
    "context"
    "fmt"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        fmt.Println("mongo ctor")
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()
    
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected")
}
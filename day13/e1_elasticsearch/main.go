package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Person struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connect to es succeed.")
	p1 := Person{Name: "eric", Age: "22", Married: false}
	put1, err := client.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

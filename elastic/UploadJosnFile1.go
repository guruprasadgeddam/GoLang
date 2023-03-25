package main

import (
	"context"
	"encoding/json"
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

func main() {
	ctx := context.Background()
	elasticSearchClinet, error := GetESlinet()
	if error != nil {
		fmt.Println("Error in initilizing", error)
		panic("Failed")
	}
	newStudent := Student{
		Name:         "Guru",
		Age:          30,
		AverageScore: 98,
	}
	josnData, err1 := json.Marshal(newStudent)
	js := string(josnData)
	//index, err1 := elastic.Index().In
	_, err1 = elasticSearchClinet.Index().Index("students").BodyJson(js).Do(ctx)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("Elastic Insertion successful")
}

func GetESlinet() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false), elastic.SetHealthcheck(false))
	fmt.Println("ES initialized")
	return client, err
}

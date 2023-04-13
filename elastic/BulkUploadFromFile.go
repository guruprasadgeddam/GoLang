package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	elastic "github.com/elastic/go-elasticsearch/v7"
)

func main() {
	// Create a new Elasticsearch client
	cfg := elastic.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	es, err := elastic.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	// Open the bulk file
	file, err := os.Open("jsontext.txt")
	if err != nil {
		log.Fatalf("Error opening bulk file: %s", err)
	}
	defer file.Close()

	// Read the file line by line and append each line to a strings.Builder object
	scanner := bufio.NewScanner(file)
	var body strings.Builder
	for scanner.Scan() {
		body.WriteString(scanner.Text())
		body.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning bulk file: %s", err)
	}

	// Send the bulk data to Elasticsearch
	res, err := es.Bulk(
		strings.NewReader(body.String()),
		es.Bulk.WithIndex("myindex"),
		es.Bulk.WithDocumentType("_doc"), // For Elasticsearch 7.x, document type should be specified here
	)

	if err != nil {
		log.Fatalf("Error bulk uploading data: %s", err)
	}

	fmt.Println(res)
}

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

func main() {
	fmt.Println("testing::::")
	uploadJsonDocuments()
}

func uploadJsonDocuments() {

	file, error := os.Open("jsondoc.json")
	if error != nil {
		log.Fatalf("os.Open() ERROR:", error)
	}
	defer file.Close()

	// Call ioutil.ReadAll() to create a bytes array from file's JSON data
	byteSlice, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("ioutil.ReadAll() ERROR:", err)
	}
	fmt.Println("bytesStr TYPE:", reflect.TypeOf(byteSlice), "n")
	// Make HTTP request using "PUT" or "POST" verb
	req, reqErr := http.NewRequest("PUT", "http://localhost:9200/_bulk?pretty=true", bytes.NewBuffer(byteSlice))
	req.Header.Set("Content-Type", "application/json")
	if reqErr != nil {
		log.Fatalf("http.NewRequest ERROR:", reqErr)
	} else {
		fmt.Println("HTTP Request:", req)
	}
	// Instantiate a new client object
	client := &http.Client{}
	// pass Http request to elasticsearch client
	res, error := client.Do(req)
	if error != nil {
		log.Fatalf("1client.Do ERROR:", err)
	}

	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)

	if error != nil {
		log.Fatalf("client.Do ERROR:", error)
	}
	// Convert the bytes object []uint8 of the JSON response to a string
	strBody := string(body)

	// Print out the response body
	fmt.Println("nresp.Body:", strBody)
}

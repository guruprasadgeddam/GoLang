package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// sample
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	chunkSize := 1024 // 1KB chunks
	var chunks [][]byte

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize

		if end > len(data) {
			end = len(data)
		}

		chunks = append(chunks, data[i:end])
	}

	fmt.Println(len(chunks))
}

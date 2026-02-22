package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	//fmt.Println("Hello, World!")
	readCsv()

	//fmt.Println("Goodbye, World!")
}

func readCsv() {
	// Open the CSV file
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Process the records
	for _, record := range records {
		fmt.Printf("%v,%v\n", record[0], record[1])
	}
}

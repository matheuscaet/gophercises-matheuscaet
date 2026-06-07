package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//fmt.Println("Hello, World!")
	readCsv()

	//fmt.Println("Goodbye, World!")
}

func readCsv() {
	var timerValue int
	flag.IntVar(&timerValue, "timer", 30, "variable to set timer size in seconds")
	flag.Parse()

	timer := time.NewTimer(time.Duration(timerValue) * time.Second)
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
		select {
		case <-timer.C:
			fmt.Println("Time's up!")
			return
		default:
			// Process the record (for example, print it)
			fmt.Printf("%v,%v\n", record[0], record[1])
		}
	}

	fmt.Println("Finished", len(records))
}

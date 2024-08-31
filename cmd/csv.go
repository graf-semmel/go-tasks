package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var headers = []string{"ID", "Done", "Description", "Created"}
var fileName = "tasks.csv"

func ReadCSV() [][]string {
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist:", fileName)
		return [][]string{headers}
	} else if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
	}

	return records
}

func WriteCSV(tasks []Task) {
	if csvExists() {
		err := os.Remove(fileName)
		if err != nil {
			log.Fatal("Error deleting file:", err)
		}
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	err = csvWriter.Write(headers)
	if err != nil {
		log.Fatal("Error writing headers:", err)
	}

	for _, task := range tasks {
		err = csvWriter.Write([]string{
			strconv.Itoa(task.ID),
			strconv.FormatBool(task.Done),
			task.Description,
			task.Created,
		})
		if err != nil {
			log.Fatal("Error writing to CSV:", err)
		}
	}
}

func csvExists() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

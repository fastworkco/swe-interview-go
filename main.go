package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
}

const csvFilePath = "./items.csv"

// Read items from the CSV file
func ReadItemsFromCSV() ([]Item, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var items []Item
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for i, record := range records {
		if i == 0 { // Skip header row
			continue
		}

		id, _ := strconv.Atoi(record[0])
		price, _ := strconv.ParseFloat(record[2], 64)

		items = append(items, Item{
			ID:    id,
			Name:  record[1],
			Price: price,
		})
	}

	return items, nil
}

// Write items to the CSV file
func WriteItemsToCSV(items []Item) error {
	file, err := os.Create(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	if err := writer.Write([]string{"id", "name", "price"}); err != nil {
		return err
	}

	// Write data rows
	for _, item := range items {
		row := []string{
			strconv.Itoa(item.ID),
			item.Name,
			strconv.FormatFloat(item.Price, 'f', 2, 64),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	r := mux.NewRouter()

	// Item Handlers
	r.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		items, err := ReadItemsFromCSV()
		if err != nil {
			http.Error(w, "Failed to read items", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(items)
	}).Methods("GET")

	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

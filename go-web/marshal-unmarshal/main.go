package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	marshal()
	unmarshal()
}

func marshal() {
	type product struct {
		Name      string
		Price     int
		Published bool
	}

	p := product{
		Name:      "MacBook Pro",
		Price:     1500,
		Published: true,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}

func unmarshal() {
	type product struct {
		Name      string
		Price     int
		Published bool
	}

	jsonData := `{"Name": "MacBook Air", "Price": 900, "Published": true}`

	var p product

	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}

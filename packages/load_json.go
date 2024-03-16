package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/2gbeh/drygin-api/models"
)

func main() {
	file, err := os.Open("../data/customers.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	var tags models.Tags
	json.Unmarshal(bytes, &tags)

	fmt.Printf("json ~ %v", tags.Data)
}
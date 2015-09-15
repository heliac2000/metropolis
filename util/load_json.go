package util

import (
	"encoding/json"
	"log"
	"os"
)

// Load data from a JSON file
//
func LoadDataFromJSONFile(s interface{}, file string) interface{} {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}

	dec := json.NewDecoder(jsonFile)
	if err := dec.Decode(&s); err != nil {
		log.Fatalln(err)
	}

	return s
}

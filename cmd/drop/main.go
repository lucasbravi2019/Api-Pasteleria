package main

import (
	"encoding/xml"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lucasbravi2019/pasteleria/core"
)

type XMLTables struct {
	XMLName xml.Name   `xml:"tables"`
	List    []XMLTable `xml:"table"`
}

type XMLTable struct {
	TableName xml.Name `xml:"table"`
	ID        string   `xml:"id,attr"`
	SQL       string   `xml:",chardata"`
}

const PATH = "bin/tables_drop.xml"

func MigrationLoader() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	if err := processXmlFile(PATH); err != nil {
		log.Fatal(err)
	}
}

func processXmlFile(filePath string) error {
	file, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	var xmlTables XMLTables

	if err := xml.NewDecoder(file).Decode(&xmlTables); err != nil {
		return err
	}

	db := core.GetDatabaseConnection()
	for _, table := range xmlTables.List {
		_, err := db.Exec(table.SQL)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func main() {
	log.Println("Starting drop")
	MigrationLoader()
	log.Println("Finished drop")
}

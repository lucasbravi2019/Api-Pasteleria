package core

import (
	"encoding/xml"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var queries map[string]string = make(map[string]string)

type XMLQueries struct {
	XMLName xml.Name   `xml:"queries"`
	List    []XMLQuery `xml:"query"`
}

type XMLQuery struct {
	XMLName xml.Name `xml:"query"`
	ID      string   `xml:"id,attr"`
	SQL     string   `xml:",chardata"`
}

func GetQueryByName(queryName string) string {
	query := queries[queryName]
	if query == "" {
		log.Fatal("query not found for query name: ", queryName)
	}

	return strings.TrimSpace(query)
}

func QueryLoader() {
	if err := processXmlFileInDirectory(QUERIES_PATH); err != nil {
		log.Fatal(err)
	}
}

func FindQueryByName(queryName string) (string, error) {
	query := queries[queryName]

	if query == STRING_EMPTY {
		return STRING_EMPTY, errors.New("query not found")
	}

	return queries[queryName], nil
}

func processXmlFile(filePath string) error {
	file, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	var xmlQueries XMLQueries

	if err := xml.NewDecoder(file).Decode(&xmlQueries); err != nil {
		return err
	}

	for _, query := range xmlQueries.List {
		queries[query.ID] = query.SQL
	}

	return nil
}

func processXmlFileInDirectory(directoryPath string) error {
	files, err := os.ReadDir(directoryPath)

	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == XML_EXT {
			path := filepath.Join(directoryPath, file.Name())
			if err := processXmlFile(path); err != nil {
				return err
			}
		}
	}

	return nil
}

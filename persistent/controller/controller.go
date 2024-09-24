package main

import "net/http"

type Parser interface {
	parseFile(content []byte, format string)
}

type PDFParser struct {
}

func (pdfParser *PDFParser) parseFile(content []byte, format string) {

}

type MetadataDetails struct {
	Name string
}

// POST: BASEURL/format/{format}/resume/
func fileHandler(w http.ResponseWriter, r http.Reader) {

	// get the body
	// validator for supported] file format

	// send read data to service

	// call the parser based on format
	//[]bytes => marshal it to MetadataDetails struct

	// db call -> pass marshalled object to DB

	// response from here
}

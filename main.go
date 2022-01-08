package main

import (
	"flat-file-parser/parser"
	"fmt"
)

func main() {

	// pass file pointer into Parser
	// add method recevier function to Parser
	// extract lines of file as a Struct of type T
	// add lines to slice
	fn := "flat1.txt"
	recs := parser.ParseTestRecord(fn)
	fmt.Println("data recs: ", recs)
}

// stretch
//add lines to db as entities
// make website to create flat files and read out db data into table
// use kafka to consume new flat files
// use producer to get new files from frontend

package parser

import (
	"bufio"
	"flat-file-parser/types"
	"github.com/ahmedalhulaibi/flatfile"
	"os"
)

type Parser struct {
	FileName string
	file     *os.File
}

func ParseTestRecord(fn string) []types.TestRecord {
	f, err := os.Open(fn)
	checkError(err)
	scanner := bufio.NewScanner(f)

	rs := make([]types.TestRecord, 0)
	for scanner.Scan() {
		l := scanner.Bytes()

		record := types.TestRecord{} // use reflect to pass in type dynamically
		err := flatfile.Unmarshal(l, &record, 0, 0, false)
		checkError(err)

		rs = append(rs, record)
	}
	return rs
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

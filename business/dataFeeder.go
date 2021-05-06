package business

import (
	"encoding/csv"
	"io"
	"os"
)

type DataFeeder struct {
	feeds []*dataFeed
}

func (df *DataFeeder) LoadFromCsv(csvFile *os.File) error {
	csvReader := csv.NewReader(csvFile)
	isHeader := true
	for csvLine, err := csvReader.Read(); err != io.EOF; csvLine, err = csvReader.Read() {
		if isHeader {
			isHeader = false
			continue
		}

		dataFeed, err := parse(csvLine[0])
		if err != nil {
			return err
		}

		df.feeds = append(df.feeds, dataFeed)
	}
	return nil
}

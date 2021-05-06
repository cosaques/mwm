package business

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
)

type DataFeeder struct {
	feeds []*dataFeed
}

func (df *DataFeeder) AdminHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://www.data.gouv.fr/fr/datasets/r/406c6a23-e283-4300-9484-54e78c8ae675")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	err = df.loadFromCsv(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(fmt.Sprintf("Loaded %d feeds from data.gouv.fr CSV !", len(df.feeds))))
	w.WriteHeader(http.StatusOK)
}

func (df *DataFeeder) loadFromCsv(csvFile io.Reader) error {
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

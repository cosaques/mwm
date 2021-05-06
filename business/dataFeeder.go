package business

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type dataFeeder struct {
	feeds []*dataFeed
	deps  map[string]bool
}

func NewDataFeeder() *dataFeeder {
	return &dataFeeder{
		deps: make(map[string]bool),
	}
}

func (df *dataFeeder) AdminHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://www.data.gouv.fr/fr/datasets/r/406c6a23-e283-4300-9484-54e78c8ae675")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	err = df.loadFromCsv(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Loaded %d feeds from data.gouv.fr CSV !", len(df.feeds))
}

func (df *dataFeeder) ApiHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[3]

	switch action {
	case "list":
		resp := struct {
			Departments []string
		}{}
		for key := range df.deps {
			resp.Departments = append(resp.Departments, key)
		}
		w.WriteHeader(200)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(resp)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Action %s not found", action)
	}
}

func (df *dataFeeder) loadFromCsv(csvFile io.Reader) error {
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

		df.addDataFeed(dataFeed)
	}
	return nil
}

func (df *dataFeeder) addDataFeed(feed *dataFeed) {
	df.feeds = append(df.feeds, feed)
	df.deps[feed.department] = true
}

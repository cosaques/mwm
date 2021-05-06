package business

import (
	"os"
	"testing"
)

func TestDataFeeder_LoadFromCsv(t *testing.T) {
	csvFile, err := os.Open("testdata/data.csv")
	if err != nil {
		t.Fatal("cannot open file")
	}

	df := &DataFeeder{}
	err = df.LoadFromCsv(csvFile)
	if err != nil {
		t.Fatalf("error %s happened", err)
	}

	if len(df.feeds) != 3 {
		t.Fatal("csv was badly parsed")
	}
}

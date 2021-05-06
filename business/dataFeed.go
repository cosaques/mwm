package business

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type dataFeed struct {
	department string
	date       time.Time
	posTests   int
	totalTests int
	age        string
	population int
}

// parse parses string in format
// 01;2020-05-13;0;16;09;83001
func parse(s string) (*dataFeed, error) {
	feed := strings.Split(s, ";")
	if len(feed) != 6 {
		return nil, errors.New("bad format")
	}

	date, err := time.Parse("2006-01-02", feed[1])
	if err != nil {
		return nil, err
	}

	posTests, err := strconv.Atoi(feed[2])
	if err != nil {
		return nil, err
	}

	totalTests, err := strconv.Atoi(feed[3])
	if err != nil {
		return nil, err
	}

	population, err := strconv.Atoi(feed[5])
	if err != nil {
		return nil, err
	}

	return &dataFeed{
		department: feed[0],
		date:       date,
		posTests:   posTests,
		totalTests: totalTests,
		age:        feed[4],
		population: population,
	}, nil
}

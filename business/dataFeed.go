package business

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type dataFeed struct {
	Department string    `json:"Department"`
	Date       time.Time `json:"Date"`
	PosTests   int       `json:"PositiveCases"`
	TotalTests int       `json:"TotalTests"`
	Age        string    `json:"Age"`
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

	return &dataFeed{
		Department: feed[0],
		Date:       date,
		PosTests:   posTests,
		TotalTests: totalTests,
		Age:        feed[4],
	}, nil
}

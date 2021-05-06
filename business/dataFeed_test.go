package business

import (
	"reflect"
	"testing"
	"time"
)

func Test_parse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *dataFeed
		wantErr bool
	}{
		{
			"CorrectString",
			args{"01;2020-05-13;0;16;09;83001"},
			&dataFeed{"01", time.Date(2020, 5, 13, 0, 0, 0, 0, time.UTC), 0, 16, "09", 83001},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

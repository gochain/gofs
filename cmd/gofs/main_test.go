package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/alecthomas/units"
)

func Test_prettySize(t *testing.T) {
	tests := []struct {
		b    int64
		want string
	}{
		{1, "1B"},
		{1000, "1.00KB"},
		{1005, "1.01KB"},
		{1000000, "1.00MB"},
		{1400000, "1.40MB"},
		{1000000000, "1.00GB"},
		{int64(units.EB), "1.00EB"},
	}
	for _, tt := range tests {
		t.Run(strconv.FormatInt(tt.b, 10), func(t *testing.T) {
			if got := prettySize(tt.b); got != tt.want {
				t.Errorf("got %q; want %q", got, tt.want)
			}
		})
	}
}

func Test_parseDate(t *testing.T) {
	tests := []struct {
		date    string
		y       int
		m       int64
		d       int
		wantErr bool
	}{
		{"2019-09-10", 2019, 9, 10, false},
		{"2019-09-2", 2019, 9, 2, false},
		{"2019-09", 2019, 9, 1, false},
		{"2019-9", 2019, 9, 1, false},
		{"2019", 2019, 1, 1, false},

		{"2019-", 0, 0, 0, true},
		{"2019-09-", 0, 0, 0, true},
		{"2019-09-10-", 0, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.date, func(t *testing.T) {
			y, m, d, err := parseDate(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("wrong err: got %v; wantErr %v", err, tt.wantErr)
				return
			}
			if y != tt.y {
				t.Errorf("wrong year got %v; want %v", y, tt.y)
			}
			if m != time.Month(tt.m) {
				t.Errorf("wrong month got1 %v; want %v", m, tt.m)
			}
			if d != tt.d {
				t.Errorf("wrong day got2 %v; want %v", d, tt.d)
			}
		})
	}
}

func Test_parseDuration(t *testing.T) {
	tests := []struct {
		dur     string
		y       int
		d       int
		wantErr bool
	}{
		{"1y", 1, 0, false},
		{"10y", 10, 0, false},
		{"100d", 0, 100, false},
		{"1y1d", 1, 1, false},
		{"10y100d", 10, 100, false},

		{"", 0, 0, true},
		{"1d1y", 0, 0, true},
		{"1yy", 0, 0, true},
		{"1y1dd", 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.dur, func(t *testing.T) {
			y, d, err := parseDuration(tt.dur)
			if (err != nil) != tt.wantErr {
				t.Errorf("wrong err: got %v; wantErr %v", err, tt.wantErr)
				return
			}
			if y != tt.y {
				t.Errorf("wrong year: got %v; want %v", y, tt.y)
			}
			if d != tt.d {
				t.Errorf("wrong day: got %v; want %v", d, tt.d)
			}
		})
	}
}

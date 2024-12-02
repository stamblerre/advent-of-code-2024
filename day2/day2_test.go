package main

import (
	"testing"
)

func TestDay2(t *testing.T) {
	for _, tc := range []struct {
		filename       string
		numSafeReports int
		useDampener    bool
	}{
		{
			filename:       "testdata/test_input.txt",
			numSafeReports: 2,
		},
		{
			filename:       "testdata/test_input.txt",
			numSafeReports: 4,
			useDampener:    true,
		},
	} {
		gotSafeReports, err := countSafeReports(tc.filename, tc.useDampener)
		if err != nil {
			t.Fatal(err)
		}
		if gotSafeReports != tc.numSafeReports {
			t.Fatalf("expected %d, got %d", tc.numSafeReports, gotSafeReports)
		}
	}
}

package main

import "testing"

func TestDay5Part1(t *testing.T) {
	for _, tc := range []struct {
		filename         string
		validUpdateCount int
		fixedUpdateCount int
	}{
		{
			filename:         "testdata/test_input.txt",
			validUpdateCount: 143,
			fixedUpdateCount: 123,
		},
	} {
		gotValidUpdateCount, gotFixedUpdateCount, err := validateUpdates(tc.filename)
		if err != nil {
			t.Fatal(err)
		}
		if gotValidUpdateCount != tc.validUpdateCount {
			t.Fatalf("expected %d, got %d", tc.validUpdateCount, gotValidUpdateCount)
		}
		if gotFixedUpdateCount != tc.fixedUpdateCount {
			t.Fatalf("expected %d, got %d", tc.fixedUpdateCount, gotFixedUpdateCount)
		}
	}
}

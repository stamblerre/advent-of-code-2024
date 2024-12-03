package main

import (
	"testing"
)

func TestDay3(t *testing.T) {
	for _, tc := range []struct {
		filename string
		result   int
	}{
		{
			filename: "testdata/test_input.txt",
			result:   161,
		},
		{
			filename: "testdata/test_input2.txt",
			result:   48,
		},
	} {
		got, err := multiply(tc.filename)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.result {
			t.Fatalf("expected %d, got %d", tc.result, got)
		}
	}
}

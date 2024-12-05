package main

import "testing"

func TestDay4Part1(t *testing.T) {
	for _, tc := range []struct {
		filename string
		result   int
	}{
		{
			filename: "testdata/small_test_1.txt",
			result:   3,
		},
		{
			filename: "testdata/test_input.txt",
			result:   18,
		},
	} {
		got, err := wordSearchForXmas(tc.filename)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.result {
			t.Fatalf("expected %d, got %d", tc.result, got)
		}
	}
}

func TestDay4Part2(t *testing.T) {
	for _, tc := range []struct {
		filename string
		result   int
	}{
		{
			filename: "testdata/test_input.txt",
			result:   9,
		},
	} {
		got, err := wordSearchForMasXed(tc.filename)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.result {
			t.Fatalf("expected %d, got %d", tc.result, got)
		}
	}
}

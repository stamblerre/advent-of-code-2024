package main

import "testing"

func TestDay6(t *testing.T) {
	for _, tc := range []struct {
		filename  string
		expected1 int
		expected2 int
	}{
		{
			filename:  "testdata/test_input.txt",
			expected1: 41,
			expected2: 6,
		},
	} {
		got1, got2, err := run(tc.filename)
		if err != nil {
			t.Fatal(err)
		}
		if got1 != tc.expected1 {
			t.Fatalf("expected %d, got %d", tc.expected1, got1)
		}
		if got2 != tc.expected2 {
			t.Fatalf("expected %d, got %d", tc.expected2, got2)
		}
	}
}

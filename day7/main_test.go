package main

import (
	"testing"

	"advent-of-code-2024.com/internal/shared"
)

func TestDay7(t *testing.T) {
	for _, tc := range []struct {
		filename  string
		expected1 int
		expected2 int
	}{
		{
			filename:  "testdata/test_input.txt",
			expected1: 3749,
			expected2: 11387,
		},
	} {
		day := &today{}
		got1, got2, err := shared.Run(day, tc.filename)
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

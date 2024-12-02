package main

import "testing"

func TestDay1(t *testing.T) {
	for _, tc := range []struct {
		list1           []int
		list2           []int
		minDistance     int
		similarityScore int
	}{
		{
			list1:           []int{3, 4, 2, 1, 3, 3},
			list2:           []int{4, 3, 5, 3, 9, 3},
			minDistance:     11,
			similarityScore: 31,
		},
	} {
		gotMinDistance := getMinDistance(tc.list1, tc.list2)
		if gotMinDistance != tc.minDistance {
			t.Fatalf("expected %d, got %d", tc.minDistance, gotMinDistance)
		}
		gotSimilarityScore := getSimilarityScore(tc.list1, tc.list2)
		if gotSimilarityScore != tc.similarityScore {
			t.Fatalf("expected %d, got %d", tc.similarityScore, gotSimilarityScore)
		}
	}
}

package histograms

import (
	"math/rand"
	"testing"
	"time"
)

func TestHistogram(t *testing.T) {
	h := NewHistogram()
	h.MaxBins = 3
	h.Granularity = time.Second

	for i := 0; i < 5; i++ {
		for i := 0; i < 1000; i++ {
			h.Add(rand.Float64())
		}
		time.Sleep(time.Second)
	}

	distributions := h.Distributions()
	assertEqual(t, len(distributions), 3, "Error on distributions number")

	for _, distribution := range distributions {
		count := 0
		for _, centroid := range distribution.Centroids {
			count += centroid.Count
		}
		assertEqual(t, count, 1000, "Error on centroids count")
	}

	distributions = h.Distributions()
	assertEqual(t, len(distributions), 0, "Error on distributions number")
}

func assertEqual(t *testing.T, a interface{}, b interface{}, e string) {
	if a != b {
		t.Fatalf("%s - %v != %v", e, a, b)
	}
}

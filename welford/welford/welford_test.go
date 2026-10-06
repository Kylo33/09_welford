package welford_test

import (
	"math"
	"testing"

	"git.devmail.group/renn/welford/welford"
)

func assertFloatsEqual(t testing.TB, got, want float64) {
	tol := 1e-6
	if math.Abs(got-want) > tol {
		t.Fatalf("got %v, want %v (tol %v)", got, want, tol)
	}
}

func TestMean(t *testing.T) {
	data := []float64{1, 2, 3, 4}

	w := welford.Welford{}
	for _, dataPoint := range data {
		w.Add(dataPoint)
	}

	got := w.Mean
	want := 2.5

	assertFloatsEqual(t, got, want)
}

func TestCount(t *testing.T) {
	data := []float64{1, 2, 3, 4}

	w := welford.Welford{}
	for _, dataPoint := range data {
		w.Add(dataPoint)
	}

	got := w.Count
	want := 4

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestStdev(t *testing.T) {
	data := []float64{1, 2, 3, 4}

	w := welford.Welford{}

	for _, dataPoint := range data {
		w.Add(dataPoint)
	}

	got := w.Stdev()
	want := math.Sqrt(1.25)

	assertFloatsEqual(t, got, want)
}

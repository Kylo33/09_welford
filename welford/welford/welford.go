package welford

import "math"

type Welford struct {
	Count          int
	Mean, Variance float64
}

func (w *Welford) Add(data float64) {
	w.Count++

	oldMean := w.Mean
	oldVariance := w.Variance

	w.Mean = oldMean + (data-oldMean)/float64(w.Count)
	w.Variance = oldVariance + math.Pow(data-oldMean, 2) - math.Pow(data-oldMean, 2)/float64(w.Count)
}

func (w *Welford) Stdev() float64 {
	return math.Sqrt(w.Variance / float64(w.Count))
}

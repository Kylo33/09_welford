package samplepath

import (
	"math"
)

type SamplePath struct {
	Time, Mean, Variance float64
}

func (s *SamplePath) Stdev() float64 {
	return math.Sqrt(s.Variance / float64(s.Time))
}

func (s *SamplePath) Add(dataPoint, time float64) {
	oldTime := s.Time
	oldMean := s.Mean
	oldVariance := s.Variance

	s.Time = time
	s.Mean = oldMean + (time-oldTime)/time*(dataPoint-oldMean)
	s.Variance = oldVariance + (time-oldTime)*oldTime/time*math.Pow(dataPoint-oldMean, 2)
}

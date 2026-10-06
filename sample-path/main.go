package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"git.devmail.group/renn/samplepath/samplepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: samplepath <dataset path>")
	}

	dataset, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading dataset: %v", err)
	}

	s := samplepath.SamplePath{}
	for line := range strings.Lines(string(dataset)) {
		fields := strings.Fields(line)

		time, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			log.Fatalf("Error parsing data point: %v", err)
		}

		dataPoint, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Fatalf("Error parsing data point: %v", err)
		}

		s.Add(dataPoint, time)
	}

	fmt.Printf("%.8e\t%.8e", s.Mean, s.Stdev())
}

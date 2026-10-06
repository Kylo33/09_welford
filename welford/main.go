package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"git.devmail.group/renn/welford/welford"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: welford <dataset path>")
	}

	dataset, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading dataset: %v", err)
	}

	w := welford.Welford{}
	for dataString := range strings.FieldsSeq(string(dataset)) {
		dataPoint, err := strconv.ParseFloat(dataString, 64)
		if err != nil {
			log.Fatalf("Error parsing data point: %v", dataString)
		}
		w.Add(dataPoint)
	}

	fmt.Printf("%.8e\t%.8e", w.Mean, w.Stdev())
}

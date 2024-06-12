package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"MathSkill/function"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s data.txt\n", os.Args[0])
	}

	filePath := os.Args[1]
	data, err := function.ReadData(filePath)
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	average := function.CalculateAverage(data)
	median := function.CalculateMedian(data)
	variance := function.CalculateVariance(data, average)
	stdDeviation := function.CalculateStdDeviation(variance)

	fmt.Printf("Average: %d\n", int(math.Round(average)))
	fmt.Printf("Median: %d\n", int(math.Round(median)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDeviation)))
}

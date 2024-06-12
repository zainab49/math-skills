package function

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"math"
)

func ReadData(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func CalculateAverage(data []int) float64 {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return float64(sum) / float64(len(data))
}

func CalculateMedian(data []int) float64 {
	sort.Ints(data)
	n := len(data)
	if n%2 == 0 {
		return float64(data[n/2-1]+data[n/2]) / 2.0
	}
	return float64(data[n/2])
}

func CalculateVariance(data []int, mean float64) float64 {
	var sumSquares float64
	for _, value := range data {
		sumSquares += (float64(value) - mean) * (float64(value) - mean)
	}
	return sumSquares / float64(len(data))
}

func CalculateStdDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}

package day01

import (
	"bufio"
	"bytes"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Locations []int64

func readInput(path string) ([]Locations, error) {
	var locations []Locations

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(file))

	var locationOne Locations
	var locationTwo Locations
	for scanner.Scan() {
		line := scanner.Text()
		lineData := strings.Split(line, "   ")
		locOne, _ := strconv.ParseInt(lineData[0], 10, 64)
		locationOne = append(locationOne, locOne)
		locTwo, _ := strconv.ParseInt(lineData[1], 10, 64)
		locationTwo = append(locationTwo, locTwo)
	}

	locations = append(locations, locationOne)
	locations = append(locations, locationTwo)

	return locations, nil
}

func ProcessAnswer(path string) int {
	locations, _ := readInput(path)

	slices.SortStableFunc(locations[0], func(a, b int64) int {
		if a < b {
			return -1
		}

		if a > b {
			return 1
		}

		return 0
	})

	slices.SortStableFunc(locations[1], func(a, b int64) int {
		if a < b {
			return -1
		}

		if a > b {
			return 1
		}

		return 0
	})

	var sum int
	for i := range locations[1] {
		sum += int(math.Abs(float64(int(locations[1][i]) - int(locations[0][i]))))

	}
	return sum
}

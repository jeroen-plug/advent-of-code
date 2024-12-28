package day6

import (
	"bufio"
	"log"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

type RaceRecord struct {
	Time     int
	Distance int
}

func Solution() (any, any) {
	f, err := os.Open(path.Join("input", "6.txt"))
	if err != nil {
		log.Fatal("Could not read input 6")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return day6a(lines), day6b(lines)
}

func day6a(lines []string) int {
	records := parse(lines)
	result := 1

	for _, r := range records {
		// x(t-x)-d => -x2 + tx - d
		low, high := calculateRoots(-1, float64(r.Time), float64(-r.Distance))
		start := int(math.Ceil(low + 0.00000000001))
		end := int(math.Floor(high - 0.00000000001))
		result *= end - start + 1
	}

	return result
}

func day6b(lines []string) int {
	time, err := strconv.Atoi(strings.Join(strings.Fields(strings.TrimPrefix(lines[0], "Time:")), ""))
	if err != nil {
		log.Fatal(err)
	}
	distance, err := strconv.Atoi(strings.Join(strings.Fields(strings.TrimPrefix(lines[1], "Distance:")), ""))
	if err != nil {
		log.Fatal(err)
	}

	// x(t-x)-d => -x2 + tx - d
	low, high := calculateRoots(-1, float64(time), float64(-distance))
	start := int(math.Ceil(low + 0.00000000001))
	end := int(math.Floor(high - 0.00000000001))

	return end - start + 1
}

func calculateRoots(a float64, b float64, c float64) (float64, float64) {
	peak := -b / (2 * a)
	delta := math.Abs(math.Sqrt(math.Pow(b, 2)-4*a*c) / (2 * a))
	return peak - delta, peak + delta
}

func parse(lines []string) []RaceRecord {
	times := strings.Fields(strings.TrimPrefix(lines[0], "Time:"))
	distances := strings.Fields(strings.TrimPrefix(lines[1], "Distance:"))

	var records []RaceRecord
	for i := range times {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			continue
		}
		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			continue
		}
		records = append(records, RaceRecord{time, distance})
	}

	return records
}

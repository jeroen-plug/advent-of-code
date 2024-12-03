package input

import (
	"bufio"
	"log"
	"os"
	"path"
	"strconv"
)

func String(day int) string {
	input, err := os.ReadFile(path.Join("input", strconv.Itoa(day)+".txt"))
	if err != nil {
		log.Fatalf("Could not read input %d", day)
	}
	return string(input)
}

func Lines(day int) []string {
	f, err := os.Open(path.Join("input", strconv.Itoa(day)+".txt"))
	if err != nil {
		log.Fatalf("Could not read input %d", day)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

package input

import (
	"log"
	"strconv"
)

func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Could not parse number %s", s)
	}
	return n
}

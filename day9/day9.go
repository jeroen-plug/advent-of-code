package day9

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

func Day9() {
	data := strings.TrimSuffix(input.String(9), "\n")

	fmt.Printf("day 9a: %d\n", day9a(data))
	fmt.Printf("day 9b: %d\n", day9b(data))
}

func day9a(data string) int {
	return calculateChecksum(defragmentFilesystem(parseDiskmap(data)))
}

func day9b(data string) int {
	return calculateChecksum(defragmentFiles(parseDiskmap(data)))
}

func parseDiskmap(diskmap string) []int {
	var fileId int
	var fs []int
	for i, b := range diskmap {
		n := -1
		if i%2 == 0 {
			n = fileId
			fileId++
		}

		fs = slices.Grow(fs, int(b-'0'))
		for j := 0; j < int(b-'0'); j++ {
			fs = append(fs, n)
		}
	}
	return fs
}

func defragmentFilesystem(fs []int) []int {
	lastBlock := len(fs) - 1
	for i := range fs {
		if fs[i] >= 0 {
			continue
		}
		for ; fs[lastBlock] < 0 && lastBlock > 0; lastBlock-- {
		}
		if lastBlock <= i {
			break
		}
		fs[i] = fs[lastBlock]
		fs[lastBlock] = -1
	}
	return fs
}

func defragmentFiles(fs []int) []int {
	firstBlock := len(fs) - 1
	lastBlock := firstBlock
	currentFile := fs[lastBlock]
	for i := lastBlock; i > 0; i-- {
		if fs[i] < 0 {
			continue
		} else if fs[i] == currentFile {
			firstBlock = i
			continue
		}

		fileSize := lastBlock - firstBlock + 1

		spotStart := 0
		spotSize := 0
		for j := 0; j < firstBlock; j++ {
			if fs[j] >= 0 {
				spotStart = j + 1
				spotSize = 0
			} else {
				if spotSize++; spotSize >= fileSize {
					break
				}
			}
		}

		if spotSize >= fileSize {
			for j := 0; j < fileSize; j++ {
				fs[spotStart+j] = fs[firstBlock+j]
				fs[firstBlock+j] = -1
			}
		}

		firstBlock = i
		lastBlock = firstBlock
		currentFile = fs[lastBlock]
	}
	return fs
}

func calculateChecksum(fs []int) int {
	checksum := 0
	for i, f := range fs {
		if f > 0 {
			checksum += i * f
		}
	}
	return checksum
}

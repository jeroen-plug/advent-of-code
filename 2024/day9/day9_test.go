package day9

import (
	"slices"
	"testing"
)

const diskmap = "2333133121414131402"
const filesystem = "00...111...2...333.44.5555.6666.777.888899"
const defragmented = "0099811188827773336446555566.............."
const checksum = 1928

const defragmentedFiles = "00992111777.44.333....5555.6666.....8888.."
const checksumFiles = 2858

func TestDay9a(t *testing.T) {
	want := checksum
	res := day9a(diskmap)

	if res != want {
		t.Fatalf("day9a() = %d, want %d", res, want)
	}
}

func TestDay9b(t *testing.T) {
	want := checksumFiles
	res := day9b(diskmap)

	if res != want {
		t.Fatalf("day9b() = %d, want %d", res, want)
	}
}

func TestParseDiskmap(t *testing.T) {
	want := toFs(filesystem)
	res := parseDiskmap(diskmap)

	if !slices.Equal(res, want) {
		t.Fatalf("parseDiskmap() = %d, want %d", res, want)
	}
}

func TestDefragmentFilesystem(t *testing.T) {
	want := toFs(defragmented)
	res := defragmentFilesystem(toFs(filesystem))

	if !slices.Equal(res, want) {
		t.Fatalf("defragmentFilesystem() = %d, want %d", res, want)
	}
}

func TestDefragmentFiles(t *testing.T) {
	want := toFs(defragmentedFiles)
	res := defragmentFiles(toFs(filesystem))

	if !slices.Equal(res, want) {
		t.Fatalf("defragmentFiles() = %d, want %d", res, want)
	}
}

func TestCalculateChecksum(t *testing.T) {
	want := checksum
	res := calculateChecksum(toFs(defragmented))

	if res != want {
		t.Fatalf("calculateChecksum() = %d, want %d", res, want)
	}
}

func toFs(s string) []int {
	var fs []int
	for _, b := range s {
		if b == '.' {
			fs = append(fs, -1)
		} else {
			fs = append(fs, int(b-'0'))
		}
	}
	return fs
}

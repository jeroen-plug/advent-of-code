package day11

import (
	"fmt"
	"testing"
)

const example = `125 17`

func TestDay11(t *testing.T) {
	var tests = []struct {
		blink int
		want  int
	}{
		{6, 22},
		{25, 55312},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("day11(%d)", tt.blink), func(t *testing.T) {
			res := day11(example, tt.blink)
			if res != tt.want {
				t.Fatalf("got %d, want %d", res, tt.want)
			}
		})
	}
}

func BenchmarkDay11(b *testing.B) {
	tests := []int{6, 12, 25, 50, 100}
	for _, blinks := range tests {
		b.Run(fmt.Sprintf("day11(%d)", blinks), func(b *testing.B) {
			for range b.N {
				day11(example, blinks)
			}
		})
	}
}

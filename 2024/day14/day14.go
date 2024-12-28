package day14

import (
	"cmp"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code/2024/input"
)

type Robot struct {
	Position [2]int
	Velocity [2]int
}

func Solution() (any, any) {
	lines := input.Lines(14)
	day14render(lines, 101, 103)
	return day14a(lines, 101, 103), fmt.Sprintf("%d; All results are in ./out14", day14b(lines, 101, 103))
	// ffmpeg -framerate 100 -i out14/step%04d.png out14/video.mp4
}

func day14a(lines []string, width, height int) int {
	robots := parse(lines)
	robots = move(robots, 100, width, height)
	quadrants := groupByQuadrant(robots, width, height)
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func day14b(lines []string, width, height int) []int {
	robots := parse(lines)

	var results [][2]int

	for i := range 9999 {
		robots = move(robots, 1, width, height)
		quadrants := groupByQuadrant(robots, width, height)
		safetyScore := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
		results = append(results, [2]int{i + 1, safetyScore})
	}

	slices.SortFunc(results, func(a, b [2]int) int { return cmp.Compare(a[1], b[1]) })
	var top3 []int
	for _, r := range results[:3] {
		top3 = append(top3, r[0])
	}
	return top3
}

func day14render(lines []string, width, height int) {
	_, err := os.Lstat("out14")
	if !os.IsNotExist(err) {
		return
	}

	os.Mkdir("out14", os.ModePerm)

	robots := parse(lines)
	img := render(robots, width, height)
	f, _ := os.Create(fmt.Sprintf("out14/step%04d.png", 0))
	png.Encode(f, img)

	for i := range 9999 {
		robots = move(robots, 1, width, height)
		img = render(robots, width, height)
		f, _ = os.Create(fmt.Sprintf("out14/step%04d.png", i+1))
		png.Encode(f, img)
	}
}

func groupByQuadrant(robots []Robot, width, height int) [4]int {
	center := [2]int{width / 2, height / 2}
	var quadrants [4]int
	for _, r := range robots {
		if r.Position[0] < center[0] && r.Position[1] < center[1] {
			quadrants[0]++
		} else if r.Position[0] > center[0] && r.Position[1] < center[1] {
			quadrants[1]++
		} else if r.Position[0] < center[0] && r.Position[1] > center[1] {
			quadrants[2]++
		} else if r.Position[0] > center[0] && r.Position[1] > center[1] {
			quadrants[3]++
		}
	}
	return quadrants
}

func move(robots []Robot, seconds, width, height int) []Robot {
	var moved []Robot
	for _, r := range robots {
		r.Position[0] = (r.Position[0] + (r.Velocity[0] * seconds)) % width
		r.Position[1] = (r.Position[1] + (r.Velocity[1] * seconds)) % height
		if r.Position[0] < 0 {
			r.Position[0] += width
		}
		if r.Position[1] < 0 {
			r.Position[1] += height
		}
		moved = append(moved, r)
	}
	return moved
}

func visualize(robots []Robot, width, height int) {
	tiles := make(map[[2]int]int)
	for _, r := range robots {
		tiles[r.Position]++
	}

	for y := range height {
		for x := range width {
			n := tiles[[2]int{x, y}]
			if n > 0 {
				fmt.Print(n)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func render(robots []Robot, width, height int) image.Image {
	tiles := make(map[[2]int]int)
	for _, r := range robots {
		tiles[r.Position]++
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := range height {
		for x := range width {
			n := tiles[[2]int{x, y}]
			if n > 0 {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}
	return img
}

func parse(lines []string) []Robot {
	var robots []Robot
	for _, l := range lines {
		fields := strings.Fields(l)
		position := strings.Split(strings.TrimPrefix(fields[0], "p="), ",")
		velocity := strings.Split(strings.TrimPrefix(fields[1], "v="), ",")
		robots = append(robots, Robot{
			[2]int{input.ParseInt(position[0]), input.ParseInt(position[1])},
			[2]int{input.ParseInt(velocity[0]), input.ParseInt(velocity[1])},
		})
	}
	return robots
}

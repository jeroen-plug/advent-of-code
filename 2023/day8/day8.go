package day8

import (
	"fmt"
	"strings"

	"github.com/jeroen-plug/advent-of-code/2023/input"
)

type Documents struct {
	Instructions []byte // Left:0, Right:1
	Network      [][2]uint16
}

func Solution() (any, any) {
	lines := input.Lines(8)
	return day8a(lines), day8b(lines)
}

func day8a(lines []string) int {
	d := parse(lines)

	location := toAddress("AAA")
	steps := 0
	for ; location != toAddress("ZZZ"); steps++ {
		step := d.Instructions[steps%len(d.Instructions)]
		location = d.Network[location][step]
	}

	return steps
}

func day8b(lines []string) int {
	d := parse(lines)

	var startLocations []uint16
	for i, next := range d.Network {
		if i&0x1F == int(toAddress("AAA"))&0x1F && next[0] != toAddress("AAA") {
			startLocations = append(startLocations, uint16(i))
		}
	}

	var cycleLength []int
	for _, start := range startLocations {
		location := start
		steps := 0
		for ; location&0x1F != toAddress("ZZZ")&0x1F; steps++ {
			step := d.Instructions[steps%len(d.Instructions)]
			location = d.Network[location][step]
		}
		cycleLength = append(cycleLength, steps)
	}

	return findLcm(cycleLength)
}

func parse(lines []string) Documents {
	var d Documents

	for _, direction := range lines[0] {
		if direction == 'L' {
			d.Instructions = append(d.Instructions, 0)
		} else if direction == 'R' {
			d.Instructions = append(d.Instructions, 1)
		}
	}

	d.Network = make([][2]uint16, toAddress("ZZZ")+1)
	for _, l := range lines[2:] {
		if toAddress(l[0:3]) > toAddress("ZZZ") {
			fmt.Println(l)
		}
		d.Network[toAddress(l[0:3])] = [2]uint16{toAddress(l[7:10]), toAddress(l[12:15])}
	}

	return d
}

func toAddress(node string) uint16 {
	return uint16(node[0]-'A')<<10 + uint16(node[1]-'A')<<5 + uint16(node[2]-'A')
}

func fromAddress(address uint16) string {
	var node strings.Builder
	node.WriteByte(byte(address>>10) + 'A')
	node.WriteByte(byte(address>>5&0x1F) + 'A')
	node.WriteByte(byte(address>>0&0x1F) + 'A')
	return node.String()
}

func findLcm(a []int) int {
	lcm := a[0]
	for _, n := range a[1:] {
		lcm = (n * lcm) / gcd(n, lcm)
	}
	return lcm
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

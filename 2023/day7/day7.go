package day7

import (
	"cmp"
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code/2023/input"
)

type HandType int

const (
	FiveOfAKind  HandType = 0x7
	FourOfAKind  HandType = 0x6
	FullHouse    HandType = 0x5
	ThreeOfAKind HandType = 0x4
	TwoPair      HandType = 0x3
	OnePair      HandType = 0x2
	HighCard     HandType = 0x1
)

func Solution() (any, any) {
	lines := input.Lines(7)
	return day7a(lines), day7b(lines)
}

func day7a(lines []string) int {
	var bidsWithStr [][2]int
	for _, l := range lines {
		fields := strings.Fields(l)

		hand := fields[0]
		bid := input.ParseInt(fields[1])
		handType := getTypeA(hand)

		strength := int(handType)<<(4*5) +
			getValueA(hand[0])<<(4*4) +
			getValueA(hand[1])<<(4*3) +
			getValueA(hand[2])<<(4*2) +
			getValueA(hand[3])<<(4*1) +
			getValueA(hand[4])<<(4*0)

		bidsWithStr = append(bidsWithStr, [2]int{bid, strength})
	}

	slices.SortFunc(bidsWithStr, func(a, b [2]int) int {
		return cmp.Compare(a[1], b[1])
	})

	winnings := 0
	for rank, bid := range bidsWithStr {
		winnings += (rank + 1) * bid[0]
	}

	return winnings
}

func getValueA(card byte) int {
	switch card {
	case 'A':
		return 0xE
	case 'K':
		return 0xD
	case 'Q':
		return 0xC
	case 'J':
		return 0xB
	case 'T':
		return 0xA
	default:
		return int(card) - '0'
	}
}

func getTypeA(hand string) HandType {
	h := strings.Split(hand, "")
	slices.Sort(h)

	i := 0
	curr := h[0]
	var groups []int
	for _, card := range h {
		if card != curr {
			groups = append(groups, i)
			curr = card
			i = 0
		}
		i++
	}
	groups = append(groups, i)

	slices.Sort(groups)
	slices.Reverse(groups)

	if groups[0] == 5 {
		return FiveOfAKind
	} else if groups[0] == 4 {
		return FourOfAKind
	} else if groups[0]+groups[1] == 5 {
		return FullHouse
	} else if groups[0] == 3 {
		return ThreeOfAKind
	} else if groups[0]+groups[1] == 4 {
		return TwoPair
	} else if groups[0] == 2 {
		return OnePair
	} else {
		return HighCard
	}
}

func day7b(lines []string) int {
	var bidsWithStr [][2]int
	for _, l := range lines {
		fields := strings.Fields(l)

		hand := fields[0]
		bid := input.ParseInt(fields[1])
		handType := getTypeB(hand)

		strength := int(handType)<<(4*5) +
			getValueB(hand[0])<<(4*4) +
			getValueB(hand[1])<<(4*3) +
			getValueB(hand[2])<<(4*2) +
			getValueB(hand[3])<<(4*1) +
			getValueB(hand[4])<<(4*0)

		bidsWithStr = append(bidsWithStr, [2]int{bid, strength})
	}

	slices.SortFunc(bidsWithStr, func(a, b [2]int) int {
		return cmp.Compare(a[1], b[1])
	})

	winnings := 0
	for rank, bid := range bidsWithStr {
		winnings += (rank + 1) * bid[0]
	}

	return winnings
}

func getValueB(card byte) int {
	switch card {
	case 'A':
		return 0xE
	case 'K':
		return 0xD
	case 'Q':
		return 0xC
	case 'J':
		return 0x1
	case 'T':
		return 0xA
	default:
		return int(card) - '0'
	}
}

func getTypeB(hand string) HandType {
	h := strings.Split(hand, "")
	slices.Sort(h)

	i := 0
	curr := h[0]
	var groups []int
	var wildcards int
	for _, card := range h {
		if card != curr {
			groups = append(groups, i)
			curr = card
			i = 0
		}
		if card == "J" {
			wildcards++
		} else {
			i++
		}
	}
	groups = append(groups, i)

	slices.Sort(groups)
	slices.Reverse(groups)

	if groups[0]+wildcards == 5 {
		return FiveOfAKind
	} else if groups[0]+wildcards == 4 {
		return FourOfAKind
	} else if groups[0]+groups[1]+wildcards == 5 {
		return FullHouse
	} else if groups[0]+wildcards == 3 {
		return ThreeOfAKind
	} else if groups[0]+groups[1]+wildcards == 4 {
		return TwoPair
	} else if groups[0]+wildcards == 2 {
		return OnePair
	} else {
		return HighCard
	}
}

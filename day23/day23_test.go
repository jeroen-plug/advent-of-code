package day23

import (
	"strings"
	"testing"
)

const example = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

func TestDay23a(t *testing.T) {
	want := 7
	res := day23a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day23a() = %d, want %d", res, want)
	}
}

func TestDay23b(t *testing.T) {
	want := "co,de,ka,ta"
	res := day23b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day23b() = %s, want %s", res, want)
	}
}

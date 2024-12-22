package day22

import (
	"slices"
	"testing"
)

func TestDay22a(t *testing.T) {
	want := 37327623
	res := day22a([]string{"1", "10", "100", "2024"})

	if res != want {
		t.Fatalf("day22a() = %d, want %d", res, want)
	}
}

func TestDay22b(t *testing.T) {
	want := 23
	res := day22b([]string{"1", "2", "3", "2024"}, 2)

	if res != want {
		t.Fatalf("day22b() = %d, want %d", res, want)
	}
}

func TestMonkeySecret(t *testing.T) {
	want := []int{15887950, 16495136, 527345, 704524, 1553684,
		12683156, 11100544, 12249484, 7753432, 5908254}
	var res []int

	var secret MonkeySecret
	secret.SetSeed(123)
	for range len(want) {
		secret.Next()
		res = append(res, int(secret))
	}

	if slices.Compare(res, want) != 0 {
		t.Fatalf("MonkeySecret.Next() = %d, want %d", res, want)
	}
}

func TestMonkeyPrice(t *testing.T) {
	want := []int{3, 0, 6, 5, 4, 4, 6, 4, 4, 2}
	var res []int

	var secret MonkeySecret
	secret.SetSeed(123)
	for range len(want) {
		res = append(res, secret.CurrentPrice())
		secret.Next()
	}

	if slices.Compare(res, want) != 0 {
		t.Fatalf("MonkeySecret.CurrentPrice() = %d, want %d", res, want)
	}
}

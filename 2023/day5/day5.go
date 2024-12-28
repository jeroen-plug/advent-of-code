package day5

import (
	"bufio"
	"log"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

type CategoryMapRange struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

type SeedRange struct {
	Start  int
	Lenght int
}

type Almanac struct {
	Seeds      []int
	SeedRanges []SeedRange

	SeedToSoil            []CategoryMapRange
	SoilToFertilizer      []CategoryMapRange
	FertilizerToWater     []CategoryMapRange
	WaterToLight          []CategoryMapRange
	LightToTemperature    []CategoryMapRange
	TemperatureToHumidity []CategoryMapRange
	HumidityToLocation    []CategoryMapRange
}

func Solution() (any, any) {
	f, err := os.Open(path.Join("input", "5.txt"))
	if err != nil {
		log.Fatal("Could not read input 5")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return day5a(lines), day5b(lines)
}

func day5a(lines []string) int {
	a := parse(lines)
	smallestLocation := -1

	for _, seed := range a.Seeds {
		soil := doMap(a.SeedToSoil, seed)
		fertilizer := doMap(a.SoilToFertilizer, soil)
		water := doMap(a.FertilizerToWater, fertilizer)
		light := doMap(a.WaterToLight, water)
		temperature := doMap(a.LightToTemperature, light)
		humidity := doMap(a.TemperatureToHumidity, temperature)
		location := doMap(a.HumidityToLocation, humidity)

		if smallestLocation == -1 || location < smallestLocation {
			smallestLocation = location
		}

		// fmt.Printf("seed %d, soil %d, fertilizer %d, water %d, light %d, temperature %d, humidity %d, location %d\n", seed, soil, fertilizer, water, light, temperature, humidity, location)
	}

	return smallestLocation
}

func day5b(lines []string) int {
	a := parse(lines)
	locations := make(chan int)
	var wg sync.WaitGroup

	for _, seedRange := range a.SeedRanges {
		wg.Add(1)
		go func() {
			defer wg.Done()
			smallestLocation := math.MaxInt
			for seed := seedRange.Start; seed < seedRange.Start+seedRange.Lenght; seed++ {
				soil := doMap(a.SeedToSoil, seed)
				fertilizer := doMap(a.SoilToFertilizer, soil)
				water := doMap(a.FertilizerToWater, fertilizer)
				light := doMap(a.WaterToLight, water)
				temperature := doMap(a.LightToTemperature, light)
				humidity := doMap(a.TemperatureToHumidity, temperature)
				location := doMap(a.HumidityToLocation, humidity)

				if location < smallestLocation {
					smallestLocation = location
				}
			}
			locations <- smallestLocation
		}()
	}

	go func() {
		wg.Wait()
		close(locations)
	}()

	smallestLocation := math.MaxInt
	for location := range locations {
		if location < smallestLocation {
			smallestLocation = location
		}
	}

	return smallestLocation
}

func doMap(ranges []CategoryMapRange, in int) int {
	for _, r := range ranges {
		if in >= r.SourceStart && in < r.SourceStart+r.Length {
			return in - r.SourceStart + r.DestinationStart
		}
	}
	return in
}

type CategoryMap int

const (
	NoCategoryMap CategoryMap = iota
	SeedToSoilMap
	SoilToFertilizerMap
	FertilizerToWaterMap
	WaterToLightMap
	LightToTemperatureMap
	TemperatureToHumidityMap
	HumidityToLocationMap
)

func parse(lines []string) Almanac {
	var a Almanac
	var currentMap CategoryMap

	for _, l := range lines {
		if strings.HasPrefix(l, "seeds: ") {
			var r SeedRange
			for i, seed := range strings.Split(strings.TrimPrefix(l, "seeds: "), " ") {
				n, err := strconv.Atoi(seed)
				if err == nil {
					a.Seeds = append(a.Seeds, n)
					if i%2 == 0 {
						r.Start = n
					} else {
						r.Lenght = n
						a.SeedRanges = append(a.SeedRanges, r)
					}
				}
			}
		} else if strings.HasSuffix(l, " map:") {
			switch strings.TrimSuffix(l, " map:") {
			case "seed-to-soil":
				currentMap = SeedToSoilMap
			case "soil-to-fertilizer":
				currentMap = SoilToFertilizerMap
			case "fertilizer-to-water":
				currentMap = FertilizerToWaterMap
			case "water-to-light":
				currentMap = WaterToLightMap
			case "light-to-temperature":
				currentMap = LightToTemperatureMap
			case "temperature-to-humidity":
				currentMap = TemperatureToHumidityMap
			case "humidity-to-location":
				currentMap = HumidityToLocationMap
			}
		} else if currentMap != NoCategoryMap {
			ns := strings.SplitN(l, " ", 3)
			if len(ns) < 3 {
				continue
			}
			n1, err := strconv.Atoi(ns[0])
			if err != nil {
				continue
			}
			n2, err := strconv.Atoi(ns[1])
			if err != nil {
				continue
			}
			n3, err := strconv.Atoi(ns[2])
			if err != nil {
				continue
			}

			r := CategoryMapRange{n1, n2, n3}
			switch currentMap {
			case SeedToSoilMap:
				a.SeedToSoil = append(a.SeedToSoil, r)
			case SoilToFertilizerMap:
				a.SoilToFertilizer = append(a.SoilToFertilizer, r)
			case FertilizerToWaterMap:
				a.FertilizerToWater = append(a.FertilizerToWater, r)
			case WaterToLightMap:
				a.WaterToLight = append(a.WaterToLight, r)
			case LightToTemperatureMap:
				a.LightToTemperature = append(a.LightToTemperature, r)
			case TemperatureToHumidityMap:
				a.TemperatureToHumidity = append(a.TemperatureToHumidity, r)
			case HumidityToLocationMap:
				a.HumidityToLocation = append(a.HumidityToLocation, r)
			}
		}
	}

	return a
}

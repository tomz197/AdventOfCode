package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Conversion struct {
	From  int
	To    int
	Range int
}

type Convmap = map[string][]Conversion

func (c *Conversion) convert(num int) (int, bool) {
	if c.From <= num && num < c.From+c.Range {
		return c.To + num - c.From, true
	}
	return num, false
}

func formatFile(file *os.File) (Convmap, []int) {
	// file:
	// seeds: 79 14 55 13
	//
	// seed-to-soil map:
	// 50 98 2
	// 52 50 48
	scanner := bufio.NewScanner(file)
	conversions := make(Convmap)
	var seeds []int

	var convs []Conversion
	var convName string
	var isReading bool = false
	for scanner.Scan() {
		line := scanner.Text()
		if len(seeds) == 0 {
			strSeeds := regexp.MustCompile("[0-9]+").FindAllString(line, -1)
			seeds = make([]int, 0, len(strSeeds))
			for _, strSeed := range strSeeds {
				seed, _ := strconv.Atoi(strSeed)
				seeds = append(seeds, seed)
			}
			continue
		}

		if line == "" {
			if isReading {
				conversions[convName] = convs
				convs = make([]Conversion, 0)
			}
			isReading = false
			continue
		}

		if !isReading {
			convName = line
			isReading = true
			continue
		}

		nums := regexp.MustCompile("[0-9]+").FindAllString(line, -1)
		if len(nums) != 3 {
			log.Fatal("invalid line: ", line)
		}
		newConv := Conversion{}
		newConv.To, _ = strconv.Atoi(nums[0])
		newConv.From, _ = strconv.Atoi(nums[1])
		newConv.Range, _ = strconv.Atoi(nums[2])

		convs = append(convs, newConv)
	}
	if isReading {
		conversions[convName] = convs
	}
	return conversions, seeds
}

func getLocation(num int, convMap *Convmap, path *[]string) int {
	for i, curr := range *path {
		if i == len(*path)-1 {
			continue
		}
		convName := curr + "-to-" + (*path)[i+1] + " map:"
		conversions := (*convMap)[convName]
		for _, conv := range conversions {
			temp, can := conv.convert(num)
			if can {
				num = temp
				break
			}
		}
	}
	return num
}

func main() {
	file, err := os.Open("day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	path := []string{
		"seed", "soil", "fertilizer", "water",
		"light", "temperature", "humidity", "location"}
	var conversions Convmap
	var seeds []int

	conversions, seeds = formatFile(file)

	var lowest1 int
	for i, num := range seeds {
		n := getLocation(num, &conversions, &path)
		if n < lowest1 || i == 0 {
			lowest1 = n
		}
	}
	fmt.Println("part1:", lowest1)

	var lowest2 int = math.MaxInt
	for i := 0; i < len(seeds); i = i + 2 {
		from := seeds[i]
		to := from + seeds[i+1]
		for j := from; j < to; j++ {
			n := getLocation(j, &conversions, &path)
			if n < lowest2 {
				lowest2 = n
				log.Print(j-from, "/", seeds[i+1])
			}
		}
		fmt.Println("-----------", i/2+1, "/", len(seeds)/2)
	}
	fmt.Println("part2:", lowest2)
}

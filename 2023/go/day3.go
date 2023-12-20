package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Number struct {
	y      int
	xStart int
	xEnd   int
	value  int
}

type Symbols = map[string]bool
type Gears = map[string]*[]Number

func getSymbolId(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func isAdjacent(num Number, symbols *Symbols) bool {
	// check if adjacent to a symbol
	from, to := num.xStart-1, num.xEnd+1
	for i := from; i <= to; i++ {
		if (*symbols)[getSymbolId(i, num.y-1)] {
			return true
		}
		if (*symbols)[getSymbolId(i, num.y+1)] {
			return true
		}
	}
	if (*symbols)[getSymbolId(from, num.y)] {
		return true
	}
	if (*symbols)[getSymbolId(to, num.y)] {
		return true
	}
	return false
}

func appendToGears(num Number, gears *Gears) {
	from, to := num.xStart-1, num.xEnd+1
	for i := from; i <= to; i++ {
		uVal := (*gears)[getSymbolId(i, num.y-1)]
		if uVal != nil {
			*uVal = append(*uVal, num)
		}
		dVal := (*gears)[getSymbolId(i, num.y+1)]
		if dVal != nil {
			*dVal = append(*dVal, num)
		}
	}
	lVal := (*gears)[getSymbolId(from, num.y)]
	if lVal != nil {
		*lVal = append(*lVal, num)
	}
	rVal := (*gears)[getSymbolId(to, num.y)]
	if rVal != nil {
		*rVal = append(*rVal, num)
	}
}

func handleLine(
	line string,
	lineNum int,
	numbers *[]Number,
	symbols *Symbols,
	gears *Gears,
) {
	// line = "8....+.58.""
	var num Number
	var found bool = false

	addNum := func(end int) {
		num.xEnd = end - 1
		val, err := strconv.Atoi(line[num.xStart:end])
		if err != nil {
			log.Fatal(err)
		}
		num.value = val
		*numbers = append(*numbers, num)
		num = Number{}
	}

	for i, val := range line {
		if val == '.' {
			if found {
				addNum(i)
				found = false
			}
			continue
		}
		if _, err := strconv.Atoi(line[i : i+1]); err == nil {
			if !found {
				num.xStart = i
				num.y = lineNum
				found = true
			}
		} else {
			if found {
				addNum(i)
				found = false
			}
			(*symbols)[getSymbolId(i, lineNum)] = true
			if val == '*' {
				(*gears)[getSymbolId(i, lineNum)] = &[]Number{}
			}
		}
	}
	if found {
		addNum(len(line))
	}
}

func main() {
	file, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	symbols := make(Symbols)
	gears := make(Gears)
	numbers := make([]Number, 0)

	var lineNum int
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(line, lineNum, &numbers, &symbols, &gears)
		lineNum++
	}

	// part 1
	filtered := make([]Number, 0)
	for _, num := range numbers {
		if num.value <= 0 {
			continue
		}
		if isAdjacent(num, &symbols) {
			filtered = append(filtered, num)
		}
	}

	var sum int
	for _, num := range filtered {
		sum += num.value
	}

	fmt.Println(sum)

	// part 2
	for _, num := range numbers {
		if num.value <= 0 {
			continue
		}
		appendToGears(num, &gears)
	}

	var sum2 int
	for _, gear := range gears {
		if len(*gear) == 2 {
			sum2 += (*gear)[0].value * (*gear)[1].value
		}
	}

	fmt.Println(sum2)
}

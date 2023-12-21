package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func formatLine(line string) ([]string, []string) {
	from := strings.Index(line, ":")
	separator := strings.Index(line, "|")

	winning := regexp.MustCompile("[0-9]+").
		FindAllString(line[from:separator], -1)
	guesses := regexp.MustCompile("[0-9]+").
		FindAllString(line[separator:], -1)

	return winning, guesses
}

func part1() {

}

func main() {
	file, err := os.Open("day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var scoreSum int
	var lines []string
	for scanner.Scan() {
		// line = "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"
		line := scanner.Text()
		lines = append(lines, line)
		winning, guesses := formatLine(line)

		var score int
		for _, win := range winning {
			for _, guess := range guesses {
				if guess == win {
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}
					break
				}
			}
		}

		scoreSum += score
	}

	fmt.Println(scoreSum)

	// part 2
	cardCount := make([]int, len(lines))
	for i := range cardCount {
		cardCount[i] = 1
	}

	for currLine, line := range lines {
		// line = "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"
		winning, guesses := formatLine(line)

		var score int
		for _, win := range winning {
			for _, guess := range guesses {
				if guess == win {
					score += 1
				}
			}
		}

		for i := 0; i < cardCount[currLine]; i++ {
			for j := currLine + 1; j <= currLine+score; j++ {
				if j >= len(lines) {
					break
				}
				cardCount[j] += 1
			}
		}
	}

	numOfCards := 0
	for _, count := range cardCount {
		numOfCards += count
	}

	fmt.Println(numOfCards)
}

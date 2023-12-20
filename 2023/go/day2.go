package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

type Game struct {
	id    int
	red   int
	green int
	blue  int
}

func (game *Game) isPossible(r int, g int, b int) bool {
	if r < game.red || g < game.green || b < game.blue {
		return false
	}
	return true
}

func getId(line string) (int, int, error) {
	// line = "Game 1:"
	lineLen := utf8.RuneCountInString(line)
	if lineLen < 7 {
		log.Fatal("Invalid line")
	}

	start := 5
	end := 6
	for i := 6; i < len(line); i++ {
		if line[i] == ':' {
			end = i
			break
		}
	}

	id, err := strconv.Atoi(line[start:end])
	if err != nil {
		log.Fatal(err)
	}
	return id, end + 2, nil
}

func handleRound(line string) (int, int, int) {
	// line = "3 blue, 4 red"
	lineLen := utf8.RuneCountInString(line)
	if lineLen < 4 {
		log.Fatal("Invalid line")
	}

	var red, green, blue int
	var value int
	var valStart int
	for i := 0; i < lineLen; i++ {
		switch line[i] {
		case ' ':
			num, err := strconv.Atoi(line[valStart:i])
			if err != nil {
				log.Fatal(err)
			}
			value = num
		case 'r':
			i += 4
			if value > red {
				red = value
			}
			valStart = i + 1
		case 'g':
			i += 6
			if value > green {
				green = value
			}
			valStart = i + 1
		case 'b':
			i += 5
			if value > blue {
				blue = value
			}
			valStart = i + 1
		}
	}
	return red, green, blue
}

func handleLine(line string) Game {
	// line = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	id, start, err := getId(line)
	if err != nil {
		log.Fatal(err)
	}
	var red, green, blue int

	for i := start; i < len(line); i++ {
		if line[i] != ';' && i != len(line)-1 {
			continue
		}
		nr, ng, nb := handleRound(line[start:i])
		if nr > red {
			red = nr
		}
		if ng > green {
			green = ng
		}
		if nb > blue {
			blue = nb
		}
	}

	return Game{id, red, green, blue}
}

func main() {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := make([]Game, 0)
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, handleLine(line))
	}

	sum := 0
	multSum := 0
	for _, game := range games {
		if game.isPossible(12, 13, 14) {
			sum += game.id
		}
		multSum += game.red * game.green * game.blue
	}
	fmt.Println("Sum of ids:", sum)
	fmt.Println("Sum of red*green*blue:", multSum)
}

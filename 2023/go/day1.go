package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := utf8.RuneCountInString(line)
		var firstNum, lastNum int

		for i := 0; i < lineLen; i++ {
			num, conv := strNumToInt(line[i:])
			if conv {
				firstNum = num
				break
			}
		}
		for i := lineLen - 1; i >= 0; i-- {
			num, conv := strNumToInt(line[i:])
			if conv {
				lastNum = num
				break
			}
		}

		sum += firstNum*10 + lastNum
	}

	fmt.Println(sum)
}

func strNumToInt(txt string) (int, bool) {
	strNums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	intNums := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	txtLen := utf8.RuneCountInString(txt)

	for i, _ := range strNums {
		if txt[0] == intNums[i] {
			return i + 1, true
		}

		strNumLen := utf8.RuneCountInString(strNums[i])
		if strNumLen > txtLen {
			continue
		}
		txtSlice := txt[:strNumLen]
		if txtSlice == strNums[i] {
			return i + 1, true
		}
	}
	return 0, false
}

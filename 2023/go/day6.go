package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func formatFile1(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	strTimes := regexp.MustCompile("[0-9]+").FindAllString(scanner.Text(), -1)
	scanner.Scan()
	strDistances := regexp.MustCompile("[0-9]+").FindAllString(scanner.Text(), -1)

	times := make([]int, 0, len(strTimes))
	distances := make([]int, 0, len(strDistances))

	for _, strTime := range strTimes {
		time, _ := strconv.Atoi(strTime)
		times = append(times, time)
	}
	for _, strDistance := range strDistances {
		distance, _ := strconv.Atoi(strDistance)
		distances = append(distances, distance)
	}

	return times, distances
}

func formatFile2(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	strTimes := regexp.MustCompile("[0-9]+").FindAllString(scanner.Text(), -1)
	scanner.Scan()
	strDistances := regexp.MustCompile("[0-9]+").FindAllString(scanner.Text(), -1)

	temp := ""
	for _, strTime := range strTimes {
		temp += strTime
	}
	time, _ := strconv.Atoi(temp)

	temp = ""
	for _, strDistance := range strDistances {
		temp += strDistance
	}
	distance, _ := strconv.Atoi(temp)

	return time, distance
}

func findPossibilities(time int, distance int) int {
	var possib int

	for i := 1; i < time; i++ {
		if (time-i)*i > distance {
			possib++
		}
	}

	return possib
}

func part1(file *os.File) int {
	times, distances := formatFile1(file)

	var res int = 1
	for i := 0; i < len(times); i++ {
		p := findPossibilities(times[i], distances[i])
		res *= p
	}

	return res
}

func part2(file *os.File) int {
	time, distance := formatFile2(file)

	res := findPossibilities(time, distance)

	return res
}

func main() {
	file1, err1 := os.Open("day6.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer file1.Close()

	fmt.Println("Part 1:", part1(file1))

	file2, err2 := os.Open("day6.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer file1.Close()

	fmt.Println("Part 2:", part2(file2))
}

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in format of a simple mathematical problem and its answer")

	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to Open the CSV file: %s", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided csv file.")
	}

	problems := parseLines(lines)

	correctAns := 0
	for i, p := range problems {
		fmt.Printf("Problem number %d: %s = \n", i+1, p.que)

		var answer string

		fmt.Scanf("%s\n", &answer)

		if answer == p.ans {
			correctAns++
		} else {
			fmt.Println("Incorrect 😭")
			break
		}
	}

	fmt.Printf("You scored %d out of %d", correctAns, len(problems))
}

func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			que: line[0],
			ans: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	que string
	ans string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

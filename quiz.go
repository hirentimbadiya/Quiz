package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in format of a simple mathematical problem and its answer")

	timeLimit := flag.Int("limit", 30, "This is the time limit to give answer in seconds")
	flag.Parse()

	// fmt.Println(time.Duration(*timeLimit) * time.Second)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	file, err := os.Open(*csvFileName)
	handleErr(err)

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	handleErr(err)

	problems := parseLines(lines)
	correctAns := 0

	for i, p := range problems {
		fmt.Printf("Problem number %d: %s = ", i+1, p.que)

		//* create an answer channel
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			//* pass the answer to channel whenever user enters answer
			answerCh <- answer
		}()

		//? The timer's channel will receive a value when the timer expires
		select {
		case <-timer.C:
			fmt.Println("\nTime Expired 😭")
			fmt.Printf("You scored %d out of %d", correctAns, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.ans {
				correctAns++
			} else {
				fmt.Println("\nIncorrect 😭")
				fmt.Printf("You scored %d out of %d", correctAns, len(problems))
				return
			}
		}
	}
}

func parseLines(lines [][]string) []problem {
	retValue := make([]problem, len(lines))

	for i, line := range lines {
		retValue[i] = problem{
			que: line[0],
			ans: strings.TrimSpace(line[1]),
		}
	}
	return retValue
}

type problem struct {
	que string
	ans string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

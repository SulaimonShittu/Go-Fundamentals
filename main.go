package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	scores := []score{}
	shouldContinue := true

	for shouldContinue {
		fmt.Println("1) Enter a score")
		fmt.Println("2) Print report")
		fmt.Println("q) Quit")
		fmt.Println()
		fmt.Println("Please select an option")

		var option string
		fmt.Scanln(&option)

		switch option {
		case "1":
			scores = append(scores, addStudent())
		case "2":
			printReport(scores)
		case "q":
			shouldContinue = false
		}
	}
}

type score struct {
	name  string
	score int
}

func printReport(scores []score) {
	fmt.Println()
	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))
	for _, v := range scores {
		fmt.Println(v.name, v.score)
	}
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println()
}

func addStudent() score {
	fmt.Println("Enter a student name and score")
	var name, rawScore string
	fmt.Scanln(&name, &rawScore)
	s, _ := strconv.Atoi(rawScore)
	return score{
		name:  name,
		score: s,
	}
}

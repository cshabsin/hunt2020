package main

import (
	"fmt"
)

var (
	clauses = []Clause{
		{0, 'E', 4},
		{1, 'J', 2},
		{2, 'L', 3},
		{3, 'D', 4},
		{4, 'N', 10},
		{5, 'G', 3},
		{6, 'K', 8},
		{7, 'A', 9},
		{8, 'O', 8},
		{9, 'I', 10},
		{10, 'C', 4},
		{11, 'B', 9},
		{12, 'F', 3},
		{13, 'H', 2},
		{14, 'M', 1},
	}
	answers = []Answer{
		{"Player Piano", "ALABAMA", "HEART", "OF", "DIXIE"},
		{"Wizard Woods", "BUREAU", "CHEST", "OF", "DRAWERS"},
		{"Pirate Ship", "CHARACTERFLAW", "FEET", "OF", "CLAY"},
		{"Magic 8-Ball", "CORNUCOPIA", "HORN", "OF", "PLENTY"},
		{"Magic Railway", "GRACEUNDERFIRE", "NERVES", "OF", "STEEL"},
		{"Espresso Stand", "HANGOVERREMEDY", "HAIR", "OFTHE", "DOG"},
		{"Monty Minotaur's Magical Menagerie", "LAUGHINGSTOCK", "BUTT", "OFTHE", "JOKE"},
		{"Checkerboard", "LEGALBALANCE", "SCALES", "OF", "JUSTICE"},
		{"Fortune Teller", "LOCALITY", "NECK", "OFTHE", "WOODS"},
		{"TT: Sand Witches", "THORNYISSUES", "BONE", "OF", "CONTENTION"},
		{
			puzzle: "Fortune Cookies",
			answer: "BA?",
			first:  "Z",
		},
		{
			puzzle: "The Wizard's Escape",
			answer: "CJ?",
			first:  "Z",
		},
		{
			puzzle: "Arts and Witchcrafts",
			answer: "CK?",
			first:  "Z",
		},
		{
			puzzle: "Drawing Board",
			answer: "K?",
			first:  "Z",
		},
		{
			puzzle: "Charming",
			answer: "Z?",
			first:  "Z",
		},
	}
)

func main() {
	ClauseBy(letter).Sort(clauses)
	//	AnswerBy(answer).Sort(answers)
	AnswerBy(expression).Sort(answers)
	zs := zip(answers, clauses)
	for _, z := range zs {
		fmt.Printf("%s: %v\n", z.IndexIntoExpression(), z)
	}
	for _, z := range zs {
		fmt.Print(z.IndexIntoExpression())
	}
	fmt.Print("\n")
}

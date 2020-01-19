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
		{"Fortune Cookies", "AMENDMENTS", "BILL", "OF", "RIGHTS"},
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
			puzzle: "The Wizard's Escape",
		},
		{
			puzzle: "Arts and Witchcrafts",
		},
		{
			puzzle: "Drawing Board",
		},
		{
			puzzle: "Charming",
		},
	}
	betweens = []string{
		"AA", "ALB", "BA", "CA", "CI", "D", "GZ", "I", "LB", "LF", "M", "Z",
	}
)

func out(zs []Zipped) string {
	var s string
	for _, z := range zs {
		s += z.IndexIntoExpression()
	}
	return s
}

func main() {
	for a := 0; a < len(betweens); a++ {
		for b := a; b < len(betweens); b++ {
			for c := b; c < len(betweens); c++ {
				for d := c; d < len(betweens); d++ {
						in := []string{
							betweens[a],
							betweens[b],
							betweens[c],
							betweens[d],
						}
						zs := handle(clauses, answers, in, false)
						fmt.Printf("%s | %s\n", out(zs), in)
				}
			}
		}
	}

}

func handle(clausesIn []Clause, answersIn []Answer, ins []string, print bool) []Zipped {
	var clauses []Clause
	for _, c := range clausesIn {
		clauses = append(clauses, Clause{c.order, c.letter, c.amount})
	}
	var answers []Answer
	for _, a := range answersIn {
		answers = append(answers, Answer{a.puzzle, a.answer, a.first, a.of, a.second})
	}
	for i := 0; i < len(ins); i++ {
		answers[i+11].answer = ins[i]
	}
	ClauseBy(letter).Sort(clauses)
	AnswerBy(answer).Sort(answers)
	//AnswerBy(expression).Sort(answers)
	zs := zip(answers, clauses)
	if print {
		for _, z := range zs {
			fmt.Printf("%s: %v\n", z.IndexIntoExpression(), z)
		}
		for _, z := range zs {
			fmt.Print(z.IndexIntoExpression())
		}
		fmt.Print("\n\n")
	}
	ZipBy(zorder).Sort(zs)
	return zs
}

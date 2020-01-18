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
			answer: "A?",
		},
		{
			puzzle: "The Wizard's Escape",
			answer: "F?",
		},
		{
			puzzle: "Arts and Witchcrafts",
			answer: "J?",
		},
		{
			puzzle: "Drawing Board",
			answer: "K?",
		},
		{
			puzzle: "Charming",
			answer: "Z?",
		},
	}
)

type Clause struct {
	order int
	letter byte
	amount int
}

type Answer struct {
	puzzle string
	answer string
	first  string
	of     string
	second string
}

func letter(c1, c2 *Clause) bool {
	return c1.letter < c2.letter
}

func answer(a1, a2 *Answer) bool {
	return a1.answer < a2.answer
}

type Zipped struct {
	Answer
	Clause
}

func zip(as []Answer, cs[]Clause) []Zipped {
	if len(as) != len(cs) {
		panic("whoa")
	}
	var zs []Zipped
	for i, a := range as {
		zs = append(zs, Zipped{a, cs[i]})
	}
	return zs
}

func (z Zipped) IndexIntoExpression() string {
	expression := z.first + z.of + z.second
	if z.amount > len(expression) {
		return "?"
	}
	
	return string(expression[z.amount-1])
}

func (z Zipped) String() string {
	return fmt.Sprintf("<%d %c %d | %s>", z.order, z.letter, z.amount, z.first+z.of+z.second)
}

func main() {
	ClauseBy(letter).Sort(clauses)
	AnswerBy(answer).Sort(answers)
	zs := zip(answers, clauses)
	fmt.Println(zs)
	for _, z := range zs {
		fmt.Print(z.IndexIntoExpression())
	}
	fmt.Print("\n")
}

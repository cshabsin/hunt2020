package main

import (
	"fmt"
	"sort"
)

type ClauseBy func(c1, c2 *Clause) bool

func (by ClauseBy) Sort(clauses []Clause) {
	csort := &clauseSorter{
		clauses: clauses,
		by:      by,
	}
	sort.Sort(csort)
}

type clauseSorter struct {
	clauses []Clause
	by      func(c1, c2 *Clause) bool
}

// Len is part of sort.Interface.
func (s *clauseSorter) Len() int {
	return len(s.clauses)
}

// Swap is part of sort.Interface.
func (s *clauseSorter) Swap(i, j int) {
	s.clauses[i], s.clauses[j] = s.clauses[j], s.clauses[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *clauseSorter) Less(i, j int) bool {
	return s.by(&s.clauses[i], &s.clauses[j])
}

type AnswerBy func(a1, a2 *Answer) bool

func (by AnswerBy) Sort(answers []Answer) {
	csort := &answerSorter{
		answers: answers,
		by:      by,
	}
	sort.Sort(csort)
}

type answerSorter struct {
	answers []Answer
	by      func(c1, c2 *Answer) bool
}

// Len is part of sort.Interface.
func (s *answerSorter) Len() int {
	return len(s.answers)
}

// Swap is part of sort.Interface.
func (s *answerSorter) Swap(i, j int) {
	s.answers[i], s.answers[j] = s.answers[j], s.answers[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *answerSorter) Less(i, j int) bool {
	return s.by(&s.answers[i], &s.answers[j])
}

type Clause struct {
	order  int
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

func (a Answer) expression() string {
	return a.first + a.of + a.second
}

func letter(c1, c2 *Clause) bool {
	return c1.letter < c2.letter
}

func answer(a1, a2 *Answer) bool {
	return a1.answer < a2.answer
}

func expression(a1, a2 *Answer) bool {
	return a1.expression() < a2.expression()
}

func zorder(z1, z2 *Zipped) bool {
	return z1.order < z2.order
}

type Zipped struct {
	Answer
	Clause
}

func zip(as []Answer, cs []Clause) []Zipped {
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
	if z.amount > len(z.expression()) {
		return "?"
	}

	return string(z.expression()[z.amount-1])
}

func (z Zipped) String() string {
	return fmt.Sprintf("<%2d %c %2d | %-20s %s>", z.order, z.letter, z.amount, z.answer, z.expression())
}

type ZipBy func(a1, a2 *Zipped) bool

func (by ZipBy) Sort(zips []Zipped) {
	csort := &zipSorter{
		zips: zips,
		by:   by,
	}
	sort.Sort(csort)
}

type zipSorter struct {
	zips []Zipped
	by   func(c1, c2 *Zipped) bool
}

// Len is part of sort.Interface.
func (s *zipSorter) Len() int {
	return len(s.zips)
}

// Swap is part of sort.Interface.
func (s *zipSorter) Swap(i, j int) {
	s.zips[i], s.zips[j] = s.zips[j], s.zips[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *zipSorter) Less(i, j int) bool {
	return s.by(&s.zips[i], &s.zips[j])
}


package main

import (
	"sort"
)

type ClauseBy func(c1, c2 *Clause) bool

func (by ClauseBy) Sort(clauses []Clause) {
	csort := &clauseSorter{
		clauses: clauses,
		by: by,
	}
	sort.Sort(csort)
}

type clauseSorter struct {
	clauses []Clause
	by func(c1, c2 *Clause) bool
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
		by: by,
	}
	sort.Sort(csort)
}

type answerSorter struct {
	answers []Answer
	by func(c1, c2 *Answer) bool
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


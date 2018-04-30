/*
 * Copyright (C) 2017-2018 Alibaba Group Holding Limited
 */
package cli

import "io"

const DefaultSuggestDistance = 2

func CalculateStringDistance(source string, target string) int {
	return DistanceForStrings([]rune(source), []rune(target), DefaultOptions)
}

// error with suggestions
type SuggestibleError interface {
	GetSuggestions() []string
}

func PrintSuggestions(w io.Writer, lang string, ss []string) {
	if len(ss) > 0 {
		Noticef(w, "\nDid you mean:\n")
		for _, s := range ss {
			Noticef(w, "  %s\n", s)
		}
	}
}

//
// helper class for Suggester
type Suggester struct {
	suggestFor string
	distance   int
	results    []string
}

func NewSuggester(v string, distance int) *Suggester {
	return &Suggester{
		suggestFor: v,
		distance:   distance,
	}
}

func (a *Suggester) Apply(s string) {
	d := CalculateStringDistance(a.suggestFor, s)
	if d <= a.distance {
		if d < a.distance {
			a.distance = d
			a.results = make([]string, 0)
		}
		a.results = append(a.results, s)
	}
}

func (a *Suggester) GetResults() []string {
	return a.results
}

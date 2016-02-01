package nlp

import (
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"regexp"
	"strings"

	log "github.com/avabot/ava/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/avabot/ava/Godeps/_workspace/src/github.com/dchest/stemmer/porter2"
)

// StructuredInput is generated by Ava and sent to packages as a helper tool.
type StructuredInput struct {
	Commands StringSlice
	Objects  StringSlice
	People   StringSlice
}

type WordClass struct {
	Word  string
	Class int
}

// SIT is a Structured Input Type. It corresponds to either a Command, Object,
// or Person
type SIT int

// Classifier is a set of common english word stems unique among their
// Structured Input Types. This enables extremely fast constant-time O(3)
// lookups of stems to their SITs with high accuracy and no training
// requirements. It consumes just a few MB in memory
type Classifier map[string]struct{}

func (c Classifier) ClassifyTokens(tokens []string) *StructuredInput {
	var s StructuredInput
	for _, t := range tokens {
		t = strings.ToLower(t)
		log.Debugln("checking", t)
		_, exists := c["C"+t]
		if exists {
			s.Commands = append(s.Commands, t)
		}
		_, exists = c["O"+t]
		if exists {
			s.Objects = append(s.Objects, t)
		}
		_, exists = c["P"+t]
		if exists {
			s.People = append(s.People, t)
		}
	}
	return &s
}

func TokenizeSentence(sent string) []string {
	tokens := []string{}
	for _, w := range strings.Fields(sent) {
		found := []int{}
		for i, r := range w {
			switch r {
			case '\'', '"', ',', '.', ':', ';', '!', '?':
				found = append(found, i)
			}
		}
		if len(found) == 0 {
			tokens = append(tokens, w)
			continue
		}
		for i, j := range found {
			tokens = append(tokens, w[:j-1])
			tokens = append(tokens, string(w[j]))
			if i+1 == len(found) {
				tokens = append(tokens, w[j+1:])
			}
		}
	}
	return tokens
}

func StemTokens(tokens []string) []string {
	eng := porter2.Stemmer
	stems := []string{}
	for _, w := range tokens {
		if len(w) == 1 {
			switch w {
			case "'", "\"", ",", ".", ":", ";", "!", "?":
				continue
			}
		}
		w = strings.ToLower(w)
		stems = append(stems, eng.Stem(w))
	}
	return stems
}

// StringSlice replaces []string, adding custom sql support for arrays in lieu
// of pq.
type StringSlice []string

// QuoteEscapeRegex replaces escaped quotes except if it is preceded by a
// literal backslash, e.g. "\\" should translate to a quoted element whose value
// is \
var QuoteEscapeRegex = regexp.MustCompile(`([^\\]([\\]{2})*)\\"`)

// Scan convert to a slice of strings
// http://www.postgresql.org/docs/9.1/static/arrays.html#ARRAYS-IO
func (s *StringSlice) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("scan source was not []bytes"))
	}
	str := string(asBytes)
	str = QuoteEscapeRegex.ReplaceAllString(str, `$1""`)
	str = strings.Replace(str, `\\`, `\`, -1)
	str = str[1 : len(str)-1]
	csvReader := csv.NewReader(strings.NewReader(str))
	slice, err := csvReader.Read()
	if err != nil && err.Error() != "EOF" {
		return err
	}
	*s = StringSlice(slice)
	return nil
}

func (s StringSlice) Value() (driver.Value, error) {
	// string escapes.
	// \ => \\\
	// " => \"
	for i, elem := range s {
		s[i] = `"` + strings.Replace(strings.Replace(elem, `\`, `\\\`, -1), `"`, `\"`, -1) + `"`
	}
	return "{" + strings.Join(s, ",") + "}", nil
}

// Last safely returns the last item in a StringSlice, which is most often the
// target of a pronoun, e.g. (In "Where is that?", "that" will most often refer
// to the last Object named in the previous sentence.
func (s StringSlice) Last() string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}

func (s StringSlice) String() string {
	if len(s) == 0 {
		return ""
	}
	var ss string
	for _, w := range s {
		ss += " " + w
	}
	return ss[1:]
}

func (s StringSlice) StringSlice() []string {
	ss := []string{}
	for _, tmp := range s {
		if len(tmp) <= 2 {
			continue
		}
		ss = append(ss, tmp)
	}
	return ss
}

var (
	ErrInvalidClass        = errors.New("invalid class")
	ErrInvalidOddParameter = errors.New("parameter count must be even")
	ErrSentenceTooShort    = errors.New("sentence too short to classify")
)

const (
	CommandI SIT = iota + 1
	PersonI
	ObjectI
)

var Pronouns map[string]SIT = map[string]SIT{
	"me":   PersonI,
	"us":   PersonI,
	"you":  PersonI,
	"him":  PersonI,
	"her":  PersonI,
	"them": PersonI,
	"it":   ObjectI,
	"that": ObjectI,
	// Ultimately Place and Time would be nice-to-have in a structured
	// input, but they don't outweigh the cost of training a full NER on
	// each new package
	// "there": PlaceI,
	// "then":  TimeI,
}

func max(slice []float64) float64 {
	if len(slice) == 0 {
		return 0.0
	}
	m := slice[0]
	for index := 1; index < len(slice); index++ {
		if slice[index] > m {
			m = slice[index]
		}
	}
	return m
}

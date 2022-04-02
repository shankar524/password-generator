package text

import (
	"fmt"
	"math/rand"
	"time"
)

type RuleType uint8

const (
	UPPERCASE RuleType = iota
	LOWERCASE
	SYMBOLS
	NUMBERS
)

const (
	UPPER_CHARS_SET     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWER_CHARS_SET     = "abcdefghijklmnopqrstuvwxyz"
	SPECIAL_CHARS_SET   = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
	NUMBERS_SET         = "0123456789"
	DEFAULT_UPPER_CHARS = 4
	DEFAULT_LOWER_CHARS = 3
)

type TextRule struct {
	For    RuleType
	Length int
}

type TextBuilder struct {
	Rules []TextRule
}

type text struct {
	rules []TextRule
}

type Builder interface {
	AddRule(TextRule) Builder
	Build() (Generator, error)
}

type Generator interface {
	Generate() string
}

var caseToCharacter = map[RuleType]string{
	UPPERCASE: UPPER_CHARS_SET,
	LOWERCASE: LOWER_CHARS_SET,
	SYMBOLS:   SPECIAL_CHARS_SET,
	NUMBERS:   NUMBERS_SET,
}

func (pb *TextBuilder) AddRule(rule TextRule) Builder {
	pb.Rules = append(pb.Rules, rule)
	return pb
}

func (tb *TextBuilder) Build() (t Generator, err error) {
	if len(tb.Rules) == 0 {
		defaultRule := []TextRule{
			{UPPERCASE, DEFAULT_UPPER_CHARS},
			{LOWERCASE, DEFAULT_LOWER_CHARS},
		}
		t = text{defaultRule}
		return
	}

	for _, rule := range tb.Rules {
		if rule.For > NUMBERS {
			err = fmt.Errorf("no such rule %d", rule.For)
			return
		}
		if rule.Length == 0 {
			err = fmt.Errorf("zero value for length is not allowed")
			return
		}
	}
	t = text{tb.Rules}
	return
}

func (t text) Generate() (text string) {
	seeder := rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)
	for _, rule := range t.rules {
		text += generateRandomTextFromSource(seeder, caseToCharacter[rule.For], rule.Length)
	}
	text = shuffle(seeder, text)
	return
}

func generateRandomTextFromSource(randomize *rand.Rand, sourceText string, length int) string {
	randomTexts := make([]byte, length)
	charsetLength := len(sourceText)
	for i := range randomTexts {
		randomNumber := randomize.Intn(charsetLength)
		randomTexts[i] = sourceText[randomNumber]
	}

	return string(randomTexts)
}

func shuffle(randomize *rand.Rand, textToShuffle string) string {
	strBytes := []rune(textToShuffle)
	randomize.Shuffle(len(strBytes), func(i, j int) {
		strBytes[i], strBytes[j] = strBytes[j], strBytes[i]
	})

	return string(strBytes)
}

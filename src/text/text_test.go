package text

import (
	"fmt"
	"strings"
	"testing"
)

func Test_TextBuilder_AddRule(t *testing.T) {
	t.Run("It appends new TextRule", func(t *testing.T) {
		tb := TextBuilder{}
		tb.AddRule(TextRule{UPPERCASE, 5}).AddRule(TextRule{LOWERCASE, 6})

		if len(tb.Rules) != 2 {
			t.Errorf("Expected Rules count to be %d but received %d", 2, len(tb.Rules))
		}
	})
}

func Test_TextBuilder_Build(t *testing.T) {
	t.Run("TextBuilder implements Builder interface", func(t *testing.T) {
		var _ Builder = new(TextBuilder)
	})
	t.Log("When no TextRule is configured")
	{
		expectedTextLength := DEFAULT_UPPER_CHARS + DEFAULT_LOWER_CHARS
		t.Run(fmt.Sprintf("it sets default text rule of length %d and does not return error", expectedTextLength), func(t *testing.T) {
			tb := TextBuilder{}
			textGen, err := tb.Build()
			if err != nil {
				t.Errorf("expected error to be %v but received '%s' error", nil, err.Error())
			}

			generatedText := textGen.Generate()
			actualTextLength := len(generatedText)
			if actualTextLength != expectedTextLength {
				t.Errorf("expected length of generated text %d but received %d instead", expectedTextLength, actualTextLength)
			}
		})
	}
	t.Log("When TextRule is configured")
	{
		t.Log(fmt.Sprintf("when For value of TextRule is other than %d, %d, %d and %d", UPPERCASE, LOWERCASE, SYMBOL, NUMBER))
		{
			t.Run("it should return error", func(t *testing.T) {
				tb := TextBuilder{}
				ruleFor := NUMBER + 1
				tb.AddRule(TextRule{RuleType(ruleFor), 2})
				_, err := tb.Build()
				if err == nil {
					t.Errorf("expected error to but did not receive")
				}

				expectedErrorMessage := fmt.Sprintf("no such rule %d", ruleFor)
				actualErrorMessage := err.Error()
				if actualErrorMessage != expectedErrorMessage {
					t.Errorf("expected error message: '%s'but received '%s'", expectedErrorMessage, actualErrorMessage)
				}
			})
		}
		t.Log("when Length value of TextRule is 0")
		{
			t.Run("it should return error", func(t *testing.T) {
				tb := TextBuilder{}
				tb.AddRule(TextRule{NUMBER, 0})
				_, err := tb.Build()
				if err == nil {
					t.Errorf("expected error to but did not receive")
				}

				expectedErrorMessage := "zero value for length is not allowed"
				actualErrorMessage := err.Error()
				if actualErrorMessage != expectedErrorMessage {
					t.Errorf("expected error message: '%s'but received '%s'", expectedErrorMessage, actualErrorMessage)
				}
			})
		}
		t.Log("When everything is fine")
		{
			t.Run("it generates text from provided TextRule and does not return error", func(t *testing.T) {
				tb := TextBuilder{}
				// Add 4 TextRule(s)
				tb.AddRule(TextRule{UPPERCASE, 2})
				tb.AddRule(TextRule{LOWERCASE, 3})
				tb.AddRule(TextRule{SYMBOL, 2})
				tb.AddRule(TextRule{NUMBER, 2})
				textGenI, err := tb.Build()
				if err != nil {
					t.Errorf("Expected error to be %v but received '%s' error", nil, err.Error())
				}

				expectedTextRuleCount := 4
				textGen := textGenI.(text)
				if len(tb.Rules) != expectedTextRuleCount {
					t.Errorf("Expected length of generated TextRules %d but received %d instead", expectedTextRuleCount, len(tb.Rules))
				}

				// returned text instance should have same text rule as TextBuilder rule
				if len(textGen.rules) != len(tb.Rules) {
					t.Errorf("Expected length of generated TextRule %d but received %d instead", len(tb.Rules), len(textGen.rules))
				}

				// check for TextRules equality
				for i, providedRules := range tb.Rules {
					assignedRule := textGen.rules[i]
					if (assignedRule.For != providedRules.For) || (assignedRule.Length != providedRules.Length) {
						t.Errorf("Expected TextRule: %+v received %+v instead", len(tb.Rules), len(textGen.rules))
					}
				}
			})
		}
	}
}

func Test_text_Generate(t *testing.T) {
	t.Run("text implements Generator interface", func(t *testing.T) {
		var _ Generator = new(text)
	})

	t.Log("When no TextRule was provided for TextBuilder")
	{
		expectedTextLength := DEFAULT_UPPER_CHARS + DEFAULT_LOWER_CHARS
		t.Run(fmt.Sprintf("it sets default TextRule(%d UPPER_CHARS and %d LOWER_CHARS) of  and does not return error", DEFAULT_UPPER_CHARS, DEFAULT_LOWER_CHARS), func(t *testing.T) {
			tb := TextBuilder{}
			textGen, err := tb.Build()
			if err != nil {
				t.Errorf("expected error to be %v but received '%s' error", nil, err.Error())
			}

			generatedText := textGen.Generate()
			actualTextLength := len(generatedText)
			if actualTextLength != expectedTextLength {
				t.Errorf("expected length of generated text %d but received %d instead", expectedTextLength, actualTextLength)
			}

			upperCaseCount := len(commonText(generatedText, caseToCharacter[UPPERCASE]))
			if upperCaseCount != DEFAULT_UPPER_CHARS {
				t.Errorf("expected %d upper characters but received %d instead", DEFAULT_UPPER_CHARS, upperCaseCount)
			}

			lowerCaseCount := len(commonText(generatedText, caseToCharacter[LOWERCASE]))
			if lowerCaseCount != DEFAULT_LOWER_CHARS {
				t.Errorf("expected %d lower characters but received %d instead", DEFAULT_UPPER_CHARS, upperCaseCount)
			}

			// returned text should not have symbol
			symbolCount := len(commonText(generatedText, caseToCharacter[SYMBOL]))
			if symbolCount != 0 {
				t.Errorf("expected %d symbol characters but received %d instead", 0, symbolCount)
			}

			// returned text should not have number
			numberCount := len(commonText(generatedText, caseToCharacter[NUMBER]))
			if numberCount != 0 {
				t.Errorf("expected %d numbered characters but received %d instead", 0, numberCount)
			}
		})
	}
	t.Log("When TextRule was provided for TextBuilder")
	{
		t.Run("it returns text based on provided text rule and doesnot return error", func(t *testing.T) {
			tb := TextBuilder{}
			// Add TextRule(s)
			tb.AddRule(TextRule{UPPERCASE, 2}) // 2 uppercase letters
			tb.AddRule(TextRule{LOWERCASE, 3}) // 3 lowercase letters
			tb.AddRule(TextRule{SYMBOL, 2})    // 2 symbol
			tb.AddRule(TextRule{NUMBER, 2})    // 2 number
			textGen, err := tb.Build()
			if err != nil {
				t.Errorf("expected error to be %v but received '%s' error", nil, err.Error())
			}

			generatedText := textGen.Generate()
			actualTextLength := len(generatedText)
			expectedTextLength := 2 + 3 + 2 + 2
			if actualTextLength != expectedTextLength {
				t.Errorf("expected length of generated text %d but received %d instead", expectedTextLength, actualTextLength)
			}

			upperCaseCount := len(commonText(generatedText, caseToCharacter[UPPERCASE]))
			if upperCaseCount != 2 {
				t.Errorf("expected %d upper characters but received %d instead", 2, upperCaseCount)
			}

			lowerCaseCount := len(commonText(generatedText, caseToCharacter[LOWERCASE]))
			if lowerCaseCount != 3 {
				t.Errorf("expected %d lower characters but received %d instead", 3, upperCaseCount)
			}

			symbolCount := len(commonText(generatedText, caseToCharacter[SYMBOL]))
			if symbolCount != 2 {
				t.Errorf("expected %d symbol characters but received %d instead", 2, symbolCount)
			}

			numberCount := len(commonText(generatedText, caseToCharacter[NUMBER]))
			if numberCount != 2 {
				t.Errorf("expected %d numbered characters but received %d instead", 2, numberCount)
			}
		})
	}
}

func commonText(firstString, secondString string) (common string) {
	for _, char := range firstString {
		stringifiedChar := string(char)
		if strings.Contains(secondString, stringifiedChar) {
			common += stringifiedChar
		}
	}
	return
}

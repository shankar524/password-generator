package cmd

import (
	"fmt"

	"github.com/shankar524/password-generator/src/text"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var (
	SymbolCountFlag    *int64
	UppercaseCountFlag *int64
	DowncaseCountFlag  *int64
	DigitsCountFlag    *int64
	CopyToClipboard    bool
	generateCmd        = &cobra.Command{
		Use:   "g",
		Short: "Generates random texts",
		Long:  `Generates random texts based on given text rule`,
		Run: func(cmd *cobra.Command, args []string) {
			var (
				symbolCount, upcaseCount, downcaseCount, numberCount int
			)
			textBuilder := text.TextBuilder{}

			// set default values and convert int64 to int
			symbolCount = int(*SymbolCountFlag)
			upcaseCount = int(*UppercaseCountFlag)
			downcaseCount = int(*DowncaseCountFlag)
			numberCount = int(*DigitsCountFlag)

			// add text rules
			addTextRule(&textBuilder, text.LOWERCASE, downcaseCount)
			addTextRule(&textBuilder, text.NUMBER, numberCount)
			addTextRule(&textBuilder, text.UPPERCASE, upcaseCount)
			addTextRule(&textBuilder, text.SYMBOL, symbolCount)

			textGenerator, err := textBuilder.Build()
			if err != nil {
				fmt.Printf("error creating text generator. Error: %s", err.Error())
				return
			}

			randomText := textGenerator.Generate()

			if !CopyToClipboard {
				fmt.Println(randomText)
				return
			}

			err = clipboard.Init()
			if err != nil {
				panic(err)
			}

			clipboard.Write(clipboard.FmtText, []byte(randomText))
		},
	}
)

func addTextRule(textBuilder *text.TextBuilder, textFor text.RuleType, length int) {
	if length > 0 {
		textBuilder.AddRule(text.TextRule{For: textFor, Length: length})
	}
}

func init() {
	// get flags
	SymbolCountFlag = generateCmd.Flags().Int64P("symbol", "s", 0, "number of symbol to generate in return text")
	UppercaseCountFlag = generateCmd.Flags().Int64P("up", "u", 0, "number of upper case letters to generate in return text")
	DowncaseCountFlag = generateCmd.Flags().Int64P("down", "d", 0, "number of down case letters to generate in return text")
	DigitsCountFlag = generateCmd.Flags().Int64P("number", "n", 0, "number of numerical digits to generate in return text")
	generateCmd.Flags().BoolVarP(&CopyToClipboard, "copy", "c", true, "copy generated text to clipboard, if set false then prints text in console itself")
	rootCmd.AddCommand(generateCmd)
}

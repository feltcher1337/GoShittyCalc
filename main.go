package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Knetic/govaluate"
	romnum "github.com/brandenc40/romannumeral"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var toReplace = strings.NewReplacer(
		"IV", "4",
		"VIII", "8",
		"VII", "7",
		"VI", "6",
		"V", "5",
		"IX", "9",
		"X", "10",
		"III", "3",
		"II", "2",
		"I", "1",

		"iv", "4",
		"viii", "8",
		"vii", "7",
		"vi", "6",
		"v", "5",
		"ix", "9",
		"x", "10",
		"iii", "3",
		"ii", "2",
		"i", "1",
	)

	re1, _ := regexp.Compile(`^ *([1-9]|10) *[+*/-] *([1-9]|10) *\r?\n?$`)
	re2, _ := regexp.Compile(`(?i)^ *(I|II|III|IV|V|VI|VII|VIII|IX|X) *[+*/-] *(I|II|III|IV|V|VI|VII|VIII|IX|X) *\r?\n?$`)

	fmt.Println("Enter an expression:")
	userInput, _ := reader.ReadString('\n')

	switch {
	case re1.MatchString(userInput):
		fmt.Println(eval(userInput))
	case re2.MatchString(userInput):
		userInput = toReplace.Replace(userInput)
		resultInt := eval(userInput)
		if resultInt > 0 {
			result, _ := romnum.IntToString(resultInt)
			fmt.Println(result)
		} else {
			fmt.Println("There are no negative numbers or zero in the Roman number system.")
		}
	default:
		fmt.Println("Invalid format. There should be 2 operands from 1 to 10 in Roman or Arabic number system and one operator between them. You cannot mix Roman and Arabic.")
	}
}

// Evaluate expression
func eval(str string) int {
	exp, _ := govaluate.NewEvaluableExpression(str)
	resultInterface, _ := exp.Eval(nil)
	resultFloat := resultInterface.(float64)
	result := int(resultFloat)
	return result
}

package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// in go charcter is a type alias for int32 also known as rune
// rune is a type alias for int32
// rune represents utf-32 code point
// rune is used to represent a unicode code point

type TextAnalyzer struct {
	charCount      int
	lineCount      int
	upperCaseCount int
	lowerCaseCount int
}

func NewTextAnalyzer() *TextAnalyzer {
	return &TextAnalyzer{}
}

func (ta *TextAnalyzer) Analyze(text string) {
	// each element in the string is of type rune
	ta.charCount = utf8.RuneCountInString(text)
	ta.lineCount = utf8.RuneCountInString(text) - utf8.RuneCountInString(strings.ReplaceAll(text, "\n", ""))

	for _, r := range text {
		if unicode.IsUpper(r) {
			ta.upperCaseCount++
		} else if unicode.IsLower(r) {
			ta.lowerCaseCount++
		}
	}
}

func (ta *TextAnalyzer) Print() {
	fmt.Printf("Character count: %d\n", ta.charCount)
	fmt.Printf("Line count: %d\n", ta.lineCount)
	fmt.Printf("Uppercase count: %d\n", ta.upperCaseCount)
	fmt.Printf("Lowercase count: %d\n", ta.lowerCaseCount)
}

func main() {
	// defining a variable of type rune
	var r rune = 'a'
	// printing the value of the variable
	fmt.Println(r) // outut :: 97

	s1 := "Hello äöü"
	// converting the string to a slice of runes
	// range keyword can be used to iterate over a string
	r1Slice := []rune(s1)
	// printing the slice of runes
	fmt.Println(s1) // output :: Hello äöü
	// prints its value in int32 that is its unicode code point
	fmt.Println(r1Slice) // output :: [72 101 108 108 111 32 228 246 252]

	// iterating over string
	// each element in the string is of type rune
	for i, r := range s1 {
		fmt.Printf("%d %c %T\n", i, r, r)

	}
	// output ::0 H int32
	// 1 e int32
	// 2 l int32
	// 3 l int32
	// 4 o int32
	// 5   int32
	// 6 ä int32
	// 8 ö int32
	// 10 ü int32

	for _, r := range s1 {
		str := string(r)
		byteCount := utf8.RuneLen(r)

		fmt.Printf("Rune : %c UTF8 %q %d\n", r, str, byteCount)

	}
	//
	// output ::
	// 	Rune : H UTF8 "H" 1
	// Rune : e UTF8 "e" 1
	// Rune : l UTF8 "l" 1
	// Rune : l UTF8 "l" 1
	// Rune : o UTF8 "o" 1
	// Rune :   UTF8 " " 1
	// Rune : ä UTF8 "ä" 2
	// Rune : ö UTF8 "ö" 2
	// Rune : ü UTF8 "ü" 2

	b := `German has 30 letters, including the standard 26 of the Latin alphabet,
	plus the umlauts (ä, ö, ü), and the eszett (ß), also known as sharp s`

	ta := NewTextAnalyzer()
	ta.Analyze(b)
	ta.Print()
}

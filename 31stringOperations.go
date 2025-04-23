package main

import (
	"fmt"
	s "strings"
)

// alias fmt.Println to a shorter name as we’ll use it a lot below.

var p = fmt.Println

func main() {
	//Here’s a sample of the functions available in strings.
	// Since these are functions from the package, not methods on the string object itself, we need pass the string in question as the first argument to the function. You can find more functions in the strings package docs.

	p("Contains:  ", s.Contains("test", "es"))        // output :: Contains:  true
	p("Count:     ", s.Count("test", "t"))            // output :: Count:     2
	p("HasPrefix: ", s.HasPrefix("test", "te"))       // output :: HasPrefix:  true
	p("HasSuffix: ", s.HasSuffix("test", "st"))       // output :: HasSuffix:  true
	p("Index:     ", s.Index("test", "e"))            // output :: Index:     1
	p("Join:      ", s.Join([]string{"a", "b"}, "-")) // output :: Join:      a-b
	p("Repeat:    ", s.Repeat("a", 5))                // output :: Repeat:    aaaaa
	p("Replace:   ", s.Replace("foo", "o", "0", -1))  // output :: Replace:   f00
	p("Replace:   ", s.Replace("foo", "o", "0", 1))   // output :: Replace:   f0o
	p("Split:     ", s.Split("a-b-c-d-e", "-"))       // output :: Split:     [a b c d e]
	p("ToLower:   ", s.ToLower("TEST"))               // output :: ToLower:   test
	p("ToUpper:   ", s.ToUpper("test"))               // output :: ToUpper:   TEST
}

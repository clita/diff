package main

import (
	"fmt"
	"github.com/clita/diff"
)

func main() {
	a := "Hello Vaibhav\n\n\n how are you doing tnight, you must be doing good"
	b := "Hello Vaibhav how are you doing night, you must be doing  good"

	// string1, string2 := diff.FindColouredChanges(a, b, "lines")
	string1, string2 := diff.FindColouredChanges(a, b, "words")
	fmt.Println("String1: \n" + string1 + "\n\n" + "String2: \n" + string2)
}

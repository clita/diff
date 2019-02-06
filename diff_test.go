package diff

import "testing"

func TestFindColouredChangesLines(t *testing.T) {
	tables := []struct {
		input1  string
		input2  string
		output1 string
		output2 string
	}{
		{"Hello\n World", "Hello", "Hello\n[- World](fg-red,fg-bold)", "Hello"},
		{"Hello", "Hello World", "[-Hello](fg-red,fg-bold)", "[+Hello World](fg-green,fg-bold)"},
		{"Hello\n\nWorld", "Hello\nWorld", "Hello\n[-](fg-red,fg-bold)\nWorld", "Hello\nWorld"},
	}

	for _, table := range tables {

		output1, output2 := FindColouredChanges(table.input1, table.input2, "lines", false)
		if (output1 != table.output1) || (output2 != table.output2) {
			t.Errorf("FindColouredChanges (by lines) of (%s & %s) was incorrect, got: %s & %s, want: %s & %s.", table.input1, table.input2, output1, output2, table.output1, table.output2)
		}
	}
}

func TestFindColouredChangesWords(t *testing.T) {
	tables := []struct {
		input1  string
		input2  string
		output1 string
		output2 string
	}{
		{"Hello\n World", "Hello", "Hello\n [World](fg-red,fg-bold)", "Hello"},
		{"Hello", "Hello World", "Hello", "Hello [World](fg-green,fg-bold)"},
		{"World", "Hello\nWorld", "World", "[Hello](fg-green,fg-bold)\nWorld"},
	}

	for _, table := range tables {

		output1, output2 := FindColouredChanges(table.input1, table.input2, "words", false)
		if output1 != table.output1 || output2 != table.output2 {
			t.Errorf("FindColouredChanges (by lines) of (%s & %s) was incorrect, got: %s & %s, want: %s & %s.", table.input1, table.input2, output1, output2, table.output1, table.output2)
		}
	}
}

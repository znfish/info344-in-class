package testing

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		name         string
		input        string
		exepectedOut string
	}{ //initializer
		{
			name:         "empty string",
			input:        "",
			exepectedOut: "",
		},
		{
			name:         "single character",
			input:        "a",
			exepectedOut: "a",
		},
		{
			name:         "two characters",
			input:        "ab",
			exepectedOut: "ba",
		},
		{
			name:         "stressed",
			input:        "stressed",
			exepectedOut: "desserts",
		},
		{
			name:         "high ascii",
			input:        "Hello, 世界",
			exepectedOut: "界世 ,olleH",
		},
	}

	for _, c := range cases {
		if output := Reverse(c.input); output != c.exepectedOut {
			t.Errorf("%s: got %s bit expected %s", c.name, output, c.exepectedOut)
		}
	}
}

func TestGetGreeting(t *testing.T) {
	cases := []struct {
		name         string
		input        string
		exepectedOut string
	}{
		{
			name:         "empty",
			input:        "",
			exepectedOut: "Hello, World!",
		},
		{
			name:         "Name = Christy",
			input:        "Christy",
			exepectedOut: "Hello, Christy!",
		},
	}

	for _, c := range cases {
		if output := GetGreeting(c.input); output != c.exepectedOut {
			t.Errorf("%s: got %s but expected %s", c.name, output, c.exepectedOut)
		}
	}
}

func TestParseSize(t *testing.T) {
	cases := []struct {
		name         string
		input        string
		exepectedOut *Size
	}{
		{
			name:         "empty string",
			input:        "",
			exepectedOut: &Size{},
		},
	}

	for _, c := range cases {
		if output := ParseSize(c.input); output.Height !=
			c.exepectedOut.Height || output.Width != c.exepectedOut.Width {
			t.Errorf("%s: got %v but expected %v", c.name, output, c.exepectedOut)
		}
	}
}

func TestLateDaysConsume(t *testing.T) {
	ld := NewLateDays()
	for i := 3; i > -10; i-- {
		expectedLateDays := i
		if expectedLateDays < 0 {
			expectedLateDays = 0
		}
		if output := ld.Consume("test"); output != expectedLateDays {
			t.Errorf("iteration %d: got %d but expected %d", i, output, expectedLateDays)
		}
	}
}

package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    charmander  bulbasaur \nsquirtle  ",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		{
			input:    "charmander bulbasaur squirtle",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		{
			input:    "CHARmander BULbaSAUR squIRtle",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
	}
	for _, c := range cases {
		result := cleanInput(c.input)
		if len(result) != len(c.expected) {
			t.Errorf("Result did not contain correct number of inputs")
		}
	}

}

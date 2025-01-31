package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "Winter     Is     Comming!",
			expected: []string{"winter", "is", "comming!"},
		},
		{
			input:    "Szeth Son Son Vallano",
			expected: []string{"szeth", "son", "son", "vallano"},
		},
		{
			input:    "Extremely-Super_Duper_Long_Input;Thing",
			expected: []string{"extremely-super_duper_long_input;thing"},
		},
		{
			input:    "               ",
			expected: []string{},
		},
	}

	for _, tc := range cases {
		t.Run("cleanInput", func(t *testing.T) {
			result := cleanInput(tc.input)
			t.Log(result, tc.expected)

			if len(result) != len(tc.expected) {
				t.Errorf("Result length doesn't match the expected.\nResult: %d\nExpected: %d\n", len(result), len(tc.expected))
			}

			for i, word := range result {
				if word != tc.expected[i] {
					t.Errorf("Result doesn't match the expected!\nResult: %+v\nExpected: %+v\n", result, tc.expected)
				}
			}
		})
	}

}

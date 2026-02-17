package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "This Should Become All Lowercase",
			expected: []string{"this", "should", "become", "all", "lowercase"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Output did not match expected: got %s, expected %s", word, expectedWord)
			}
		}
	}
}

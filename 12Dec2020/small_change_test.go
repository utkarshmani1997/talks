package main

import (
	"fmt"
	"testing"
)

func TestSmallChange(t *testing.T) {
	cases := []struct {
		name           string
		input1, input2 string
		expected       string
		isErr          bool
	}{
		{name: "Case01", input1: "Small", input2: "change", expected: "Small change"},
		{name: "Case02", input1: "", input2: "hello", expected: "", isErr: true},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("Test case: %s, function: small_change(%s, %s)", tt.name, tt.input1, tt.input2), func(t *testing.T) {
			actual, err := smallChange(tt.input1, tt.input2)
			if actual != tt.expected {
				t.Errorf("expected %s but got %s", tt.expected, actual)
			}

			if tt.isErr && err == nil {
				t.Errorf("expected err, but got no err")
			}
		})
	}
}

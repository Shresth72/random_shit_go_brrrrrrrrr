package decode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeBencodedValue(t *testing.T) {
	tests := []struct {
		input         string
		expected      interface{}
		expectedError bool
	}{
		// Integer parsing
		{"i54e", 54, false},
		{"isde", 0, true},
		{"i0e", 0, false},
		{"345e", 0, true},
		{"i00e", 0, true},
		{"i1a2e", 0, true},
		{"i123ee", 0, true},
		{"i-0e", 0, true},
		{"ie", 0, true},
		{"i54", 0, true},
		{"i054e", 0, true},
		{"i-10e", -10, false},
		{"i54 e", 0, true},
		{"i541235e235235", 541235, false},
		// Strings parsing
		{"4:spam", "spam", false},
		{"0:", "", false},
		{"7:Hello69", "Hello69", false},
		{"", "", true},
		{"5:Help", "", true},
		{"5 :Hello", "", true},
		{"0:extra", "", false},
		{"abc:abc", "", true},
		{"4spam", "", true},
	}

	for _, test := range tests {
		res, err := DecodeBencodedValue(test.input)
		if test.expectedError {
			assert.Error(t, err, "Expected error for input: %s and result: %v", test.input, res)
		} else {
			assert.NoError(t, err, "Unexpected error for input: %s", test.input)
			assert.Equal(t, test.expected, res, "Mismatch for input: %s", test.input)
		}
	}
}

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
		{"4:spam", "spam", false},
		{"0:", "", false},
		{"7:Hello69", "Hello69", false},
		{"", "", true},
		{"5:Help", "", true},
		{"abc:abc", "", true},
		{"4spam", "", true},
	}

	for _, test := range tests {
		res, err := DecodeBencodedValue(test.input)
		if test.expectedError {
			assert.Error(t, err, "Expected error for input: %s", test.input)
		} else {
			assert.NoError(t, err, "Unexpected error for input: %s", test.input)
			assert.Equal(t, test.expected, res, "Mismatch for input: %s", test.input)
		}
	}
}

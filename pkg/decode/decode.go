package decode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func DecodeBencodedValue(encodedValue string) (interface{}, string, error) {
	if len(encodedValue) == 0 {
		return nil, "", fmt.Errorf("Empty encoded value")
	}

	switch encodedValue[0] {
	case 'i':
		return decodeInteger(encodedValue)
	case 'l':
		return decodeList(encodedValue)
	case 'd':
		return decodeDict(encodedValue)
	default:
		if unicode.IsDigit(rune(encodedValue[0])) {
			return decodeString(encodedValue)
		}
	}

	return nil, "", fmt.Errorf("Unhandled encoded value: %s", encodedValue)
}

func decodeInteger(encodedValue string) (interface{}, string, error) {
	if encodedValue[0] != 'i' || !strings.Contains(encodedValue, "e") {
		return nil, "", fmt.Errorf("Invalid integer encoding")
	}

	encodedValue = encodedValue[1:]
	parts := strings.SplitN(encodedValue, "e", 2)
	if len(parts) != 2 || parts[0] == "" {
		return nil, "", fmt.Errorf("Invalid integer encoding")
	}

	numStr := parts[0]
	if !isValidNumber(numStr) {
		return nil, "", fmt.Errorf("Invalid integer encoding: leading zeroes are not allowed")
	}

	decoded, err := strconv.Atoi(numStr)
	if err != nil {
		return nil, "", err
	}
	return decoded, parts[1], err
}

func decodeString(encodedValue string) (interface{}, string, error) {
	splits := strings.SplitN(encodedValue, ":", 2)
	if len(splits) != 2 {
		return nil, "", fmt.Errorf("Malformed bencode string")
	}

	lenStr := splits[0]
	if !isValidNumber(lenStr) {
		return nil, "", fmt.Errorf("Invalid integer encoding: leading zeroes are not allowed")
	}

	strLen, err := strconv.Atoi(lenStr)
	if err != nil {
		return nil, "", err
	}

	if len(splits[1]) < strLen {
		return nil, "", fmt.Errorf("Not enough characters for the string")
	}
	decodedString := splits[1][:strLen]
	return decodedString, splits[1][strLen:], nil
}

func decodeList(encodedValue string) (interface{}, string, error) {
	if encodedValue[0] != 'l' {
		return nil, "", fmt.Errorf("Invalid list encoding")
	}

	var values []interface{}
	rest := encodedValue[1:]

	for len(rest) > 0 && rest[0] != 'e' {
		value, remainder, err := DecodeBencodedValue(rest)
		if err != nil {
			return nil, "", err
		}
		values = append(values, value)
		rest = remainder
	}

	if len(rest) == 0 || rest[0] != 'e' {
		return nil, "", fmt.Errorf("list not terminated correctly")
	}
	return values, rest[1:], nil
}

func decodeDict(encodedValue string) (interface{}, string, error) {
	if encodedValue[0] != 'd' {
		return nil, "", fmt.Errorf("Invalid dict encoding")
	}

	values := make(map[string]interface{})
	rest := encodedValue[1:]

	for len(rest) > 0 && rest[0] != 'e' {
		keyRaw, remainder, err := DecodeBencodedValue(rest)
		if err != nil {
			return nil, "", err
		}

		keyStr, ok := keyRaw.(string)
		if !ok {
			return nil, "", fmt.Errorf("Key must be a string")
		}

		val, remainder, err := DecodeBencodedValue(remainder)
		if err != nil {
			return nil, "", err
		}

		values[keyStr] = val
		rest = remainder
	}

	if len(rest) == 0 || rest[0] != 'e' {
		return nil, "", fmt.Errorf("list not terminated correctly")
	}
	return values, rest[1:], nil
}

func isValidNumber(numStr string) bool {
	if len(numStr) == 0 {
		return false
	}
	if numStr[0] == '0' && len(numStr) > 1 {
		return false
	}
	if strings.HasPrefix(numStr, "-0") {
		return false
	}
	return true
}

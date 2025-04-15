package decode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func DecodeBencodedValue(encodedValue string) (interface{}, error) {
	if len(encodedValue) == 0 {
		return nil, fmt.Errorf("Empty encoded value")
	}

	switch encodedValue[0] {
	case 'i':
		return decodeInteger(encodedValue)
	default:
		if unicode.IsDigit(rune(encodedValue[0])) {
			return decodeString(encodedValue)
		}
	}

	return nil, fmt.Errorf("Unhandled encoded value: %s", encodedValue)
}

func decodeInteger(encodedValue string) (int, error) {
	if !strings.Contains(encodedValue, "e") {
		return 0, fmt.Errorf("Invalid integer encoding: missing 'e' terminator")
	}

	encodedValue = encodedValue[1:]
	decoded, err := strconv.Atoi(strings.SplitN(encodedValue, "e", 2)[0])
	if err != nil {
		return 0, err
	}

	return decoded, nil
}

func decodeString(encodedValue string) (string, error) {
	splits := strings.SplitN(encodedValue, ":", 2)
	if len(splits) != 2 {
		return "", fmt.Errorf("Malformed bencode string")
	}

	strLen, err := strconv.Atoi(splits[0])
	if err != nil {
		return "", err
	}

	if len(splits[1]) < strLen {
		return "", fmt.Errorf("Not enough characters for the string")
	}
	decodedString := splits[1][:strLen]
	return decodedString, nil
}

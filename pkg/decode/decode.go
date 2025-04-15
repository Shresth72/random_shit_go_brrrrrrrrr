package decode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func DecodeBencodedValue(encodedValue string) (interface{}, error) {
	if len(encodedValue) > 0 && unicode.IsDigit(rune(encodedValue[0])) {
		splits := strings.SplitN(encodedValue, ":", 2)
		if len(splits) != 2 {
			return nil, fmt.Errorf("Malformed bencode string")
		}

		strLen, err := strconv.Atoi(splits[0])
		if err != nil {
			return nil, err
		}

		if len(splits[1]) < strLen {
			return nil, fmt.Errorf("Not enough characters for the string")
		}
		decodedString := splits[1][:strLen]
		return decodedString, nil
	}
	return nil, fmt.Errorf("Unhandled encoded value: %s", encodedValue)
}

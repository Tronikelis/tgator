package utils

import "unicode"

func HasUppercase(str string) bool {
	for _, r := range str {
		if unicode.IsUpper(r) {
			return true
		}
	}

	return false
}

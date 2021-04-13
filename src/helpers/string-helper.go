package helpers

import "strings"

func IsEmptyOrWhiteSpace(stringToValidate string) bool {
	stringWithoutWhiteSpace := strings.TrimSpace(stringToValidate)

	return len(stringWithoutWhiteSpace) == 0
}

package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetSuccessWhenStringIsFullOfWhiteSpace(t *testing.T) {
	stringFullOfWhiteSpace := "         "

	result := IsEmptyOrWhiteSpace(stringFullOfWhiteSpace)

	assert.True(t, result)
}

func TestShouldGetSuccessWhenStringIsEmpty(t *testing.T) {
	emptyString := ""

	result := IsEmptyOrWhiteSpace(emptyString)

	assert.True(t, result)
}

func TestShouldGetFalseWhenStringHasSomething(t *testing.T) {
	emptyString := "VC"

	result := IsEmptyOrWhiteSpace(emptyString)

	assert.False(t, result)
}

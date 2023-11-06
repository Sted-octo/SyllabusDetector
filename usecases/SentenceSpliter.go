package usecases

import (
	"errors"
	"strings"
)

const BAD_ARGUMENTS_ERROR string = "bad arguments"

func split(sentence string) ([]string, error) {
	trimmed := strings.TrimSpace(sentence)
	if trimmed == "" {
		return nil, errors.New(BAD_ARGUMENTS_ERROR)
	}

	return strings.FieldsFunc(trimmed, delimiters), nil
}

func delimiters(r rune) bool {
	return r == ' ' || r == ',' || r == ';'
}

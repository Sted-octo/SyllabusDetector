package usecases

import "errors"

const BAD_ARGUMENTS_ERROR string = "bad arguments"

func split(sentence *string) ([]string, error) {
	if sentence == nil {
		return nil, errors.New(BAD_ARGUMENTS_ERROR)
	}
	return nil, nil
}

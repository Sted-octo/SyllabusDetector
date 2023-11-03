package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Nil_Input_Sentence_Should_Return_Error_Bad_Argument(t *testing.T) {
	var sentence *string = nil

	_, err := split(sentence)

	assert.EqualError(t, err, BAD_ARGUMENTS_ERROR)
}

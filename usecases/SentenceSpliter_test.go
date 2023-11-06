package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Empty_Input_Sentence_Should_Return_Error_Bad_Argument(t *testing.T) {
	var sentence string = ""

	_, err := split(sentence)

	assert.EqualError(t, err, BAD_ARGUMENTS_ERROR)
}

func Test_Non_Empty_Input_Sentence_without_space_Should_Return_String_List_Of_length_1(t *testing.T) {
	var sentence string = "Steeve"

	words, err := split(sentence)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(words))
	assert.Equal(t, "Steeve", words[0])
}

func Test_Sentence_with_2_words_and_space_Should_Return_String_List_Of_length_2(t *testing.T) {
	var sentence string = "Steeve Dess"

	words, err := split(sentence)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(words))
	assert.Equal(t, "Steeve", words[0])
	assert.Equal(t, "Dess", words[1])
}

func Test_Sentence_with_2_words_and_comas_Should_Return_String_List_Of_length_2(t *testing.T) {
	var sentence string = "Steeve, Dess"

	words, err := split(sentence)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(words))
	assert.Equal(t, "Steeve", words[0])
	assert.Equal(t, "Dess", words[1])
}

func Test_Sentence_with_2_words_and_semicolon_Should_Return_String_List_Of_length_2(t *testing.T) {
	var sentence string = "Steeve; Dess"

	words, err := split(sentence)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(words))
	assert.Equal(t, "Steeve", words[0])
	assert.Equal(t, "Dess", words[1])
}

func Test_Sentence_with_3_words_and_spaces_Should_Return_String_List_Of_length_3(t *testing.T) {
	var sentence string = "Steeve Elfie Lise"

	words, err := split(sentence)

	assert.Nil(t, err)
	assert.Equal(t, 3, len(words))
	assert.Equal(t, "Steeve", words[0])
	assert.Equal(t, "Elfie", words[1])
	assert.Equal(t, "Lise", words[2])
}

func Test_One_Word_Sentence_ending_with_space_and_Cariage_Return_Should_Return_String_List_Of_length_1(t *testing.T) {
	var sentence string = "Steeve \r\n"

	words, err := split(sentence)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(words))
	assert.Equal(t, "Steeve", words[0])
}

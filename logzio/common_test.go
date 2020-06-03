package logzio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripWhitespaceReturnsDesiredResult(t *testing.T) {
	input := "test string with	tabs"
	expected := "teststringwithtabs"

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}

func TestStripWhitespaceReturnsDesiredResultWhenInputStringIsEmpty(t *testing.T) {
	input := ""
	expected := ""

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}

func TestStripWhitespaceReturnsDesiredResultWhenInputStringIsJustSpaces(t *testing.T) {
	input := "        "
	expected := ""

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}

func TestStripWhitespaceReturnsDesiredResultWhenInputStringIsJustTabs(t *testing.T) {
	input := "				"
	expected := ""

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}

func TestStripWhitespaceReturnsDesiredResultWhenInputStringIsJustCharacters(t *testing.T) {
	input := "abcdefghijklmnoplqrstuvwxyz"
	expected := "abcdefghijklmnoplqrstuvwxyz"

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}

func TestStripWhitespaceReturnsDesiredResultWhenInputStringHasWhiteSpaceAtFrontAndBack(t *testing.T) {
	input := "     text with whitespace at front and back    	"
	expected := "textwithwhitespaceatfrontandback"

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}

func TestStripWhitespaceReturnsDesiredResultWhenInputStringHasCarraigeReturn(t *testing.T) {
	input := `this string
	spans multiple
lines`
	expected := "thisstringspansmultiplelines"

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}

func TestDiffJsonReturnsTrueWhenStructurallyTheSame(t *testing.T) {
	old := "{\"a\":{\"c\": []}, \"b\": true}"
	new := "{\"b\": true, \"a\":{\"c\": []}}"

	assert.True(t, jsonEqual(old, new))
}

func TestDiffJsonReturnsNotEqualWhenJsonIsDifferent(t *testing.T) {
	old := "{\"a\":{\"c\": []}, \"b\": true}"
	new := "{\"b\": true, \"a\":{\"d\": []}}"

	assert.False(t, jsonEqual(old, new))
}

func TestDiffJsonReturnsTrueWhenStringsTheSameAndJson(t *testing.T) {
	old := "{}"
	new := "{}"

	assert.True(t, jsonEqual(old, new))
}

func TestDiffJsonReturnsTrueWhenStringsTheSameButNotJson(t *testing.T) {
	old := "{a"
	new := "{a"

	assert.True(t, jsonEqual(old, new))
}

func TestDiffJsonReturnsTrueWithExcessWhiteSpace(t *testing.T) {
	old := "{\"a\":{\"c\": []}, \"b\": true}"
	new := "{\"b\": true, \"a\":{\"c\": []}}"

	assert.True(t, jsonEqual(old, new))
}

func TestDiffJsonReturnsTrueWhenAreTheSame(t *testing.T) {
	old := "{\"a\":{\"c\": []}, \"b\": true}"
	new := "{\"a\":{\"c\": []}, \"b\": true}"

	assert.True(t, jsonEqual(old, new))
}

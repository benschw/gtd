package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	// given
	expected := &Request{
		Action:  "a",
		Context: "@work",
		Tags:    []string{"#foo", "#bar"},
		Subject: "Hello World",
	}
	args := []string{"a", "@work", "#foo", "#bar", "Hello", "World"}

	// when
	r, err := ParseArgs(args, "@foo")

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", r), "parsed request doesn't match")
	//assert.True(t, reflect.DeepEqual(expected, r), fmt.Sprintf("\n   %+v\n!= %+v", expected, r))
}

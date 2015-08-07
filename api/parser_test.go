package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	// given
	expected := &Request{
		Action:  "+",
		Context: "@work",
		Tags:    []string{"#foo", "#bar"},
		Subject: "Hello World",
	}
	args := []string{"+", "@work", "#foo", "#bar", "Hello", "World"}

	// when
	r, err := ParseArgs(args, "@foo")

	// then
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expected, r), fmt.Sprintf("%+v != %+v", expected, r))
}

package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Scenario struct {
	R *Request
	A []string
}

func TestParseArgsAdd(t *testing.T) {
	// given
	scenarios := []*Scenario{
		&Scenario{
			R: &Request{
				Action:  "a",
				Context: "@work",
				Tags:    []string{"#foo", "#bar"},
				Subject: "Hello World",
			},
			A: []string{"a", "@work", "#foo", "#bar", "Hello", "World"},
		},
	}

	for _, s := range scenarios {
		// when
		found, err := ParseArgs(s.A, "@foo")

		// then
		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprintf("%+v", s.R), fmt.Sprintf("%+v", found), "parsed request doesn't match")
	}
}
func TestParseArgsModify(t *testing.T) {
	// given
	scenarios := []*Scenario{
		&Scenario{
			R: &Request{
				Action:       "m",
				Id:           "123",
				Context:      "@test",
				Tags:         []string{"#foo", "#bar"},
				TagsToRemove: []string{"#baz"},
				Subject:      "Hello World",
			},
			A: []string{"m", "123", "@test", "#foo", "-#baz", "#bar", "Hello", "World"},
		},
		&Scenario{
			R: &Request{
				Action:       "m",
				Id:           "123",
				TagsToRemove: []string{"#baz"},
			},
			A: []string{"m", "123", "-#baz"},
		},
		&Scenario{
			R: &Request{
				Action:  "m",
				Id:      "123",
				Context: "@test",
			},
			A: []string{"m", "123", "@test"},
		},
		&Scenario{
			R: &Request{
				Action:  "m",
				Id:      "123",
				Subject: "Hello World",
			},
			A: []string{"m", "123", "Hello", "World"},
		},
	}

	for _, s := range scenarios {
		// when
		found, err := ParseArgs(s.A, "@foo")

		// then
		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprintf("%+v", s.R), fmt.Sprintf("%+v", found), "parsed request doesn't match")
	}
}
func TestParseArgsList(t *testing.T) {
	// given
	scenarios := []*Scenario{
		&Scenario{
			R: &Request{
				Action:  "l",
				Context: "@work",
				Tags:    []string{"#foo", "#bar"},
			},
			A: []string{"l", "@work", "#foo", "#bar"},
		},
	}

	for _, s := range scenarios {
		// when
		found, err := ParseArgs(s.A, "@foo")

		// then
		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprintf("%+v", s.R), fmt.Sprintf("%+v", found), "parsed request doesn't match")
	}
}

package gtd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// given
	r := NewMemRepo()

	todo := &Todo{
		Id: "0",
		Meta: &Meta{
			Context: "@work",
			Tags:    []string{"#foo", "#bar"},
		},
		Subject: "Hello World",
	}
	req := &Request{
		Action:  "+",
		Context: todo.Meta.Context,
		Tags:    todo.Meta.Tags,
		Subject: todo.Subject,
	}

	// when
	_, err := Dispatch(req, r)

	// then

	assert.Nil(t, err)
	assert.Equal(t, todo.String(), r.Get("0").String(), "should be equal")
}
func TestList(t *testing.T) {
	// given
	r := NewMemRepo()

	todo := &Todo{
		Id: "0",
		Meta: &Meta{
			Context: "@work",
			Tags:    []string{"#foo", "#bar"},
		},
		Subject: "Hello World",
	}
	todo2 := &Todo{
		Id: "1",
		Meta: &Meta{
			Context: "@home",
			Tags:    []string{"#foo", "#bar"},
		},
		Subject: "Hello Galaxy",
	}
	r.Save(todo)
	r.Save(todo2)

	// when
	out, err := Dispatch(&Request{Action: "l", Context: "@work"}, r)

	// then

	assert.Nil(t, err)
	assert.Equal(t, todo.String()+"\n", out)
}
func TestClose(t *testing.T) {
	// given
	r := NewMemRepo()

	todo := &Todo{
		Id:     "0",
		Status: StatusOpen,
	}
	r.Save(todo)

	req := &Request{
		Action: "-",
		Id:     "0",
	}

	// when
	out, err := Dispatch(req, r)

	// then

	assert.Nil(t, err)
	assert.Equal(t, "0", out, "should output deleted id")
	assert.Equal(t, StatusClosed, r.Get("0").Status, "Status should be closed")
}

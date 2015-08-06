package gtd

import "testing"

func TestAdd(t *testing.T) {
	// given
	r := NewMemRepo()
	d := &Dispatcher{Repo: r}

	args := []string{"+", "@work", "#foo", "hello", "world"}

	// when
	_, err := d.Dispatch(&Meta{}, args)

	// then

	if err != nil {
		t.Error(err)
	}

	if r.Get("0").Subject != "hello world" {
		t.Errorf("got '%+v'", r.Get("0"))
	}
}
func TestList(t *testing.T) {
	// given
	r := NewMemRepo()
	d := &Dispatcher{Repo: r}

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

	args := []string{"list", "@work"}

	// when
	out, err := d.Dispatch(&Meta{}, args)

	// then

	if err != nil {
		t.Error(err)
	}

	if (todo.String() + "\n") != out {
		t.Errorf("got '%s'", out)
	}
}

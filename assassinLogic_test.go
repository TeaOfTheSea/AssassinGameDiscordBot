package assassinlogic

import (
	"container/list"
	"errors"
	"fmt"
	"testing"
)

func TestStringToLL(t *testing.T) {
	t.Run("Passing an empty string and checking for error", func(t *testing.T) {
		_, got := StringToLL("")
		want := errors.New("No input given.")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Passing a single user", func(t *testing.T) {
		// Getting output and checking for errors before output
		got, err := StringToLL("Colin")
		if err != nil {
			t.Errorf(fmt.Sprint(err))
		}
		want := list.New()
		want.PushFront("Colin")
		eGot := got.Front()
		for eWant := want.Front(); eWant != nil; eWant = eWant.Next() {
			if eGot == nil {
				t.Errorf("list got storter than list want")
				break
			}
			if eGot.Value != eWant.Value {
				t.Errorf("In linked list, got %q want %q", eGot.Value, eWant.Value)
			}
			eGot = eGot.Next()
		}
	})
}

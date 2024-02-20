package assassinlogic

import (
	"container/list"
	"errors"
	"fmt"
	"testing"
)

func TestLLToString(t *testing.T) {
	t.Run("Passing an empty list and checking for error", func(t *testing.T) {
		input := list.New()
		_, got := LLToString(input)
		want := errors.New("Input list is empty")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

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
		compareLists(got, want, t)
	})
	t.Run("Passing two users", func(t *testing.T) {
		got, err := StringToLL("Colin -> Tan10o")
		if err != nil {
			t.Errorf(fmt.Sprint(err))
		}
		want := list.New()
		want.PushFront("Colin")
		want.PushFront("Tan10o")
		compareLists(got, want, t)
	})
}

func compareLists(got, want *list.List, t *testing.T) {
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
}

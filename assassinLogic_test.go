package assassinlogic

import (
	"container/list"
	"errors"
	"fmt"
	"strings"
	"testing"
)

/*
############################################################

              LL / String Conversion

############################################################
*/

/*
##############################
Function Tests
##############################
*/

func TestStringToLL(t *testing.T) {
	t.Run("Passing an empty string and checking for error", func(t *testing.T) {
		_, got := StringToLL("")
		want := errors.New("Input string is empty")
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
		want.PushBack("Colin")
		err = compareLists(got, want, t)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Passing two users", func(t *testing.T) {
		got, err := StringToLL("Colin -> Tan10o")
		if err != nil {
			t.Errorf(fmt.Sprint(err))
		}
		want := list.New()
		want.PushBack("Colin")
		want.PushBack("Tan10o")
		err = compareLists(got, want, t)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Converting from string to list and back", func(t *testing.T) {
		wants := [3]string{
			"Colin",
			"Colin -> Tan10o",
			"Colin -> Tan10o -> Dr. Bob",
		}
		for _, want := range wants {
			list, err := StringToLL(want)
			if err != nil {
				t.Error(err)
			}
			got, err := LLToString(list)
			if err != nil {
				t.Error(err)
			}
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		}
	})
}

func TestLLToString(t *testing.T) {
	t.Run("Passing an empty list and checking for error", func(t *testing.T) {
		input := list.New()
		_, got := LLToString(input)
		want := errors.New("Input list is empty")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Passing a list of one element", func(t *testing.T) {
		input := list.New()
		input.PushBack("Colin")
		got, err := LLToString(input)
		if err != nil {
			t.Error(err)
		}
		want := "Colin"
		if strings.Compare(got, want) != 0 {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Passing a lits of one element", func(t *testing.T) {
		input := list.New()
		input.PushBack("Colin")
		input.PushBack("Tan10o")
		got, err := LLToString(input)
		if err != nil {
			t.Error(err)
		}
		want := "Colin -> Tan10o"
		if strings.Compare(got, want) != 0 {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

/*
############################################################

                      Helper Functions

############################################################
*/

/*
##############################
Linked Lists
##############################
*/

func compareLists(got, want *list.List, t *testing.T) error {
	eGot := got.Front()
	for eWant := want.Front(); eWant != nil; eWant = eWant.Next() {
		if eGot == nil {
			t.Errorf("list got storter than list want")
			break
		}
		if eGot.Value != eWant.Value {
			return errors.New(fmt.Sprintf("In linked list, got %q want %q", eGot.Value, eWant.Value))
		}
		eGot = eGot.Next()
	}
	return nil
}

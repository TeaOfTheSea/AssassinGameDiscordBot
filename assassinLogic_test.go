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
		err = compareLists(got, want)
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
		err = compareLists(got, want)
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

                    Linked List Operations

############################################################
*/

func TestBuildLL(t *testing.T) {
	t.Run("Passing empty string slice", func(t *testing.T) {
		_, got := BuildLL([]string{})
		want := errors.New("Input slice is empty")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Passing in one name", func(t *testing.T) {
		got, err := BuildLL([]string{"Walter"})
		if err != nil {
			t.Error(err)
		}
		want := list.New()
		want.PushBack("Walter")
		err = compareLists(got, want)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Passing in two names", func(t *testing.T) {
		got, err := BuildLL([]string{"Walter", "Colin"})
		if err != nil {
			t.Error(err)
		}
		want1 := list.New()
		want1.PushBack("Walter")
		want1.PushBack("Colin")
		want2 := list.New()
		want2.PushBack("Colin")
		want2.PushBack("Walter")
		// The function being tested here returns in a random
		// order, so we're happy as long as one of the errors
		// turns up nil
		if compareLists(got, want1) != nil && compareLists(got, want2) != nil {
			t.Errorf("Got an output which mached neither want permutation.")
		}
	})
	t.Run("Checking output list as a string using LLToString", func(t *testing.T) {
		list, err := BuildLL([]string{"Walter", "Colin", "Tan10o"})
		if err != nil {
			t.Error(err)
		}
		got, err := LLToString(list)
		if err != nil {
			t.Error(err)
		}
		wants := []string{"Walter -> Colin -> Tan10o",
			"Walter -> Tan10o -> Colin",
			"Colin -> Walter -> Tan10o",
			"Colin -> Tan10o -> Walter",
			"Tan10o -> Walter -> Colin",
			"Tan10o -> Colin -> Walter"}
		// BuildLL generates a list in a random order, which
		// is kept when LLToString is called. This means our
		// functions were successful as long as we find a match
		// in one of the above strings.
		found := 0
		for _, want := range wants {
			if got == want {
				found++
			}
		}
		if found != 1 {
			t.Errorf("%d wants found got %q", found, got)
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

func compareLists(got, want *list.List) error {
	eGot := got.Front()
	for eWant := want.Front(); eWant != nil; eWant = eWant.Next() {
		if eGot == nil {
			return errors.New("List got shorter than list want")
		}
		if eGot.Value != eWant.Value {
			return errors.New(fmt.Sprintf("In linked list, got %q want %q", eGot.Value, eWant.Value))
		}
		eGot = eGot.Next()
	}
	return nil
}

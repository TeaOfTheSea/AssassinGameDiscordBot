package assassinlogic

import (
	"container/list"
	"errors"
	"fmt"
	"slices"
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
		_, err = compareLists(got, want)
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
		_, err = compareLists(got, want)
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
		_, got := BuildLL([]string{"Walter"})
		want := errors.New("Input slice has only one element")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("got %q want %q", got, want)
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
		order1, _ := compareLists(got[0], want1)
		order2, _ := compareLists(got[0], want2)
		if order1 == false && order2 == false {
			t.Errorf("Got an output which mached neither want permutation.")
		}
	})
	t.Run("Checking output list as a string using LLToString", func(t *testing.T) {
		lists, err := BuildLL([]string{"Walter", "Colin", "Tan10o"})
		if err != nil {
			t.Error(err)
		}
		got, err := LLToString(lists[0])
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

		gotInWant := slices.Index(wants, got)
		if gotInWant == -1 {
			t.Errorf("Index %d found. Got %q wanted:\n%v", gotInWant, got, wants)
		}
	})
	t.Run("Using previous method to check outputs for four inputs", func(t *testing.T) {
		lists, err := BuildLL([]string{"Walter", "Colin", "Tan10o", "Waldo"})
		if err != nil {
			t.Error(err)
		}

		// if len(lists) == 1 {
		// 	got, err := LLToString(lists[0])
		// 	if err != nil {
		// 		t.Error(err)
		// 	}
		// 	wants := []string{"Walter -> Colin -> Tan10o -> Waldo",
		// 		"Walter -> Colin -> Waldo -> Tan10o",
		// 		"Walter -> Tan10o -> Colin -> Waldo",
		// 		"Walter -> Tan10o -> Waldo -> Colin",
		// 		"Walter -> Waldo -> Colin -> Tan10o",
		// 		"Walter -> Waldo -> Tan10o -> Colin",
		// 		"Colin -> Walter -> Tan10o -> Waldo",
		// 		"Colin -> Walter -> Waldo -> Tan10o",
		// 		"Colin -> Tan10o -> Walter -> Waldo",
		// 		"Colin -> Tan10o -> Waldo -> Walter",
		// 		"Colin -> Waldo -> Walter -> Tan10o",
		// 		"Colin -> Waldo -> Tan10o -> Walter",
		// 		"Tan10o -> Walter -> Colin -> Waldo",
		// 		"Tan10o -> Walter -> Waldo -> Colin",
		// 		"Tan10o -> Colin -> Walter -> Waldo",
		// 		"Tan10o -> Colin -> Waldo -> Walter",
		// 		"Tan10o -> Waldo -> Walter -> Colin",
		// 		"Tan10o -> Waldo -> Colin -> Walter",
		// 		"Waldo -> Walter -> Colin -> Tan10o",
		// 		"Waldo -> Walter -> Tan10o -> Colin",
		// 		"Waldo -> Colin -> Walter -> Tan10o",
		// 		"Waldo -> Colin -> Tan10o -> Walter",
		// 		"Waldo -> Tan10o -> Walter -> Colin",
		// 		"Waldo -> Tan10o -> Colin -> Walter"}
		// 	// BuildLL generates a list in a random order, which
		// 	// is kept when LLToString is called. This means our
		// 	// functions were successful as long as we find a match
		// 	// in one of the above strings.
		// 	found := 0
		// 	for _, want := range wants {
		// 		if got == want {
		// 			found++
		// 		}
		// 	}
		// 	if found != 1 {
		// 		errorString := ""
		// 		for i := range lists {
		// 			for e := lists[i].Front(); e != nil; e = e.Next() {
		// 				errorString += fmt.Sprint(e.Value)
		// 				if e.Next() != nil {
		// 					errorString += " -> "
		// 				}
		// 			}
		// 			if i+1 < len(lists) {
		// 				errorString += "\n"
		// 			}
		// 		}
		// 		t.Errorf("%d found got %q chains:\n%q", found, got, errorString)
		// 	}
		// } else {
		// 	var gots [2]string
		// 	gots[0], err = LLToString(lists[0])
		// 	if err != nil {
		// 		t.Error(err)
		// 	}
		// 	gots[1], err = LLToString(lists[1])
		// 	if err != nil {
		// 		t.Error(err)
		// 	}
		// 	wants := [][]string{{"Walter -> Colin", "Tan10o -> Waldo"},
		// 		{"Walter -> Colin", "Waldo -> Tan10o"},
		// 		{"Walter -> Tan10o", "Colin -> Waldo"},
		// 		{"Walter -> Tan10o", "Waldo -> Colin"},
		// 		{"Walter -> Waldo", "Colin -> Tan10o"},
		// 		{"Walter -> Waldo", "Tan10o -> Colin"},
		// 		{"Colin -> Walter", "Tan10o -> Waldo"},
		// 		{"Colin -> Walter", "Waldo -> Tan10o"},
		// 		{"Colin -> Tan10o", "Walter -> Waldo"},
		// 		{"Colin -> Tan10o", "Waldo -> Walter"},
		// 		{"Colin -> Waldo", "Walter -> Tan10o"},
		// 		{"Colin -> Waldo", "Tan10o -> Walter"},
		// 		{"Tan10o -> Walter", "Colin -> Waldo"},
		// 		{"Tan10o -> Walter", "Waldo -> Colin"},
		// 		{"Tan10o -> Colin", "Walter -> Waldo"},
		// 		{"Tan10o -> Colin", "Waldo -> Walter"},
		// 		{"Tan10o -> Waldo", "Walter -> Colin"},
		// 		{"Tan10o -> Waldo", "Colin -> Walter"},
		// 		{"Waldo -> Walter", "Colin -> Tan10o"},
		// 		{"Waldo -> Walter", "Tan10o -> Colin"},
		// 		{"Waldo -> Colin", "Walter -> Tan10o"},
		// 		{"Waldo -> Colin", "Tan10o -> Walter"},
		// 		{"Waldo -> Tan10o", "Walter -> Colin"},
		// 		{"Waldo -> Tan10o", "Colin -> Walter"},
		// 	}
		// 	found := 0
		// 	for _, v := range wants {
		// 		matches := 0
		// 		for i := range v {
		// 			if v[i] == gots[i] {
		// 				matches += 1
		// 			}
		// 		}
		// 		if matches == 2 {
		// 			found += 1
		// 		}
		// 	}
		// }
		wants := [][]string{{"Walter -> Colin -> Tan10o -> Waldo"},
			{"Walter -> Colin -> Waldo -> Tan10o"},
			{"Walter -> Tan10o -> Colin -> Waldo"},
			{"Walter -> Tan10o -> Waldo -> Colin"},
			{"Walter -> Waldo -> Colin -> Tan10o"},
			{"Walter -> Waldo -> Tan10o -> Colin"},
			{"Colin -> Walter -> Tan10o -> Waldo"},
			{"Colin -> Walter -> Waldo -> Tan10o"},
			{"Colin -> Tan10o -> Walter -> Waldo"},
			{"Colin -> Tan10o -> Waldo -> Walter"},
			{"Colin -> Waldo -> Walter -> Tan10o"},
			{"Colin -> Waldo -> Tan10o -> Walter"},
			{"Tan10o -> Walter -> Colin -> Waldo"},
			{"Tan10o -> Walter -> Waldo -> Colin"},
			{"Tan10o -> Colin -> Walter -> Waldo"},
			{"Tan10o -> Colin -> Waldo -> Walter"},
			{"Tan10o -> Waldo -> Walter -> Colin"},
			{"Tan10o -> Waldo -> Colin -> Walter"},
			{"Waldo -> Walter -> Colin -> Tan10o"},
			{"Waldo -> Walter -> Tan10o -> Colin"},
			{"Waldo -> Colin -> Walter -> Tan10o"},
			{"Waldo -> Colin -> Tan10o -> Walter"},
			{"Waldo -> Tan10o -> Walter -> Colin"},
			{"Waldo -> Tan10o -> Colin -> Walter"},
			{"Walter -> Colin", "Tan10o -> Waldo"},
			{"Walter -> Colin", "Waldo -> Tan10o"},
			{"Walter -> Tan10o", "Colin -> Waldo"},
			{"Walter -> Tan10o", "Waldo -> Colin"},
			{"Walter -> Waldo", "Colin -> Tan10o"},
			{"Walter -> Waldo", "Tan10o -> Colin"},
			{"Colin -> Walter", "Tan10o -> Waldo"},
			{"Colin -> Walter", "Waldo -> Tan10o"},
			{"Colin -> Tan10o", "Walter -> Waldo"},
			{"Colin -> Tan10o", "Waldo -> Walter"},
			{"Colin -> Waldo", "Walter -> Tan10o"},
			{"Colin -> Waldo", "Tan10o -> Walter"},
			{"Tan10o -> Walter", "Colin -> Waldo"},
			{"Tan10o -> Walter", "Waldo -> Colin"},
			{"Tan10o -> Colin", "Walter -> Waldo"},
			{"Tan10o -> Colin", "Waldo -> Walter"},
			{"Tan10o -> Waldo", "Walter -> Colin"},
			{"Tan10o -> Waldo", "Colin -> Walter"},
			{"Waldo -> Walter", "Colin -> Tan10o"},
			{"Waldo -> Walter", "Tan10o -> Colin"},
			{"Waldo -> Colin", "Walter -> Tan10o"},
			{"Waldo -> Colin", "Tan10o -> Walter"},
			{"Waldo -> Tan10o", "Walter -> Colin"},
			{"Waldo -> Tan10o", "Colin -> Walter"}}

		got := make([]string, len(lists))
		for i := range got {
			got[i], err = LLToString(lists[i])
			if err != nil {
				t.Errorf(fmt.Sprint(err))
			}
		}

		found := false
		for _, want := range wants {
			if slices.Equal(want, got) {
				found = true
			}
		}
		if found == false {
			t.Errorf("Got %v", got)
		}
	})
}

/*
############################################################

                  Testing Helper Functions

############################################################
*/

/*
##############################
Linked Lists
##############################
*/

func compareLists(got, want *list.List) (bool, error) {
	eGot := got.Front()
	for eWant := want.Front(); eWant != nil; eWant = eWant.Next() {
		if eGot == nil {
			return false, errors.New("List got shorter than list want")
		}
		if eGot.Value != eWant.Value {
			return false, errors.New(fmt.Sprintf("In linked list, got %q want %q", eGot.Value, eWant.Value))
		}
		eGot = eGot.Next()
	}
	return true, nil
}

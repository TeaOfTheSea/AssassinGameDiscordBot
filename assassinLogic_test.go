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

                    	 Game Operations

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

		got := make([]string, len(lists))
		for i := range got {
			got[i], err = LLToString(lists[i])
			if err != nil {
				t.Errorf(fmt.Sprint(err))
			}
		}

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

func TestPlayerKilled(t *testing.T) {
	t.Run("Chains slice is nil", func(t *testing.T) {
		var nilSlice []*list.List
		_, _, _, got := PlayerKilled(nilSlice, "DeadPlayer")
		want := errors.New("Input slice is nil")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("Chains slice is empty", func(t *testing.T) {
		emptySlice := make([]*list.List, 0)
		_, _, _, got := PlayerKilled(emptySlice, "DeadPlayer")
		want := errors.New("Input slice is empty")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("Passing a slice containing an empty list", func(t *testing.T) {
		inputSlice := []*list.List{list.New()}
		_, _, _, got := PlayerKilled(inputSlice, "DeadPlayer")
		want := errors.New("Input linked list empty")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("Passing a slice which does not contain the dead player", func(t *testing.T) {
		list1 := list.New()
		list1.PushBack("Gibberish1")
		list1.PushBack("Gibberish2")
		list2 := list.New()
		list2.PushBack("Gibberish3")
		list2.PushBack("Gibberish4")
		inputSlice := []*list.List{list1, list2}
		_, _, _, got := PlayerKilled(inputSlice, "DeadPlayer")
		want := errors.New("Desired string was not an element in any given list")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("Checking trivial single list combinations", func(t *testing.T) {
		testStrings := []string{"Walter -> Bob -> Tan10o", "Bob -> Tan10o", "Walter -> Bob", "Bob"}
		expectations := [][]string{{"Walter", "Tan10o"}, {"", "Tan10o"}, {"Walter", ""}, {"", ""}}
		for i := range testStrings {
			wantList, err := StringToLL(testStrings[i])
			if err != nil {
				t.Error(err)
			}
			wantSlice := []*list.List{wantList}
			gotList, gotHunter, gotTarget, err := PlayerKilled(wantSlice, "Bob")
			if err != nil {
				t.Error(err)
			}
			if gotHunter != expectations[i][0] {
				t.Errorf("Hunter: got %v, want %v", gotHunter, expectations[i][0])
			}
			if gotTarget != expectations[i][1] {
				t.Errorf("Hunter: got %v, want %v", gotTarget, expectations[i][1])
			}
			if wantList != gotList {
				t.Errorf("List: got %v, got %v", gotList, wantList)
			}
		}
	})
}

/*
############################################################

                    Linked List Operations

############################################################
*/

func TestFindElementInChain(t *testing.T) {
	t.Run("Input empty list", func(t *testing.T) {
		_, got := FindElementInChain(list.New(), "SearchTerm")
		want := errors.New("Input linked list empty")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Input list without desired object", func(t *testing.T) {
		testList := list.New()
		testList.PushBack("A string that won't return a match")
		testList.PushBack("A second string that won't return a match")
		_, got := FindElementInChain(testList, "SearchTerm")
		want := errors.New("Desired string was not an element in this list")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Desired object is in list", func(t *testing.T) {
		testList := list.New()
		testList.PushBack("Gibberish Element")
		testList.PushBack("Second Gibberish Element")
		want := testList.PushBack("Target")
		got, err := FindElementInChain(testList, "Target")
		if err != nil {
			t.Error(err)
		}
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	})
}

func TestFindElementInChains(t *testing.T) {
	t.Run("Passing a nil slice", func(t *testing.T) {
		var nilSlice []*list.List
		_, _, got := FindElementInChains(nilSlice, "TargetString")
		want := errors.New("Slice recieved is nil")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("Passing an empty slice", func(t *testing.T) {
		emptySlice := make([]*list.List, 0)
		_, _, got := FindElementInChains(emptySlice, "TargetString")
		want := errors.New("Slice recieved is empty")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("Passing a slice containing an empty list", func(t *testing.T) {
		inputSlice := []*list.List{list.New()}
		_, _, got := FindElementInChains(inputSlice, "TargetString")
		want := errors.New("Input linked list empty")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("Passing a slice which does not contain the target string", func(t *testing.T) {
		inputList := list.New()
		inputList.PushBack("Gibberish1")
		inputList.PushBack("Gibberish2")
		inputSlice := []*list.List{inputList}
		_, _, got := FindElementInChains(inputSlice, "TargetString")
		want := errors.New("Desired string was not an element in any given list")
		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("Passing a slice that does contain the target string", func(t *testing.T) {
		worthlessList := list.New()
		worthlessList.PushBack("Gibberish1")
		worthlessList.PushBack("Gibberish2")
		wantList := list.New()
		wantList.PushBack("Gibberish3")
		wantElement := wantList.PushBack("TargetString")
		inputSlice := []*list.List{worthlessList, wantList}
		gotList, gotElement, err := FindElementInChains(inputSlice, "TargetString")
		if err != nil {
			t.Error(err)
		}
		if gotList != wantList {
			t.Errorf("got %v, want %v", gotList, wantList)
		}
		if gotElement != wantElement {
			t.Errorf("got %v, want %v", gotElement, wantElement)
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

package assassinlogic

import (
	"container/list"
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

/*
############################################################

              LL / String Conversion

############################################################
*/

// This function returns a linked list if fed a string
// in the format "A -> B -> C"
func StringToLL(s string) (*list.List, error) {
	// Return empty list and error if input string is empty
	if s == "" {
		return list.New(), errors.New("Input string is empty")
	}
	var sValues []string = strings.Split(s, " -> ")
	output := list.New()
	for _, v := range sValues {
		output.PushBack(v)
	}
	return output, nil
}

// This function returns a string in the format
// "A -> B -> C" if fed a linked list
func LLToString(l *list.List) (string, error) {
	firstElement := l.Front()
	if firstElement == nil {
		return "", errors.New("Input list is empty")
	}
	output := ""
	for e := l.Front(); e != nil; e = e.Next() {
		output += fmt.Sprint(e.Value)
		if e.Next() != nil {
			output += " -> "
		}
	}
	return output, nil
}

/*
############################################################

                    Linked List Operations

############################################################
*/

func BuildLL(s []string) (*list.List, error) {
	if len(s) == 0 {
		return list.New(), errors.New("Input slice is empty")
	}
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	output := list.New()
	for _, v := range s {
		output.PushBack(v)
	}
	return output, nil
}

package assassinlogic

import (
	"container/list"
	"errors"
	"strings"
)

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
		output.PushFront(v)
	}
	return output, nil
}

// This function returns a string in the format
// "A -> B -> C" if fed a linked list
func LLToString(l *list.List) (string, error) {
	e := l.Front()
	if e == nil {
		return "", errors.New("Input list is empty")
	}
	return "", errors.New("Function unimplemented")
}

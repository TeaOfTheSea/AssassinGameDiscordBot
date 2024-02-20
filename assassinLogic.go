package assassinlogic

import (
	"container/list"
	"errors"
	"fmt"
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

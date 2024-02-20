package assassinlogic

import (
	"container/list"
	"errors"
	"strings"
)

func StringToLL(s string) (*list.List, error) {
	// Return empty list and error if input string is empty
	if s == "" {
		return list.New(), errors.New("No input given.")
	}
	var sValues []string = strings.Split(s, " -> ")
	output := list.New()
	for _, v := range sValues {
		output.PushFront(v)
	}
	return output, nil
}

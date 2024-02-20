package assassinlogic

import (
	"container/list"
	"errors"
)

func StringToLL(s string) (*list.List, error) {
	return list.New(), errors.New("No input given.")
}

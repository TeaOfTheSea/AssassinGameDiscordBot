package assassinlogic

import (
	"testing"
)

func TestStringToLL(t *testing.T) {
	//want := list.New()
	//for e := want.Front(); e != nil; e = e.Next() {
	//	t.Error(e.Value)
	//}
	_, err := StringToLL("")
	if err != nil {
		t.Error(err)
		return
	}
}

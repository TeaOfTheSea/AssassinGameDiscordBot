package assassinlogic

import (
	"testing"
  "errors"
  "fmt"
)

func TestStringToLL(t *testing.T) {
  t.Run("Passing an empty string and checking for error", func(t* testing.T){
	  _, got := StringToLL("")
    want := errors.New("No input given.")
    if fmt.Sprint(got) != fmt.Sprint(want) {
      t.Errorf("got %q want %q", got, want)
    }
  })
}

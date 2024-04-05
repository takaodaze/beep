package csv_test

import (
	"testing"

	"github.com/takaodaze/beep/csv"
)

func TestToRow(t *testing.T) {
	got := csv.ToRow([]string{"a", "b", "c"})
	want := "a,b,c"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

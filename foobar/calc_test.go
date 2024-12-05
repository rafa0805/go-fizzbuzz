package foobar

import "testing"

func TestMax(t *testing.T) {
  want := 5
  got := Tasu(2, 3)
  if want != got {
    t.Errorf("Tasu(2, 3) == %d, want %d", got, want)
  }
}

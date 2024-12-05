package fizzbuzzhoge

import "testing"

func TestMax(t *testing.T) {
  want := "fizz"
  got := FizzBuzz(9)
  if want != got {
    t.Errorf("FizzBuzz(9) == %s, want %s", got, want)
  }
}

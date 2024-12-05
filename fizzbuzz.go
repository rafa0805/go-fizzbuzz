package fizzbuzz

func FizzBuzz(n int) string {
  if n % 3 == 0 {
    return "fizz"
  }
  if n % 5 == 0 {
    return "buzz"
  }
  if n % 15 == 0 {
    return "fizzbuzz"
  }
  return ""
}


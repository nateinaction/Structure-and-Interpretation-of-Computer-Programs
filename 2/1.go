package main

import (
  "fmt"
)

func main() {
  oneHalf := makeRat(1,2)
  oneThird := makeRat(1,3)
  printRat(addRat(oneThird, oneThird))
  fmt.Println(equalRat(oneThird, oneHalf))
}

type Rat struct {
    numer, denom int
}

func makeRat(n int, d int) Rat {
  // check if fraction is negative and remove negatives
  neg, n, d := isNeg(n, d)

  // find greatest common divisor
  comDiv := gcd(n, d)
  n = n / comDiv
  d = d / comDiv

  // if negative, add negative to numerator
  if (neg) {
    n *= -1
  }

  return Rat{n, d}
}

func numer(x Rat) int {
  return x.numer
}

func denom(x Rat) int {
  return x.denom
}

func addRat(x Rat, y Rat) Rat {
  thisNumer := (x.numer * y.denom) + (x.denom * y.numer)
  thisDenom := x.denom * y.denom
  return makeRat(thisNumer, thisDenom)
}

func subRat(x Rat, y Rat) Rat {
  thisNumer := (x.numer * y.denom) - (x.denom * y.numer)
  thisDenom := x.denom * y.denom
  return makeRat(thisNumer, thisDenom)
}

func mulRat(x Rat, y Rat) Rat {
  thisNumer := x.numer * y.numer
  thisDenom := x.denom * y.denom
  return makeRat(thisNumer, thisDenom)
}

func divRat(x Rat, y Rat) Rat {
  thisNumer := x.numer * y.denom
  thisDenom := x.denom * y.numer
  return makeRat(thisNumer, thisDenom)
}

func equalRat(x Rat, y Rat) bool {
  if (x.numer == y.numer && x.denom == y.denom) {
    return true
  }
  return false
}

func printRat(x Rat) {
   fmt.Printf("%d/%d\n", x.numer, x.denom)
}

// Euclid's Algorithm
func gcd(a int, b int) int {
  if (b == 0) {
    return a
  }
  return gcd(b, a % b)
}

func isNeg(n int, d int) (bool, int, int) {
  neg := false
  if (n < 0) {
    n *= -1
    neg = !neg
  }
  if (d < 0) {
    d *= -1
    neg = !neg
  }
  return neg, n, d
}

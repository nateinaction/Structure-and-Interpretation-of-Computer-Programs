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

func makeRat(n, d int) Rat {
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

func addRat(x, y Rat) Rat {
  thisNumer := (x.numer * y.denom) + (x.denom * y.numer)
  thisDenom := x.denom * y.denom
  return makeRat(thisNumer, thisDenom)
}

func subRat(x, y Rat) Rat {
  thisNumer := (x.numer * y.denom) - (x.denom * y.numer)
  thisDenom := x.denom * y.denom
  return makeRat(thisNumer, thisDenom)
}

func mulRat(x, y Rat) Rat {
  thisNumer := x.numer * y.numer
  thisDenom := x.denom * y.denom
  return makeRat(thisNumer, thisDenom)
}

func divRat(x, y Rat) Rat {
  thisNumer := x.numer * y.denom
  thisDenom := x.denom * y.numer
  return makeRat(thisNumer, thisDenom)
}

func equalRat(x, y Rat) bool {
  if (x.numer == y.numer && x.denom == y.denom) {
    return true
  }
  return false
}

func printRat(x Rat) {
   fmt.Printf("%d/%d\n", x.numer, x.denom)
}

func numer(x Rat) int {
  return x.numer
}

func denom(x Rat) int {
  return x.denom
}

// Euclid's Algorithm
func gcd(a, b int) int {
  if (b == 0) {
    return a
  }
  return gcd(b, a % b)
}

func isNeg(n, d int) (bool, int, int) {
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

type Rat struct {
    numer, denom int
}

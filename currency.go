package main

import (
	"fmt"
)

func main() {
  var currencies = map[string]float32{
    "EUR": 0.939,
    "JPY": 130.180,
    "CAD": 1.266,
    "KWD": 0.306,
    "OMR": 0.385,
  }

  var dollarAmount float32
  var currency string

  fmt.Scan(&dollarAmount)
}

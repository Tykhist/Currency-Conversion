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

  fmt.Println("Currencies & Rates:", currencies)
  fmt.Println()

  var dollarAmount float32
  var currency string

  fmt.Print("Enter dollar amount: ")
  fmt.Scan(&dollarAmount)
  fmt.Println()

  if dollarAmount == 0 {
    fmt.Println("Error")
  } else {
    fmt.Print("Enter currency type: ")
    fmt.Scan(&currency)
    fmt.Println()

    rate, isValid := currencies[currency]

    if isValid == true {
      fmt.Println("Result: ", rate*dollarAmount, currency)
    } else {
      fmt.Println("Currency not found...")
    }

  }
}

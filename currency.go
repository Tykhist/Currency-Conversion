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
  fmt.Print("Enter a dollar amount above zero: ")
  fmt.Scan(&dollarAmount)
  for { 
    if dollarAmount == 0 {
      fmt.Print("Error, enter a number: ")
      fmt.Scan(&dollarAmount)
    } else {
      fmt.Print("Which currency would you like to convert to? ")
      fmt.Scan(&currency)
      rate, isValid := currencies[currency]
      if isValid == true {
        money := rate * dollarAmount
        return fmt.Println("You now have", money)
      } else {
        fmt.Println("Error, currency not found. Enter dollar amount again: ")
        fmt.Scan(&dollarAmount)
      }
      break
    }
  }
}

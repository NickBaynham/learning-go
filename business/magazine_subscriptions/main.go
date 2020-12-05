package main

import (
  "fmt"
  "github.com/nickbaynham/magazine"
)

func main() {
  // Create a Subscriber value
  var subscriber magazine.Subscriber
  subscriber.Rate = 4.99
  fmt.Println(subscriber.Rate)

  // Create an Employee value
  var employee magazine.Employee
  employee.Name = "Scott Steele"
  employee.Salary = 100000
  fmt.Println(employee.Name)
  fmt.Println(employee.Salary)

  // Create an Address value
  var address magazine.Address
  address.Street = "123 Oak St"
  address.City = "Omaha"
  address.State = "NE"
  address.PostalCode = "68111"
  fmt.Println(address)

  // Add the address to the Subscriber
  subscriber.HomeAddress = address
  fmt.Println(subscriber.HomeAddress)

  // Anonymous field of type Address
  employee.Street = "42 Main St"
  employee.City = "Birmingham"
  employee.PostalCode = "31008"
  employee.State = "GA"
}

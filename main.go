/*
Author: Resul Emre AYGAN
*/

package main

import (
	"fmt"
	"strconv"
)

type Calculator struct {
	result float64
}

func (c *Calculator) Add(number float64) {
	c.result += number
}

func (c *Calculator) Mul(number float64) {
	c.result *= number
}

func (c *Calculator) Division(number float64) error {
	if number != 0 {
		c.result /= number
		return nil
	}
	return fmt.Errorf("cannot divide by zero")
}

func (c *Calculator) Subtraction(number float64) {
	c.result -= number
}

func getInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return input
}

func main() {
	fmt.Println("Welcome to Calculator!")
	fmt.Println("Enter + for Addition: ")
	fmt.Println("Enter * for Multiplication: ")
	fmt.Println("Enter / for Division: ")
	fmt.Println("Enter - for Subtraction: ")
	fmt.Println("Enter reset for Reset: ")
	fmt.Println("Enter exit for Exit.")

	var calc Calculator

	for {
		fmt.Println("Enter the operator (+, -, *, /) or 'reset' to reset or 'exit' to quit: ")
		operator := getInput()

		if operator == "exit" {
			fmt.Println("Exiting calculator!")
			break
		}

		if operator == "reset" {
			calc.result = 0
			fmt.Println("Resetting calculator! Result:", calc.result)
			continue
		}

		fmt.Print("Enter the number: ")
		numberStr := getInput()

		number, err := strconv.ParseFloat(numberStr, 64)

		if err != nil {
			fmt.Println("Invalid input for the second number!")
			continue
		}

		switch operator {
		case "+":
			calc.Add(number)
		case "*":
			calc.Mul(number)
		case "/":
			err := calc.Division(number)
			if err != nil {
				return
			}
		case "-":
			calc.Subtraction(number)
		default:
			fmt.Println("Invalid operator!")
			continue
		}

		fmt.Printf("Result: %f\n", calc.result)
	}
}

/*
Author: Resul Emre AYGAN
*/

package main

import "fmt"

type Calculator struct {
	numberOne int
	numberTwo int
}

func (c *Calculator) Add() {
	fmt.Println("Addition of two numbers: ", c.numberOne+c.numberTwo)
}

func (c *Calculator) Mul() {
	fmt.Println("Multiplication of two numbers: ", c.numberOne*c.numberTwo)
}

func (c *Calculator) Division() {
	fmt.Println("Division of two numbers: ", c.numberOne/c.numberTwo)
}

func (c *Calculator) Subtraction() {
	fmt.Println("Subtraction of two numbers: ", c.numberOne-c.numberTwo)
}

func main() {
	var numberOne, numberTwo int

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanf("%d", &numberOne)
	if err != nil {
		return
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanf("%d", &numberTwo)
	if err != nil {
		return
	}

	calc := Calculator{numberOne: numberOne, numberTwo: numberTwo}

	value := 1

	for value >= 1 {
		fmt.Println("Enter 1 for Addition: ")
		fmt.Println("Enter 2 for Multiplication: ")
		fmt.Println("Enter 3 for Division: ")
		fmt.Println("Enter 4 for Subtraction: ")
		fmt.Println("Enter 5 for Exit.")

		_, err2 := fmt.Scanf("%d", &value)
		if err2 != nil {
			return
		}

		switch value {
		case 1:
			calc.Add()
		case 2:
			calc.Mul()
		case 3:
			calc.Division()
		case 4:
			calc.Subtraction()
		case 5:
			value = 0
			break
		default:
			fmt.Println("Enter valid number!")
		}
	}
}

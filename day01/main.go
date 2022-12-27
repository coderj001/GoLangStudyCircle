package main

import (
	"fmt"
	"math"
)

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func sub(num1, num2 float64) float64 {
	return num1 - num2
}

func multi(num1, num2 float64) float64 {
	return num1 * num2
}

func div(num1, num2 float64) float64 {
	return num1 / num2
}

func sin(num float64) float64 {
	return math.Sin(num)
}

func cos(num float64) float64 {
	return math.Cos(num)
}

func tan(num float64) float64 {
	return math.Tan(num)
}

func sqr(num float64) float64 {
	return math.Sqrt(num)
}

func main() {
	fmt.Println("Welcome to the simple calculator!")
	fmt.Println("Enter the operation you want to perform (add, sub, multi, div, sin, cos, tan, sqr):")

	var operation string
	fmt.Scanln(&operation)

	fmt.Println("Enter the two numbers you want to perform the operation on:")

	var num1, num2 float64
	fmt.Scanln(&num1, &num2)

	var result float64
	switch operation {
	case "add":
		result = add(num1, num2)
	case "sub":
		result = sub(num1, num2)
	case "multi":
		result = multi(num1, num2)
	case "div":
		result = div(num1, num2)
	case "sin":
		result = sin(num1)
	case "cos":
		result = cos(num1)
	case "tan":
		result = tan(num1)
	case "sqr":
		result = sqr(num1)
	default:
		fmt.Println("Invalid operation")
	}

	fmt.Println("Result:", result)
}

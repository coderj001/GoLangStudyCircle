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

func Div(num1, num2 float64) float64 {
	return num1 / num2
}

func Sin(num float64) float64 {
	return math.Sin(num)
}

func Cos(num float64) float64 {
	return math.Cos(num)
}

func Tan(num float64) float64 {
	return math.Tan(num)
}

func Sqr(num float64) float64 {
	return math.Sqrt(num)
}

func main() {
	fmt.Println("Hello Go!!")
}

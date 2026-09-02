package cmd

import (
	"fmt"
	"strconv"
)

func Add(first, second string) string {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
	}

	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is invalid")
	}

	return fmt.Sprintf("%.2f", num1+num2)
}

func Subtract(first, second string) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Printf("error: First value %s is invalid", first)
		return
	}

	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Printf("error: Second value %s is invalid", second)
		return
	}

	return fmt.Sprintf("%.2f", num1-num2)
}

func Multiply(first, second string, shouldRoundUp bool) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}

	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is invalid")
		return
	}

	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1*num2)
	}

	return fmt.Sprintf("%f", num1*num2)
}

func Divide(first, second string, shouldRoundUp bool) (string, error) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return "", fmt.Errorf("first value is invalid")
	}

	if num1 <= 0 {
		return "", fmt.Errorf("first value is required to be greater than zero")
	}

	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return "", fmt.Errorf("second value is invalid")
	}

	if num2 <= 0 {
		return "", fmt.Errorf("second value is required to be greater than zero")
	}

	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1/num2), nil
	}

	return fmt.Sprintf("%f", num1/num2), nil
}

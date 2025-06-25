package main

import (
	"fmt"

	"github.com/ignratnan/go-calculator/internal/user"
	// You'll import your internal packages using the full module path:
	// "github.com/yourusername/my-web-app/internal/models"
)

func main() {
	//Variable description
	var first_num float64
	var second_num float64
	var operator string
	var precision uint = 2

	fmt.Println("Welcome to Go Calculator")

	fmt.Print("Input the first number: ")
	_, err := fmt.Scanln(&first_num)
	if err != nil {
		fmt.Println("Error reding first number: ", err)
		return
	}

	fmt.Print("Input the operator ( + , - , * , / ): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Error reding first number: ", err)
		return
	}

	fmt.Print("Input the second number: ")
	_, err = fmt.Scanln(&second_num)
	if err != nil {
		fmt.Println("Error reding second number: ", err)
		return
	}

	// If Else Version
	/*
		if operator == "+" {
			result := user.Add(first_num, second_num)
			fmt.Println("Result:", result)
		} else if operator == "-" {
			result := user.Sub(first_num, second_num)
			fmt.Println("Result:", result)
		} else if operator == "*" {
			result := user.Mul(first_num, second_num)
			fmt.Println("Result:", result)
		} else if operator == "/" {
			result := user.Div(first_num, second_num)
			fmt.Println("Result:", result)
		} else {
			fmt.Println("Invalid operator")
		}
	*/

	// Switch Case Version
	switch operator {
	case "+":
		result := user.Add(first_num, second_num)
		f_result := user.RoundFloat(result, precision)
		fmt.Println("Result:", f_result)
	case "-":
		result := user.Sub(first_num, second_num)
		f_result := user.RoundFloat(result, precision)
		fmt.Println("Result:", f_result)
	case "*":
		result := user.Mul(first_num, second_num)
		f_result := user.RoundFloat(result, precision)
		fmt.Println("Result:", f_result)
	case "/":
		result := user.Div(first_num, second_num)
		f_result := user.RoundFloat(result, precision)
		fmt.Println("Result:", f_result)
	default:
		fmt.Println("Invalid operator")
	}

}

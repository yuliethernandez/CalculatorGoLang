package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	for {
		var val string
		fmt.Println("*** Welcome o the Calculator with goLang **** \nChoose a mathematical operation:\n1. Addition \n2. Subtraction \n3. Multiplication \n4. Division \n5. Square root \n6. Factorial \n7. Modulus \nType the number of the option you wish use or exit or quit for finish the program")
		fmt.Scan(&val)
		if !isFinish(val) {
			switch val {
			case "1": //Addition
				{
					/*var num1, num2 float64
					fmt.Println("Enter two numbers:")
					fmt.Scan(&num1, &num2)
					fmt.Println(num1, " + ", num2, " = ", add(int(num1), int(num2)))*/
					var a, b int
					fmt.Print("Enter the first number: ")
					fmt.Scan(&a)
					if !isFinish(strconv.Itoa(a)) {
						fmt.Print("Enter the second number: ")
						fmt.Scan(&b)
						if !isFinish(strconv.Itoa(b)) {
							fmt.Println(a, " + ", b, " = ", add(a, b))
						} else {
							break
						}
					} else {
						break
					}
				}
			case "2": //Substraccion
				{
					var a, b int
					fmt.Print("Enter the first number: ")
					fmt.Scan(&a)
					if !isFinish(strconv.Itoa(a)) {
						fmt.Print("Enter the second number: ")
						fmt.Scan(&b)
						if !isFinish(strconv.Itoa(b)) {
							subs := func() int { //closure or anonymus function
								return a - b
							}
							fmt.Println(a, " - ", b, " = ", subs())
						} else {
							goto FINAL_STATEMENT
						}
					} else {
						goto FINAL_STATEMENT
					}
				}
			case "3": //Multiplication
				{
					var a, b int
					fmt.Print("Enter the first operand: ")
					fmt.Scan(&a)
					if !isFinish(strconv.Itoa(a)) {
						fmt.Print("Enter the second operand: ")
						fmt.Scan(&b)
						if !isFinish(strconv.Itoa(b)) {
							fmt.Println(a, " * ", b, " = ", multiply(a, b))
						} else {
							break
						}
					} else {
						break
					}
				}
			case "4": //Division
				{
					var a, b int
					fmt.Print("Enter the first operand: ")
					fmt.Scan(&a)
					if !isFinish(strconv.Itoa(a)) {
						fmt.Print("Enter the second operand: ")
						fmt.Scan(&b)
						if !isFinish(strconv.Itoa(b)) {
							fmt.Println(a, " / ", b, " = ", division(a, b))
						} else {
							break
						}
					} else {
						break
					}
				}
			case "5": //Square Root
				{
					var a int
					fmt.Print("Enter the number: ")
					fmt.Scan(&a)
					if !isFinish(strconv.Itoa(a)) {
						fmt.Println(a, " = ", squareRoot(a))
					} else {
						break
					}
				}
			case "6": //Factorial
				{
					var a int
					fmt.Print("Enter the number: ")
					fmt.Scan(&a)
					if !isFinish(strconv.Itoa(a)) {
						fmt.Println("!", a, " = ", factorial(a))
					} else {
						break
					}
				}
			case "7": //Modulus
				{
					var a, b int
					fmt.Print("Enter the first operand: ")
					fmt.Scan(&a)
					if !isFinish(strconv.Itoa(a)) {
						fmt.Print("Enter the second operand: ")
						fmt.Scan(&b)
						if !isFinish(strconv.Itoa(b)) {
							fmt.Println(a, " % ", b, " = ", modulus(a, b))
						} else {
							break
						}
					} else {
						break
					}
				}
			}
		} else {
			goto FINAL_STATEMENT
		}
	}
FINAL_STATEMENT:
	fmt.Print("Program finished!")
}

func isFinish(val string) bool {
	if strings.ToLower(val) == "quit" || strings.ToLower(val) == "exit" {
		return true
	}
	return false
}

func add(a, b int) (c int) {
	c = a + b
	return
}

func multiply(a, b int) int {
	return a * b
}

func division(a, b int) int {
	// Use this deferred function to handle errors.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR!!!!")
			fmt.Println(err)
		}
	}()
	// Cause an error ... Go will run the defer func above.
	return int(a / b)
}

func modulus(a, b int) int {
	return a % b
}

// The math.Sqrt function signature: func Sqrt(x float64) float64
func squareRoot(a int) int {
	c := int(math.Sqrt(float64(a)))
	return c
}

func factorial(a int) int {
	if a == 0 {
		return 1
	} else {
		return a * factorial(a-1)
	}
}

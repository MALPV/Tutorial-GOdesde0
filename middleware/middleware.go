package middleware

import "fmt"

func addNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func subtractNumbers(num1 int, num2 int) int {
	return num1 - num2
}

func MyMiddleware() {

	fmt.Println("Starting Middleware")

	result := operationsMidd(addNumbers)(2,3)
	fmt.Println(result)

	result = operationsMidd(subtractNumbers)(10,3)
	fmt.Println(result)

}

func operationsMidd(fun func(int, int) int) func(int, int) int {
	return func(n1, n2 int) int {
		fmt.Println("Starting Operations") 
		return fun(n1, n2)
	}
}
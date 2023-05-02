package funciones

import "fmt"

var addNumbers = func(num1 int, num2 int) int {
	return num1 + num2
}

var subtractNumbers = func(num1 int, num2 int) int {
	return num1 - num2
}

var multiplyNumbers = func(num1 int, num2 int) int {
	return num1 * num2
}

var divideNumbers = func(num1 int, num2 int) int {
	return num1 / num2
}

func Calculate() {

	fmt.Println("Suma ", addNumbers(7, 7))
	fmt.Println("Resta ", subtractNumbers(7, 7))
	fmt.Println("Multiplicacion ", multiplyNumbers(7, 7))
	fmt.Println("Division ", divideNumbers(7, 7))

}

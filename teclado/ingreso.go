package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var msg string
var err error

func MultiplyNumbers() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el primer numero :")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El valor ingresado es incorrecto :" + err.Error())
		}
	}

	fmt.Println("Ingrese el segundo numero :")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El valor ingresado es incorrecto :" + err.Error())
		}
	}

	fmt.Println("Ingrese el mensaje :")
	if scanner.Scan() {
		msg = scanner.Text()
	}

	fmt.Println(msg, num1*num2)
}

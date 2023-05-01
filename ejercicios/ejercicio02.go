package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error

func CalculateMultiplicationTable() string {

	scanner := bufio.NewScanner(os.Stdin)

	var text string

	for {
		fmt.Println("Ingrese un numero")

		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("El valor ingresado es incorrecto")
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		text += fmt.Sprintf("%d x %d = %d \n", i, num, i*num)
	}

	return text
}

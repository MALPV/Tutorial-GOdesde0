package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error

func CalculateMultiplicationTable() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
	
			if err != nil {
				fmt.Println("El valor ingresado es incorrecto")
			}else{
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i, "x", num, "=", i*num)
	}
}

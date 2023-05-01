package main

import (
	"fmt"
	"runtime"

	"github.com/MALPV/Tutorial-Godesde0/ejercicios"
	"github.com/MALPV/Tutorial-Godesde0/iteraciones"
	"github.com/MALPV/Tutorial-Godesde0/teclado"
	"github.com/MALPV/Tutorial-Godesde0/variables"
)

var showOldFunctions bool

func main() {
	fmt.Println("Hola Mundo...")

	showOldFunctions = false

	if showOldFunctions {

		variables.ShowInterger()
		variables.ShowOtherValues()

		fmt.Println(variables.ConvertToText(15032023))

		if os := runtime.GOOS; os == "windows" {
			fmt.Println("Es Windows ")
		} else {
			fmt.Println("No es Windows ")
		}

		switch os := runtime.GOOS; os {
		case "windows":
			fmt.Println("Es Windows")
		case "linux":
			fmt.Println("Es Linux")
		default:
			fmt.Printf("%s \n", os)
		}

		fmt.Println(ejercicios.ConvertToInt("1998"))

		teclado.MultiplyNumbers()

		iteraciones.Iterar()

	} else {

		ejercicios.CalculateMultiplicationTable()

	}
}
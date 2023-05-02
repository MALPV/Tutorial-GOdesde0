package main

import (
	"fmt"
	"runtime"

	"github.com/MALPV/Tutorial-Godesde0/arreglos_slices"
	"github.com/MALPV/Tutorial-Godesde0/defer_panic"
	"github.com/MALPV/Tutorial-Godesde0/ejercicios"
	"github.com/MALPV/Tutorial-Godesde0/files"
	"github.com/MALPV/Tutorial-Godesde0/funciones"
	"github.com/MALPV/Tutorial-Godesde0/goroutines"
	"github.com/MALPV/Tutorial-Godesde0/iteraciones"
	"github.com/MALPV/Tutorial-Godesde0/mapas"
	"github.com/MALPV/Tutorial-Godesde0/modelos"
	"github.com/MALPV/Tutorial-Godesde0/teclado"
	"github.com/MALPV/Tutorial-Godesde0/users"
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

		fmt.Println(ejercicios.CalculateMultiplicationTable())

		files.SaveTable()
		files.AddNewTable()
		files.ReadFile()

		funciones.Calculate()
		funciones.UseClosure(7)
		funciones.Exponential(2)

		arreglos_slices.UseArrays()
		arreglos_slices.ShowSlices()
		arreglos_slices.Capacity()

		mapas.CreateMaps()

		users.CreateAdmin()

		Marcos := new(modelos.Hombre)
		ejercicios.HumanoEscribiendo(Marcos)
		ejercicios.EstaVivo(Marcos)

		Carolina := new(modelos.Mujer)
		ejercicios.HumanoEscribiendo(Carolina)
		ejercicios.EstaVivo(Carolina)

		defer_panic.PanicExample()

	} else {

		channel := make(chan bool)
		go goroutines.ShowStringLower("Martina", channel)

		defer func ()  {
			<- channel
		}()

		fmt.Println("Principal")

	}
}

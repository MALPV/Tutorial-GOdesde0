package arreglos_slices

import "fmt"

var vector [10]int
var vectorWithValues [4]int = [4]int{1, 2, 3}

var matriz [5][5]int

func UseArrays() {

	vector[5] = 5

	vector[8] = 66

	names := [5]string{"Martina","Carolina","Marcos"}

	fmt.Println(vector)
	fmt.Println(vectorWithValues)
	fmt.Println(names)
	fmt.Println(matriz)
}
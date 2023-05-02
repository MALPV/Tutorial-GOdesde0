package funciones

import "fmt"

func generateMultiplicationTable(value int) func() int {

	num := value
	sequence := 0

	return func() int {
		sequence++
		return num * sequence
	}
}

func UseClosure(tableOf int) {

	miTable := generateMultiplicationTable(tableOf)

	for i := 1; i <= 10; i++ {
		fmt.Println(miTable())
	}
}

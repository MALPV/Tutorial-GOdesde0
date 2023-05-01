package iteraciones

import "fmt"

func Iterar() {

	// Iteracion normal

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Iterar con saltos de 10

	for i := 0; i < 100; i += 10 {
		fmt.Println(i)
	}

	// Iterar con break y continue

	for i := 10; i > 1; i-- {

		if i == 5 {
			continue
		}

		if i == 2 {
			break
		}

		fmt.Println(i)
	}
}

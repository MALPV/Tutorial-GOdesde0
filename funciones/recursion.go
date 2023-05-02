package funciones

import "fmt"

func Exponential(value int) {
	if value > 10000 {
		return
	}

	fmt.Println(value)
	Exponential(value * 2)
}
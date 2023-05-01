package ejercicios

import "strconv"

func ConvertToInt(value string) (int, string) {

	if num, err := strconv.Atoi(value); err != nil {
		return num, "El valor ingresado no es un numero"
	} else if num > 100 {
		return num, "El valor ingresado es mayor a 100"
	} else {
		return num, "El valor ingresado es menor a 100"
	}

}

package files

import (
	"fmt"
	"os"

	"github.com/MALPV/Tutorial-Godesde0/ejercicios"
)

func SaveTable() {

	var text string = ejercicios.CalculateMultiplicationTable()

	// abre nun buffer para crear un archivo
	archive, err := os.Create("./files/txt/result-table.txt")

	if err != nil {
		fmt.Println("Error al crear archivo " + err.Error())
		return
	}

	fmt.Fprintln(archive, text)
	archive.Close()
}

package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MALPV/Tutorial-Godesde0/ejercicios"
)

var filePath string = "./files/txt/result-table.txt"

func SaveTable() {

	var text string = ejercicios.CalculateMultiplicationTable()

	// abre nun buffer para crear un archivo
	archive, err := os.Create(filePath)

	if err != nil {
		fmt.Println("Error al crear archivo " + err.Error())
		return
	}

	fmt.Fprintln(archive, text)
	archive.Close()
}

func AddNewTable() {

	var text string = ejercicios.CalculateMultiplicationTable()

	if !Append(filePath, text) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(file string, text string) bool {

	archive, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el OpenFile " + err.Error())
		return false
	}

	textLn := text + "\n"

	_, err = archive.WriteString(textLn)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	archive.Close()
	return true
}

func ReadFile() {

	archive, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Hubo un error en la lectura del archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archive)

	for scanner.Scan() {
		register := scanner.Text()
		fmt.Println("-> " + register)
	}

	archive.Close()
}

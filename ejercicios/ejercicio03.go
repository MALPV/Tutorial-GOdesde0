package ejercicios

import (
	"fmt"

	"github.com/MALPV/Tutorial-Godesde0/interfaces"
	"github.com/MALPV/Tutorial-Godesde0/modelos"
)

func PersonaEstaLeyendo(persona modelos.Persona) {

	persona.Leer()
	fmt.Printf("La persona esta leyendo? %t\n", persona.EstaLeyendo)

}

func HumanoEscribiendo(humando interfaces.Humano) {

	humando.Escribir()
	fmt.Printf("El humano es de sexo %s y esta escribiendo...\n", humando.Sexo())

}

func EstaVivo(serVivo interfaces.SerVivo) {

	fmt.Printf("El ser vivo es de sexo %s y esta esta vivo? %t\n", serVivo.Sexo(), serVivo.EstaVivo())

}

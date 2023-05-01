package variables

import (
	"fmt"
	"strconv"
	"time"
)

func ShowInterger() {

	var num1 int
	num2 := 2

	fmt.Println("num 1 ", num1)
	fmt.Println("num 2 ", num2)
}

var Name string
var State bool
var Quantity float32
var Date time.Time

func ShowOtherValues() {

	Name = "Martina"
	State = true
	Quantity = 15.03
	Date = time.Now()

	fmt.Println("Name ", Name, "State ", State, "Quantity ", Quantity, "Date ", Date)

}

func ConvertToText(number int) (bool, string) {

	text := strconv.Itoa(number)
	return true, text

}

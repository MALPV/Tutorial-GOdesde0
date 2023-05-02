package arreglos_slices

import "fmt"

var numbers []int = []int{1, 2, 3, 4}
var array [10]int = [10]int{77, 50, 44, 1000, 1, 7, 64, 23, 25}

func ShowSlices() {

	fmt.Println(numbers)

	slice1 := array[5:]  // slice creado desde un arreglo, desde la posicion 5 hasta la final
	slice2 := array[:5]  // slice creado desde un arreglo, desde la posicion inicial hasta la 5
	slice3 := array[4:8] // slice creado desde un arreglo, desde la posicion 4 hasta la 8

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

}

func Capacity() {

	elements := make([]int, 5, 20)
	fmt.Printf("elements -> Large %d, capacity %d", len(elements), cap(elements))

	nums := make([]int, 0)
	fmt.Printf("\nnums -> Large %d, capacity %d", len(nums), cap(nums))
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nnums -> Large %d, capacity %d", len(nums), cap(nums))

}

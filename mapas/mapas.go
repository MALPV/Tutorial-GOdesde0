package mapas

import "fmt"

func CreateMaps() {

	brands := make(map[string]string)

	brands["sp"] = "Supreme"
	brands["nk"] = "Nike"

	fmt.Println(brands)

	pointsResult := map[string]int{
		"U. de Chile": 44,
		"U. Catolica": 38,
		"Colo Colo": 39,
		"Huachipato": 32,
	}

	fmt.Println(pointsResult)

	for team, point := range pointsResult {
		fmt.Printf("Team: %s - Point: %d\n", team, point)
	}

	delete(pointsResult, "Huachipato")
	fmt.Println(pointsResult)

	point, exist := pointsResult["Huachipato"]
	fmt.Printf("Point: %d, ifExist %t\n", point, exist)
	
}
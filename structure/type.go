package main

import (
	"encoding/json"
	"fmt"
)

type vehicle interface {
	start() string
}

type Car struct {
	CarName string `json:"car"`
	CarYear int    `json:"name"`
}

func (c Car) drive() {
	fmt.Println(c.CarName, "Andou")
}

func (c Car) start() string {
	return "vrum vrum"
}

func exemplo(car vehicle) {

}

func main() {
	car1 := Car{
		CarName: "Fusion",
		CarYear: 2020,
	}
	j := []byte(`{"car":"BMW","year":"2020"}`)
	var car Car
	json.Unmarshal(j, &car)

	fmt.Println(car.CarName)
	car1.drive()

	result, _ := json.Marshal(car1)

	fmt.Println(string(result))

	exemplo(car1)
}

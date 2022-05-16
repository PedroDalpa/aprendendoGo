package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	print("Hello\n")
	name := "Pedro"
	println(name)

	fmt.Printf("%v \n", name)

	res, err := http.Get("https://www.google.com/")

	if err != nil {
		panic("erroooo")
	}

	fmt.Println(res.Body)
	value := 10
	appoint := &value
	fmt.Println(appoint) // endereço na memoria
	fmt.Println(*appoint)

	fmt.Println(pointer(appoint))

	// a, b, err := learnErr("Pedro", 11)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Println(a, b)

	x := [10]int{4, 5, 6, 7}

	x[4] = 1

	fmt.Println(x[5])

	slice := make([]int, 5)

	slice = append(slice, 1)

	fmt.Println(slice)

}

func pointer(pointer *int) int {
	return *pointer
}

func learnErr(a string, b int) (string, int, error) {

	if b > 10 {
		return "", 0, errors.New("Deu ruim ")
	}
	return a, b, nil
}

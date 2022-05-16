package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	~int | ~float64 | ~int64
}

type MyNumber int

func SumGeneric[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}

	return sum
}

// Comparar é melhor usar o tipo comparable ao invés do any quando é
// nescessario comparar as variaveis, assim o go nao permite na chamada
// de função mandar um objeto e uma string, ou um inteiro e uma string
func Comparar[T comparable](numberOne T, numberTwo T) T {
	if numberOne == numberTwo {
		return numberOne
	}
	return numberTwo
}

func Max[T constraints.Ordered](one T, two T) T {
	if one > two {
		return one
	}
	return two
}

type stringer interface {
	String() string
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func concat[T stringer](vals []T) string {
	result := ""

	for _, v := range vals {
		result += v.String()
	}

	return result
}

func main() {
	Comparar("a", "b") //✅
	Comparar(1.1, 1.3) //✅
	// Comparar("a", 1) //❌

	fmt.Println(Max(1, 2))
	fmt.Println(concat([]MyString{"a", "b"}))

}

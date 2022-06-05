package main

import "fmt"

var message string = "nolan"
var matrixx [][]int = make([][]int, 10)

func handling_panic() {
	for x, y := 0, 0; y < 10 && x < 10; {
		matrixx[y] = make([]int, 10)
		x++
		y++
	}
	fmt.Println()
	counter := 0
	for i := range matrixx {
		for j := range matrixx[i] {
			matrixx[i][j] = counter
			counter++
		}
		fmt.Println(matrixx[i])
	}
	assignVal(&matrixx)
	fmt.Println(matrixx)
	fmt.Println("program shutted down")
}

func assignVal(matrice *[][]int) {
	defer handlePanic()
	(*matrice)[10] = make([]int, 10)
}

func handlePanic() {
	if rec := recover(); rec != nil {
		fmt.Println(rec, "handled with no termination")
	}
}

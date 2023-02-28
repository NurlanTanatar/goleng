package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func printSlice(s []int) {
	if len(s) == 0 {
		return
	}
	if s[0] == 0 {
		fmt.Print("O")
	}
	printSlice(s[1:])
}

func B_matrix() {
	var height int
	fmt.Scan(&height)

	var width int
	fmt.Scan(&width)
	a := make([][]int, height)
	for i := range a {
		a[i] = make([]int, width)
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	for i := range a {
		printSlice(a[i])
		fmt.Println()
	}
}

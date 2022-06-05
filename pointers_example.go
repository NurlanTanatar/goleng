package main

import "fmt"

func WorkingWithPointer() {
	message := "nolan"
	fmt.Printf("%p ", &message)
	fmt.Println(message)

	message = withoutPointer(message)
	fmt.Printf("%p --without pointer-- ", &message)
	fmt.Println(message)

	message = "nolan"
	withPointer(&message)
	fmt.Printf("%p --with pointer-- ", &message)
	fmt.Println(message)

}

func withoutPointer(randString string) string {
	fmt.Printf("%p in func\n", &randString)
	randString += randString
	return randString
}

func withPointer(randString *string) {
	*randString += *randString
}

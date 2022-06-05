package main

import "fmt"

func Mapping() {
	users := map[int]string{
		1: "Nolan",
		2: "Nurlan",
		3: "wh1t34ox",
	}
	users[2] = "NurlanT"  // resassign
	delete(users, 1)      // delete by key
	users[4] = "wh1t3fox" // init new pair - key and val
	if val, exists := users[2]; exists {
		fmt.Println(val, exists) // expected -empty string-, false
	} else {
		fmt.Println(val, exists) // expected NurlanT, true
	}
	for key, val := range users {
		fmt.Println(key, val)
	}

}

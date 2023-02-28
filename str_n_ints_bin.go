package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Converter() {
	// You are given a String, which contains letters and/or digits. You have to transform the String by converting ALL the letters first (in the order the letters appear in the string),
	// and after that, converting the digits (in the order the digits appear in the string). The two parts are then joined together.
	// Each letter is converted into its ASCII code (in base 10), and then converted to binary.
	// Each digit is converted directly to binary.

	var s string
	fmt.Scan(&s)
	conv := make([]string, len(s))
	nums := make([]string, len(s))
	for i, char := range s {
		if unicode.IsLetter(char) {
			conv[i] = fmt.Sprint(int(char))
		} else {
			conv[i] = string(char)
		}
	}

	for i, num := range conv {
		conv_to_int, _ := strconv.Atoi(num)
		conv_to_int64 := int64(conv_to_int)
		nums[i] = strconv.FormatInt(conv_to_int64, 2)
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(strings.Join(nums, "")) // Write answer to stdout
}

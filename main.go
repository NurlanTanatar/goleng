package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	f := scanner.Text()
	formatted := make([]string, len(f))
	inds := 0
	for i, chr := range f {
		if chr == 'X' {
			formatted[i] = strings.ToUpper(string(s[inds]))
			inds++
		} else if chr == 'x' {
			formatted[i] = strings.ToLower(string(s[inds]))
			inds++
		} else {
			formatted[i] = string(f[i])
		}
	}
	fmt.Println(strings.Join(formatted, ""))
}

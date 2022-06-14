package main

import (
	"fmt"
	"math/big"
	"time"
)

func main_factorial() {
	result := make(chan *big.Int)
	calc := make(chan bool)
	var number *big.Int = big.NewInt(100000)
	go func() {
		result <- factorial(number)
		close(result)
	}()
	go waitin(calc)
	calc <- true
	fmt.Println(<-result)
	close(calc)
}

func factorial(n *big.Int) *big.Int {
	one := big.NewInt(1)
	if n.Cmp(big.NewInt(0)) == 0 {
		return one
	}
	return one.Mul(n, factorial(one.Sub(n, one)))
}

func waitin(quit chan bool) {
	for {
		select {
		case <-quit:
			fmt.Println()
			break
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}

}

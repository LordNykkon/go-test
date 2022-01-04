package main

import (
	"fmt"
	"math/rand"
	"time"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Quote of the day is:")
	quotes := map[int]func() string{
		0: quote.Glass,
		1: quote.Go,
		2: quote.Hello,
		3: quote.Opt,
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(4)]())
}

package main

import (
	"flag"
	"fmt"

	"search-engine-golang/task-1/pkg/fibonacci"
)

func main() {
	num := flag.Int("n", 2, "Fibonacchi number")
	flag.Parse()

	if *num < 1 || *num > 20 {
		fmt.Println("Number must be between 1 and 20")
		return
	}

	fmt.Println(fibonacci.Number(*num))
}

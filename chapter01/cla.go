package main

import (
	"fmt"
	"os"
	"strconv"
)

const FIRST_ARG = 1

func main() {
	args := os.Args
	argsLen := len(args)

	if argsLen == FIRST_ARG {
		fmt.Println("Need one or more arguments!")
		return
	}

	var min, max float64
	for i := FIRST_ARG; i < argsLen; i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}

		if i == FIRST_ARG {
			max = n
			min = n
			continue
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var total, nInts, nFloats int
	invalid := make([]string, 0)

	for _, v := range os.Args[1:] {
		_, err := strconv.Atoi(v)
		if err == nil {
			total++
			nInts++
			continue
		}

		_, err1 := strconv.ParseFloat(v, 64)
		if err1 == nil {
			total++
			nFloats++
			continue
		}

		invalid = append(invalid, v)
	}

	fmt.Printf("#read: %v #ints: %v #floats: %v", total, nInts, nFloats)
	fmt.Println()

	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))

		for _, v := range invalid {
			fmt.Println(v)
		}
	}
}

package main

import (
	"fmt"
	"io"
	"os"
)

type User struct {
	Name    string
	SurName string
	Tel     int
}

func Search(phoneBook []User, surName string, out io.Writer) {
	for _, u := range phoneBook {
		if u.SurName == surName {
			fmt.Fprint(out, u)
		}
	}
}

func List(phoneBook []User, out io.Writer) {
	for _, v := range phoneBook {
		fmt.Fprintln(out, v)
	}
}

var phoneBook = []User{}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := arguments[0]
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	phoneBook = append(phoneBook, User{"Mihalis", "Tsoukalos", 2109416471})
	phoneBook = append(phoneBook, User{"Mary", "Doe", 2109416471})
	phoneBook = append(phoneBook, User{"John", "Black", 2109416123})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search SurName")
			return
		}

		Search(phoneBook, arguments[2], os.Stdout)

	case "list":
		List(phoneBook, os.Stdout)
	default:
		fmt.Println("Not a valid option")
	}

}

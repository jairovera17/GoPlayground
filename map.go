package main

import "fmt"

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Hex code for ", key, " color is ", value)
	}
}

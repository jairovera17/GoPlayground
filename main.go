package main

import "fmt"

func main() {
	colors := make(map[int]string)
	colors[10] = "#ffffff"
	delete(colors, 10)
	// colors := map[string]string{
	// 	"red":   "#ff000",
	// 	"green": "#4bf745",
	// }

	fmt.Println(colors)
}

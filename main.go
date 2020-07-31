package main

func main() {
	cards := newDeckFromFile("my_file.txt")
	cards.shuffle()
	cards.print()
}

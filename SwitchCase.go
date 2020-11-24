package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("hello")
	var rez string

	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Unesite ocjenu:")
	Scanner.Scan()
	input := Scanner.Text()

	inputrez, _ := strconv.Atoi(input)

	switch inputrez {
	case 1:
		rez = "nedovoljan"
	case 2:
		rez = "dovoljan"
	case 3:
		rez = "dobar"
	case 4:
		rez = "odlican"
	case 5:
		rez = "vrlo dobar"
	default:
		rez = "pogresan unos"
	}

	fmt.Println(rez)

}

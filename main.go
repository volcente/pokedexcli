package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			return
		}

		formattedInput := cleanInput(scanner.Text())
		if len(formattedInput) == 0 {
			return
		}

		fmt.Printf("Your command was: %s\n", formattedInput[0])
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanText := cleanInput(scanner.Text())
		textList := strings.Split(cleanText[0], " ")
		command := textList[0]
		fmt.Println("Your command was:", command) // Println will add back the final '\n'
	}
}

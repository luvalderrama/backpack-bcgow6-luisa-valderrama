package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	runes := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your word: ")
	input, _ := runes.ReadString('\n')
	fmt.Println("total de palabras: ", len(input))
	for i := 0; i < len(input) ; i++ {
    	fmt.Printf("%c\n", input[i])
	}
}
package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	cvsFile, err := os.ReadFile("productos.csv")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.ReplaceAll(string(cvsFile), ",", "\t\t\t"))
}

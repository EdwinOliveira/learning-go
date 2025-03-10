package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	names := []string{}

	for range 3 {
		fmt.Println("Hello. Name?")
		response, _ := reader.ReadString('\n')
		trimmedResponse := strings.TrimSpace(response);

		names = append(names, trimmedResponse)
	}

	fullName := ""

	for _, name := range names {
		fmt.Println("Name:", name)
		fullName += name 
	}

	fmt.Println(fullName)
}
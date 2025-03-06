package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello. First Name?")
	firstName, _ := reader.ReadString('\n')

	fmt.Println("Hello. Second Name?")
	middleName, _ := reader.ReadString('\n')

	fmt.Println("Hello. Last Name?")
	lastName, _ := reader.ReadString('\n')

	fmt.Println("Hello", strings.TrimSpace(firstName), strings.TrimSpace(middleName), strings.TrimSpace(lastName))
}
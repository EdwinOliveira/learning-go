package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	fullName string
	dateBirth string
	age string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	students := [2]Student{}

	for i := range len(students) {
		fmt.Println("Hello. Name?")
		response, _ := reader.ReadString('\n')
		trimmedResponse := strings.TrimSpace(response);
		students[i].fullName = trimmedResponse 

		fmt.Println("Hello. Date of Birth?")
		dateBirth, _ := reader.ReadString('\n')
		trimmedDateBirth := strings.TrimSpace(dateBirth);
		students[i].dateBirth = trimmedDateBirth

		fmt.Println("Hello. Age?")
		age, _ := reader.ReadString('\n')
		trimmedAge := strings.TrimSpace(age);
		students[i].age = trimmedAge
	}

	for _, student := range students {
		fmt.Println(student)
	}
}
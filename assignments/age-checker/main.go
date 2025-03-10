package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bearbin/go-age"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello. Date of birth?")
	response, _ := reader.ReadString('\n')
	trimmedResponse := strings.TrimSpace(response);
	parsedDate, parsedDateErr := time.Parse("2006-01-02", trimmedResponse)

	if parsedDateErr != nil {
		fmt.Println("Error parsing date:", parsedDateErr)
		return
	}

	parsedDate = time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, parsedDate.Location())

 	fmt.Println(age.AgeAt(parsedDate, time.Now()))
}
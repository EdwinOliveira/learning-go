package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addDigits(firstDigit, secondDigit, thirdDigit int) int {
	return firstDigit + secondDigit + thirdDigit
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	numberStorage := []int{}

	fmt.Println("Hello. Please input 3 single digits?")
	response, _ := reader.ReadString('\n')
	trimmedResponse := strings.TrimSpace(response);

	for _, number := range strings.Split(trimmedResponse, "") {
		parsedResponse, parsedResponseError := strconv.Atoi(number)

		if parsedResponseError != nil {
			fmt.Println("Error converting string to integer:", parsedResponseError)
			return
		}

		numberStorage = append(numberStorage, parsedResponse);
	}
}
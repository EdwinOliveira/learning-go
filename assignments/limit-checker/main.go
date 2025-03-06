package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello. Please input your number?")
	response, _ := reader.ReadString('\n')

	parsedResponse, parsedResponseError := strconv.Atoi(strings.TrimSpace(response))

	if parsedResponseError != nil {
		fmt.Println("Error converting string to integer:", parsedResponseError)
		return
	}

	if parsedResponse >= 1 && parsedResponse <= 10 {
		fmt.Println("Give number lies between 1 and 10")
	} else {
		fmt.Println("Give number doesn't lies between 1 and 10")
	}

}
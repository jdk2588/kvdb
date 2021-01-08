package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := Inmemory{sets: map[string]string{}}
	fmt.Println("Type `exit` to quit!")
	var multi [][]string
	for {
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		str := strings.Trim(input, "\n")
		if str == "exit" {
			break
		}

		arr := strings.Split(str, " ")

		output := batchCommand(arr, s, &multi)
		if len(output) > 0 {
			fmt.Println(output)
		}
	}
}
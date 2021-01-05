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
	s := Inmemory{sets: map[string]string{}}
	fmt.Println("Type `exit` to quit!")
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

		if len(arr) < 2 {
			fmt.Println("Minimum two arguments are required!")
			break
		}

		switch arr[0] {
		case "SET":
			s.add(arr[1], arr[2])
		case "DEL":
			s.delete(arr[1])
		case "GET":
			val, e := s.get(arr[1])
			if e == nil {
				fmt.Println(val)
			} else {
				fmt.Println(e.Error())
			}
		case "INCR":
			if _, e := s.increment(arr[1]); e != nil {
				fmt.Println(e.Error())
			}
		case "INCRBY":
			v, err := strconv.Atoi(arr[2])
			if err == nil {
				if _, err := s.incrementBy(arr[1], v); err != nil {
					fmt.Println(err.Error())
				}
			} else {
				fmt.Printf("%s does not look like a counter!", arr[1])
			}
		default:
			fmt.Println("Not a valid command!")
		}
	}
}
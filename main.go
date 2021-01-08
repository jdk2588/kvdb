package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkCommand(arr []string, s Inmemory) {
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
	case "COMPACT":
		for k, v := range s.sets {
			fmt.Printf("SET %s %s\n", k, v)
		}
	default:
		fmt.Println("Not a valid command!")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := Inmemory{sets: map[string]string{}}
	fmt.Println("Type `exit` to quit!")
	multi := [][]string{}
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

		//if len(arr) < 2 {
		//	fmt.Println("Minimum two arguments are required!")
		//	break
		//}

		switch arr[0] {
		case "MULTI":
			multi = append(multi, arr)
		case "DISCARD":
			multi = [][]string{}
		case "EXEC":
			for _, i := range multi {
				if i[0] != "MULTI" {
					checkCommand(i, s)
				}
			}
			multi = [][]string{}
		default:
			if len(multi) > 0 {
				multi = append(multi, arr)
			} else {
				checkCommand(arr, s)
			}
		}
	}
}
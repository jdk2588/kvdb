package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getKey(s Inmemory, key string) string {
	val, e := s.get(key)
	var output string
	if e == nil {
		output = val
	} else {
		output = e.Error()
	}
	return output
}
func checkCommand(arr []string, s Inmemory) string {
	var output string
	switch arr[0] {
	case "SET":
		s.add(arr[1], arr[2])
	case "DEL":
		s.delete(arr[1])
	case "GET":
			fmt.Println(getKey(s, arr[1]))
	case "INCR":
		if _, e := s.increment(arr[1]); e != nil {
			fmt.Println(e.Error())
		}
	case "INCRBY":
		v, err := strconv.Atoi(arr[2])
		if err == nil {
			if _, err := s.incrementBy(arr[1], v); err != nil {
				output = err.Error()
			}
		} else {
			output = fmt.Sprintf("%s does not look like a counter!", arr[1])
		}

	default:
		fmt.Println("Not a valid command!")
	}

	return output
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

		switch arr[0] {
		case "MULTI":
			multi = append(multi, arr)
		case "DISCARD":
			multi = [][]string{}
		case "EXEC":
			output := []string{}
			for _, i := range multi {
				if i[0] != "MULTI" {
					checkCommand(i, s)
					output = append(output, getKey(s, i[1]))
				}
			}
			fmt.Println(output)
			multi = [][]string{}
		case "COMPACT":
			for k, v := range s.sets {
				fmt.Printf("SET %s %s\n", k, v)
			}
		default:
			if len(multi) > 0 {
				multi = append(multi, arr)
			} else {
				checkCommand(arr, s)
			}
		}
	}
}
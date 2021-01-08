package main

import (
	"fmt"
	"strconv"
)

func getKey(s Storer, key string) string {
	val, e := s.get(key)
	var output string
	if e == nil {
		output = val
	} else {
		output = e.Error()
	}
	return output
}

func checkCommand(arr []string, s Storer) string {
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

func batchCommand(arr []string, s Storer, multi *[][]string) []string {

	var output []string
	switch arr[0] {
	case "MULTI":
		*multi = append(*multi, arr)
	case "DISCARD":
		*multi = [][]string{}
	case "EXEC":
			o := []string{}
			for _, i := range *multi {
				if i[0] != "MULTI" {
					checkCommand(i, s)
					o = append(o, getKey(s, i[1]))
				}
			}
			output = o
			*multi = [][]string{}
		case "COMPACT":
			for k, v := range s.getAll() {
				output = append(output, fmt.Sprintf("SET %s %s\n", k, v))
			}
	default:
		if len(*multi) > 0 {
			*multi = append(*multi, arr)
		} else {
			checkCommand(arr, s)
		}
	}

	return output
}

package main

import (
	"fmt"
	"strconv"
)

type Inmemory struct {
	sets map[string]string
}

func (in Inmemory) increment(k string, add int) (bool, error) {
	v, e := in.get(k)
	if e != nil {
		return false, e
	}

	if n, e := strconv.Atoi(v); e == nil {
		in.add(k, strconv.Itoa(n+add))
		return true, nil
	} else {
		return false, e
	}
}

func (in Inmemory) add(k string, v string) bool {
	in.sets[k] = v
	return true
}

func (in Inmemory) get(k string) (string, error) {
	val := in.sets[k]

	if val == "" {
		return val, fmt.Errorf("value not found for the key %s", k)
	}

	return val, nil
}

func (in Inmemory) delete(k string) (bool, error) {
	_, ok := in.sets[k]
	if ok {
		delete(in.sets, k)
		return true, nil
	}

	return false, fmt.Errorf("key %s does not exists", k)
}



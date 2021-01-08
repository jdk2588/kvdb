package main

type Storer interface {
	add(string, string) bool
	get(string) (string, error)
	delete(string) (bool, error)
	incrementBy(string, int) (bool, error)
	increment(string) (bool, error)
	getAll() map[string]string
}
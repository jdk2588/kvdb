package main

type Storer interface {
	add(string, string) bool
	get(string) (string, error)
	delete(string) (bool, error)
	increment(string, int) (bool, error)
}
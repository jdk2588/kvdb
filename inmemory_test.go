package main

import "testing"

var s Storer

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	s = Inmemory{sets: map[string]string{}}
}

func TestAdd(t *testing.T) {
	ret := s.add("key1", "value1")

	if ret != true {
		t.Errorf("Could not add key1 and value1.")
	}
}

func TestGet(t *testing.T) {
	_, e := s.get("key1")
	if e != nil {
		t.Errorf("The key stored is not key1.")
	}

	_, e1 := s.get("key2")
	if e1 == nil {
		t.Log("Key not found key2.")
	}
}

func TestDelete(t *testing.T) {
	b, _ := s.delete("key1")
	if b {
		t.Log("key1 deleted")
	}

	_, e2 := s.delete("key2")
	if e2 != nil {
		t.Log(e2.Error())
	}
}

func TestAddCounter(t *testing.T) {
	s.add("counter", "0")
	s.increment("counter", 11)
	s.increment("counter", 12)
	if v, e := s.get("counter"); e == nil {
		if v == "23" {
			t.Logf("Value for counter is %s", v)
		} else {
			t.Errorf("Value mismatch. Expected 23, Got %s", v)
		}
	}
}
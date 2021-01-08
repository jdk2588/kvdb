package main

import (
	"testing"
)

func TestMulti(t *testing.T) {
	multi := [][]string{}
	multi = append(multi, []string{"MULTI"})
	batchCommand([]string{"MULTI"}, s, &multi)
	batchCommand([]string{"INCR", "foo"}, s, &multi)
	batchCommand([]string{"SET", "key1", "value1"}, s, &multi)
	output := batchCommand([]string{"EXEC"}, s, &multi)
	if output[0] != "1" && output[2] != "value1" {
		t.Errorf("Output not matching")
	}
}

func TestMultiDiscard(t *testing.T) {
	multi := [][]string{}
	multi = append(multi, []string{"MULTI"})
	batchCommand([]string{"MULTI"}, s, &multi)
	batchCommand([]string{"INCR", "foo"}, s, &multi)
	batchCommand([]string{"DISCARD"}, s, &multi)
	batchCommand([]string{"GET", "key1"}, s, &multi)
	output := batchCommand([]string{"EXEC"}, s, &multi)
	if len(output) != 0 {
		t.Errorf("Discard did not work as expected!")
	}
}

func TestCompact(t *testing.T) {
	multi := [][]string{}
	batchCommand([]string{"SET", "counter", "10"}, s, &multi)
	batchCommand([]string{"INCR", "counter"}, s, &multi)
	batchCommand([]string{"INCR", "counter"}, s, &multi)
	batchCommand([]string{"SET", "foo", "bar"}, s, &multi)
	batchCommand([]string{"GET", "counter"}, s, &multi)
	batchCommand([]string{"INCR", "counter"}, s, &multi)
	output := batchCommand([]string{"COMPACT"}, s, &multi)
	if len(output) != 2 {
		t.Errorf("Compact did not work as expected!")
	}
}

func TestCompactNone(t *testing.T) {
	multi := [][]string{}
	batchCommand([]string{"INCR", "counter"}, s, &multi)
	batchCommand([]string{"INCRBY", "counter", "10"}, s, &multi)
	batchCommand([]string{"GET", "counter"}, s, &multi)
	batchCommand([]string{"DEL", "counter"}, s, &multi)
	output := batchCommand([]string{"COMPACT"}, s, &multi)
	if len(output) != 0 {
		t.Errorf("Compact did not work as expected!")
	}
}
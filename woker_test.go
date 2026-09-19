package main

import (
	"testing"
	"time"
)

// TestCreateWorkerDeliversValue is the one thing this repository had no test
// for at all: that the worker goroutine in woker.go actually runs the work
// function it is given, with the value sent on the channel it returns. main.go
// depends on this to report progress while it checks 200,000,000 UUIDs; if
// createWorker stopped delivering, or delivered the wrong value, this fails.
func TestCreateWorkerDeliversValue(t *testing.T) {
	received := make(chan int, 1)
	work := createWorker(func(number int) {
		received <- number
	})

	work <- 42

	select {
	case got := <-received:
		if got != 42 {
			t.Fatalf("createWorker delivered %d, want 42", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("createWorker did not run the work function within 2s")
	}
}

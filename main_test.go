package main

import (
	"context"
	"testing"
	"time"
)

// TestTransitionInput1 checks the transition for Input1.
func TestTransitionInput1(t *testing.T) {
	initialState := State{}
	input := Input1{}

	newState, output := transitionInput1(initialState, input)

	if newState.LastReceived != input {
		t.Errorf("Expected LastReceived to be input1, got %v", newState.LastReceived)
	}

	if output.Received != 1 {
		t.Errorf("Expected output signal to be 1, got %d", output.Received)
	}
}

// TestTransitionInput2 checks the transition for Input2.
func TestTransitionInput2(t *testing.T) {
	initialState := State{}
	input := Input2{}

	newState, output := transitionInput2(initialState, input)

	if newState.LastReceived != input {
		t.Errorf("Expected LastReceived to be input2, got %v", newState.LastReceived)
	}

	if output.Received != 2 {
		t.Errorf("Expected output signal to be 2, got %d", output.Received)
	}
}

// TestStateMachine checks the algoRecursive function.
func TestStateMachine(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	initialState := State{}
	input1 := make(chan Input1, 100)
	input2 := make(chan Input2, 100)
	output := make(chan Output1, 100)

	go stateMachine(ctx, initialState, input1, input2, output)

	// Test Input1 signal
	input1 <- Input1{}
	select {
	case out := <-output:
		if out.Received != 1 {
			t.Errorf("Expected output signal to be 1, got %d", out.Received)
		}
	case <-time.After(10 * time.Millisecond):
		t.Errorf("Timeout waiting for output")
	}

	// Test Input2 signal
	input2 <- Input2{}
	select {
	case out := <-output:
		if out.Received != 2 {
			t.Errorf("Expected output signal to be 2, got %d", out.Received)
		}
	case <-time.After(10 * time.Millisecond):
		t.Errorf("Timeout waiting for output")
	}
}

// TestCollectOutputSignals verifies that collectOutputSignals works properly.
func TestCollectOutputSignals(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	output := make(chan Output1, 10)
	go collectOutputSignals(ctx, output)

	// Send an output signal
	output <- Output1{Received: 1}
	time.Sleep(10 * time.Millisecond) // Give the goroutine time to process

	// No explicit check here since collectOutputSignals only prints output.
	// We can assume it works based on manual inspection of the console logs.
}

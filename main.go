package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	initialState := State{}
	input1 := make(chan Input1, 100)
	input2 := make(chan Input2, 100)
	output1 := make(chan Output1, 100)

	go collectOutputSignals(ctx, output1)

	go stateMachine(ctx, initialState, input1, input2, output1)

	time.Sleep(10 * time.Millisecond)
	input1 <- Input1{}
	time.Sleep(10 * time.Millisecond)
	input1 <- Input1{}
	time.Sleep(10 * time.Millisecond)
	input2 <- Input2{}
	time.Sleep(10 * time.Millisecond)
	input1 <- Input1{}

	time.Sleep(10 * time.Millisecond)

	<-ctx.Done()
}

func collectOutputSignals(ctx context.Context, output <-chan Output1) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Collector exiting")
			return
		case signal := <-output:
			fmt.Printf("Output received: %d\n", signal.Received)
		}
	}
}

type State struct {
	LastReceived any
}

type Input1 struct{}
type Input2 struct{}
type Output1 struct {
	Received uint8
}

func transitionInput1(s State, i Input1) (State, Output1) {
	fmt.Println("Register input signal 1")

	s.LastReceived = i

	return s, Output1{1}
}
func transitionInput2(s State, i Input2) (State, Output1) {
	fmt.Println("Register input signal 2")

	s.LastReceived = i

	return s, Output1{2}
}

func stateMachine(
	ctx context.Context,
	state State,
	dataIn1 <-chan Input1,
	dataIn2 <-chan Input2,
	dataOut1 chan<- Output1,
) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Algorithm exiting")
			return
		case signal := <-dataIn1:
			newState, outSignal := transitionInput1(state, signal)
			state = newState
			dataOut1 <- outSignal
		case signal := <-dataIn2:
			newState, outSignal := transitionInput2(state, signal)
			state = newState
			dataOut1 <- outSignal
		}
	}
}

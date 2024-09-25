# State Machine Example in Go

This project demonstrates a simple state machine implementation using Go, with recursive state transitions based on input signals.

## Overview

The core of the project is a function called `stateMachine`, which processes two types of input signals (`Input1` and `Input2`) and transitions between states. The output is emitted as a response to the signals, with the state being updated at each step.

### Key Components
- **State**: Tracks the current state of the system.
- **Input1 / Input2**: Input signals that trigger state transitions.
- **Output1**: Output signal emitted after a state transition.
- **stateMachine**: Recursively processes inputs, updates the state, and generates outputs.
- **collectOutputSignals**: Collects and prints output signals for display purposes.

## Usage

1. Run the `main.go` file to see the state machine in action.

```sh
go run main.go
```

2. Run the tests.

```sh
go test
```

package main

import "fmt"

type serverState int

// The iota is a keyword that automatically generates ints from 0 to 1 for each element of the const
const (
	StateIdle serverState = iota
	StateConnected
	StateError
	StateReconnect
)

var stateName = map[serverState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateReconnect: "reconnecting",
}

func (ss serverState) String() string {
	return stateName[ss]
}

func learnEnums() {
	ns := transition(StateIdle)

	fmt.Println(ns)

	ns2 := transition(ns)

	fmt.Println(ns2)

}
func transition(s serverState) serverState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateReconnect:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Unknown State %s", s))
	}
}

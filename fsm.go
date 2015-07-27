package fsm

import (
	"errors"
	"fmt"
)

// FSMs are the "interface" to manage state. They should not be
// created directly. Rather, the should make made by blueprints.
type Machine struct {
	// List of available transitions
	transitions list
	// Current fsm state
	state uint8
}

// Returns whether a transition to state s is legal.
func (f *Machine) isLegal(a uint8, b uint8) bool {
	return f.transitions.Search(serialize(a, b)) != nil
}

// Returns whether this can transition to the target state.
func (f *Machine) Allows(b uint8) bool {
	return f.isLegal(f.state, b)
}

// Returns whether this is disallowed to transition to the target state.
func (f *Machine) Disallows(b uint8) bool {
	return !f.Allows(b)
}

// Returns the current fsm state.
func (f *Machine) State() uint8 {
	return f.state
}

// Moves the fsm to the target state. Panics if disallowed.
func (f *Machine) Goto(state uint8) error {
	t := f.transitions.Search(serialize(f.state, state))
	if t == nil {
		return errors.New(fmt.Sprintf("Transition %d to %d not permitted.", f.state, state))
	}

	f.state = state
	if t.fn != nil {
		t.fn(f)
	}

	return nil
}

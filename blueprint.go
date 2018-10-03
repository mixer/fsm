package fsm

type Blueprint struct {
	transitions list
	start       uint8
}

// New creates a new finite state machine blueprint.
func New() *Blueprint {
	return &Blueprint{
		transitions: make(list, 0),
	}
}

// From returns a new transition for the blueprint.
// The transition will be added to the blueprint automatically when it has both
// "from" and "to" values.
func (b *Blueprint) From(start uint8) *Transition {
	return (&Transition{blueprint: b}).From(start)
}

// Add adds a complete transition to the blueprint.
func (b *Blueprint) Add(t *Transition) {
	idx := b.transitions.InsertPos(t)
	trans := make(list, len(b.transitions)+1)

	copy(trans, b.transitions[:idx])
	copy(trans[idx+1:], b.transitions[idx:])
	trans[idx] = t
	b.transitions = trans
}

// Start sets the start state for the machine.
func (b *Blueprint) Start(state uint8) {
	b.start = state
}

// Machine returns a new machine created from the blueprint.
func (b *Blueprint) Machine() *Machine {
	fsm := &Machine{
		state:       b.start,
		transitions: b.transitions,
	}

	return fsm
}

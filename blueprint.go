package fsm

type Blueprint struct {
	transitions list
	start       uint8
}

// Creates a new blueprint, from which fsms can be made.
func New() *Blueprint {
	return &Blueprint{
		transitions: make(list, 0),
	}
}

// Starts creating a new transition on the blueprint.
func (b *Blueprint) From(start uint8) *Transition {
	return (&Transition{blueprint: b}).From(start)
}

// Adds an already complete transition to the blueprint.
func (b *Blueprint) Add(t *Transition) {
	idx := b.transitions.InsertPos(t)
	trans := make(list, len(b.transitions)+1)

	copy(trans, b.transitions[:idx])
	copy(trans[idx+1:], b.transitions[idx:])
	trans[idx] = t
	b.transitions = trans
}

// Marks the state state of the fsm.
func (b *Blueprint) Start(state uint8) {
	b.start = state
}

// Creates a new FSM from the blueprint.
func (b *Blueprint) Machine() *Machine {
	fsm := &Machine{
		state:       b.start,
		transitions: b.transitions,
	}

	return fsm
}

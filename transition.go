package fsm

import (
	"sort"
)

// Handler represents a callback to be called when the machine performs a
// certain transition between states.
type Handler func(m *Machine)

// Transition represents a transition between two states.
type Transition struct {
	start     bool
	from      uint8
	fromSet   bool
	to        uint8
	toSet     bool
	hash      uint16
	blueprint *Blueprint
	fn        func(m *Machine)
}

// From sets the source state of the transition.
func (t *Transition) From(from uint8) *Transition {
	t.from = from
	t.fromSet = true
	t.recalculate()
	return t
}

// To sets the destination state of the transition.
func (t *Transition) To(to uint8) *Transition {
	t.to = to
	t.toSet = true
	t.recalculate()
	return t
}

// recalculate calculates the hash for this transition if both "from" and "to"
// have been set. If both "from" and "to" are set then this transition will
// also be added to the blueprint.
func (t *Transition) recalculate() {
	if !t.toSet || !t.fromSet {
		return
	}

	t.hash = serialize(t.from, t.to)
	t.blueprint.Add(t)
}

// Then sets the callback function for when the transition has occurred.
func (t *Transition) Then(fn Handler) *Transition {
	t.fn = fn
	return t
}

// serialize serializes a transition between two states into a single value,
// where the first 8 bits are the "from" and the last 8 bits are the "to"
// state.
func serialize(from, to uint8) uint16 {
	return (uint16(from) << 8) | uint16(to)
}

// list represents a list of transitions.
type list []*Transition

// Len returns the length of the list.
func (t list) Len() int {
	return len(t)
}

// Swap swaps the two elements with indexes a and b.
func (t list) Swap(a, b int) {
	t[a], t[b] = t[b], t[a]
}

// Less returns whether the element at index a should appear before b.
func (t list) Less(a, b int) bool {
	return t[a].hash < t[b].hash
}

// Search searches for the specified hash in the list and returns it if it is
// present.
func (t list) Search(x uint16) *Transition {
	low, high := 0, len(t)-1
	for low <= high {
		i := (low + high) / 2
		if t[i].hash > x {
			high = i - 1
		} else if t[i].hash < x {
			low = i + 1
		} else {
			return t[i]
		}
	}

	return nil
}

// InsertPos returns the index at which the specified transition should be
// inserted into the slice to retain it's order.
func (t list) InsertPos(v *Transition) int {
	return sort.Search(len(t), func(i int) bool {
		return t[i].hash >= v.hash
	})
}

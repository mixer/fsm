package fsm

import (
	"sort"
)

type Handler func(m *Machine)

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

// Sets the "from" transition
func (t *Transition) From(from uint8) *Transition {
	t.from = from
	t.fromSet = true
	t.recalculate()
	return t
}

// Sets the target state.
func (t *Transition) To(to uint8) *Transition {
	t.to = to
	t.toSet = true
	t.recalculate()
	return t
}

// Calculates the hash if both `from` and `to` have been set,
// then adds itself to the blueprint.
func (t *Transition) recalculate() {
	if !t.toSet || !t.fromSet {
		return
	}

	t.hash = serialize(t.from, t.to)
	t.blueprint.Add(t)
}

// Sets the transition handler function.
func (t *Transition) Then(fn Handler) *Transition {
	t.fn = fn
	return t
}

// Serializes the transition to a uint16, where the first 8 bits
// are the "from" and the last 8 bits are the "to" state
func serialize(from, to uint8) uint16 {
	return (uint16(from) << 8) | uint16(to)
}

type list []*Transition

// Implements sort.Interface.Len
func (t list) Len() int {
	return len(t)
}

// Implements sort.Interface.Swap
func (t list) Swap(a, b int) {
	t[a], t[b] = t[b], t[a]
}

// Implements sort.Interface.Less
func (t list) Less(a, b int) bool {
	return t[a].hash < t[b].hash
}

// Searches for the value in the slice. If it's not found, then
// -1 is returned. Otherwise, its index is returned.
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

// Returns the position in the slice that a new transition should
// be inserted at to preseve its sortedness.
func (t list) InsertPos(v *Transition) int {
	return sort.Search(len(t), func(i int) bool {
		return t[i].hash >= v.hash
	})
}

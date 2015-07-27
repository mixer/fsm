package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	A uint8 = iota
	B
	C
	D
)

func makeMachine() *Machine {
	bp := New()
	bp.Start(A)
	bp.From(A).To(B)
	bp.From(B).To(C)
	bp.From(B).To(B)
	return bp.Machine()
}

func TestWorksNormally(t *testing.T) {
	m := makeMachine()
	assert.Equal(t, A, m.State())
	assert.Nil(t, m.Goto(B))
	assert.Equal(t, B, m.State())
	assert.Nil(t, m.Goto(C))
	assert.Equal(t, C, m.State())
	err := m.Goto(B)
	assert.NotNil(t, err)
	assert.Equal(t, C, m.State())
	assert.Equal(t, "Transition 2 to 1 not permitted.", err.Error())
}

func TestAddsHandler(t *testing.T) {
	bp := New()
	out := []uint8{}
	bp.From(A).To(B).Then(func(m *Machine) { out = append(out, 1) })
	bp.From(B).To(C)
	bp.From(C).To(D).Then(func(m *Machine) { out = append(out, 2) })
	m := bp.Machine()

	assert.Equal(t, []uint8{}, out)
	m.Goto(B)
	assert.Equal(t, []uint8{1}, out)
	m.Goto(C)
	assert.Equal(t, []uint8{1}, out)
	m.Goto(D)
	assert.Equal(t, []uint8{1, 2}, out)
}

func BenchmarkTransitions(b *testing.B) {
	m := makeMachine()
	m.Goto(B)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.Goto(B)
	}
}

func BenchmarkAllows(b *testing.B) {
	m := makeMachine()
	for n := 0; n < b.N; n++ {
		m.Allows(B)
	}
}

func BenchmarkGetState(b *testing.B) {
	m := makeMachine()
	for n := 0; n < b.N; n++ {
		m.State()
	}
}

func BenchmarkBuildMachine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		makeMachine()
	}
}

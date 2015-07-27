# fsm [![GoDoc](https://godoc.org/github.com/WatchBeam/fsm?status.svg)](https://godoc.org/github.com/WatchBeam/fsm) [![Build Status](https://travis-ci.org/WatchBeam/fsm.svg)](https://travis-ci.org/WatchBeam/fsm)


Why another FSM implementation? Because we didn't see one suited for smaller-scale, programmatic use in Go which was very efficient. Example use:

```go
bp := fsm.New()
bp.Start(0)
bp.From(0).To(1)
bp.From(1).To(2).Then(func (m *fsm.Machine) { fmt.Println("hola!") })

m := bp.Machine()
m.State() // => a
m.Goto(1) // => error(nil)
m.State() // => b
m.Goto(2) // => error(nil)
// => hola!
m.Goto(1) // => error, "Transition 2 to 1 not permitted."
```

See the godocs for more information.

Benchmarks well, especially against [comparable solutions](https://github.com/ryanfaerman/fsm#benchmarks).

```
➜  fsm git:(fsm) ✗ go test -bench=.
PASS
PASS
BenchmarkTransitions    100000000           20.8 ns/op
BenchmarkAllows          50000000           20.6 ns/op
BenchmarkGetState      2000000000           0.48 ns/op
```

package step

type Value interface {
	NextStep() []Value
	Finished() bool
	Print()
}

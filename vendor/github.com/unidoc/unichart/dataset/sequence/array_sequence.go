package sequence

var (
	_ Sequence = (*ArraySequence)(nil)
)

// ArraySequence is a wrapper for an array of floats that implements
// the `Sequence` interface.
type ArraySequence []float64

// NewArraySequence returns a new array from a given set of values.
// ArraySequence implements Sequence, which allows it to be used with
// the sequence helpers.
func NewArraySequence(values ...float64) ArraySequence {
	return ArraySequence(values)
}

// NewArrayWrapper returns a new array sequence wrapper for a given values set.
func NewArrayWrapper(values ...float64) Wrapper {
	return Wrapper{NewArraySequence(values...)}
}

// Len returns the value provider length.
func (a ArraySequence) Len() int {
	return len(a)
}

// GetValue returns the value at a given index.
func (a ArraySequence) GetValue(index int) float64 {
	return a[index]
}

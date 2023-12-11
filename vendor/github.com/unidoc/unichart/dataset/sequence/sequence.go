package sequence

// Sequence is a provider for values for a seq.
type Sequence interface {
	Len() int
	GetValue(int) float64
}

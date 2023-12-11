package sequence

// Range is a common interface for a range of values.
type Range interface {
	GetMin() float64
	SetMin(min float64)

	GetMax() float64
	SetMax(max float64)

	GetDelta() float64

	GetDomain() int
	SetDomain(domain int)

	IsDescending() bool

	// Translate translates the range to the domain.
	Translate(value float64) int

	String() string

	IsZero() bool
}

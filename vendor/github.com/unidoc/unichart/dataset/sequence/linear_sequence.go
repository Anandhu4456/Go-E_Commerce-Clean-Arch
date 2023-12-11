package sequence

// Assert types implement interfaces.
var (
	_ Sequence = (*LinearSequence)(nil)
)

// LinearRange returns an array of values representing the range from
// start to end, incremented by 1.0.
func LinearRange(start, end float64) []float64 {
	return Wrapper{NewLinearSequence().WithStart(start).WithEnd(end).WithStep(1.0)}.Values()
}

// LinearRangeWithStep returns the array values of a linear sequence with
// a given start, end and optional step.
func LinearRangeWithStep(start, end, step float64) []float64 {
	return Wrapper{NewLinearSequence().WithStart(start).WithEnd(end).WithStep(step)}.Values()
}

// LinearSequence is a stepwise sequence.
type LinearSequence struct {
	start float64
	end   float64
	step  float64
}

// NewLinearSequence returns a new linear sequence.
func NewLinearSequence() *LinearSequence {
	return &LinearSequence{step: 1.0}
}

// Start returns the start value.
func (l LinearSequence) Start() float64 {
	return l.start
}

// End returns the end value.
func (l LinearSequence) End() float64 {
	return l.end
}

// Step returns the step value.
func (l LinearSequence) Step() float64 {
	return l.step
}

// Len returns the number of elements in the sequence.
func (l LinearSequence) Len() int {
	if l.start < l.end {
		return int((l.end-l.start)/l.step) + 1
	}
	return int((l.start-l.end)/l.step) + 1
}

// GetValue returns the value at a given index.
func (l LinearSequence) GetValue(index int) float64 {
	fi := float64(index)
	if l.start < l.end {
		return l.start + (fi * l.step)
	}
	return l.start - (fi * l.step)
}

// WithStart sets the start and returns the linear generator.
func (l *LinearSequence) WithStart(start float64) *LinearSequence {
	l.start = start
	return l
}

// WithEnd sets the end and returns the linear generator.
func (l *LinearSequence) WithEnd(end float64) *LinearSequence {
	l.end = end
	return l
}

// WithStep sets the step and returns the linear generator.
func (l *LinearSequence) WithStep(step float64) *LinearSequence {
	l.step = step
	return l
}

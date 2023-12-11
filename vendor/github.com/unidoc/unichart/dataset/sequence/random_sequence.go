package sequence

import (
	"math"
	"math/rand"
	"time"
)

var (
	_ Sequence = (*RandomSequence)(nil)
)

// RandomValues returns an array of random values.
func RandomValues(count int) []float64 {
	return Wrapper{NewRandomSequence().WithLen(count)}.Values()
}

// RandomValuesWithMax returns an array of random values with a given average.
func RandomValuesWithMax(count int, max float64) []float64 {
	return Wrapper{NewRandomSequence().WithMax(max).WithLen(count)}.Values()
}

// RandomSequence is a random number sequence.
type RandomSequence struct {
	rnd    *rand.Rand
	max    *float64
	min    *float64
	length *int
}

// NewRandomSequence returns a new random sequence.
func NewRandomSequence() *RandomSequence {
	return &RandomSequence{
		rnd: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Len returns the number of elements that will be generated.
func (r *RandomSequence) Len() int {
	if r.length != nil {
		return *r.length
	}
	return math.MaxInt32
}

// GetValue returns the value.
func (r *RandomSequence) GetValue(_ int) float64 {
	if r.min != nil && r.max != nil {
		var delta float64
		if *r.max > *r.min {
			delta = *r.max - *r.min
		} else {
			delta = *r.min - *r.max
		}

		return *r.min + (r.rnd.Float64() * delta)
	} else if r.max != nil {
		return r.rnd.Float64() * *r.max
	} else if r.min != nil {
		return *r.min + (r.rnd.Float64())
	}

	return r.rnd.Float64()
}

// WithLen sets a maximum len.
func (r *RandomSequence) WithLen(length int) *RandomSequence {
	r.length = &length
	return r
}

// Min returns the minimum value.
func (r RandomSequence) Min() *float64 {
	return r.min
}

// WithMin sets the scale and returns the RandomSequence.
func (r *RandomSequence) WithMin(min float64) *RandomSequence {
	r.min = &min
	return r
}

// Max returns the maximum value.
func (r RandomSequence) Max() *float64 {
	return r.max
}

// WithMax sets the average and returns the RandomSequence.
func (r *RandomSequence) WithMax(max float64) *RandomSequence {
	r.max = &max
	return r
}

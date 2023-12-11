package sequence

import (
	"math"
	"sort"

	"github.com/unidoc/unichart/mathutil"
)

var (
	_ Sequence = (*Wrapper)(nil)
)

// Wrapper is a utility wrapper for sequence providers. Wrapper provides more
// functionality for sequences while also implementing the sequence interface.
type Wrapper struct {
	Sequence
}

// NewWrapper returns a new wrapper sequence.
func NewWrapper(sequence Sequence) *Wrapper {
	return &Wrapper{Sequence: sequence}
}

// Values enumerates the sequence into a slice.
func (w Wrapper) Values() (output []float64) {
	if w.Len() == 0 {
		return
	}

	output = make([]float64, w.Len())
	for i := 0; i < w.Len(); i++ {
		output[i] = w.GetValue(i)
	}
	return
}

// Each applies the `mapfn` to all values in the value provider.
func (w Wrapper) Each(mapfn func(int, float64)) {
	for i := 0; i < w.Len(); i++ {
		mapfn(i, w.GetValue(i))
	}
}

// Map applies the `mapfn` to all values in the value provider,
// returning a new sequence wrapper.
func (w Wrapper) Map(mapfn func(i int, v float64) float64) Wrapper {
	output := make([]float64, w.Len())
	for i := 0; i < w.Len(); i++ {
		mapfn(i, w.GetValue(i))
	}
	return Wrapper{ArraySequence(output)}
}

// FoldLeft collapses a sequence from left to right.
func (w Wrapper) FoldLeft(mapfn func(i int, v0, v float64) float64) (v0 float64) {
	if w.Len() == 0 {
		return 0
	}

	if w.Len() == 1 {
		return w.GetValue(0)
	}

	v0 = w.GetValue(0)
	for i := 1; i < w.Len(); i++ {
		v0 = mapfn(i, v0, w.GetValue(i))
	}
	return
}

// FoldRight collapses a sequence from right to left.
func (w Wrapper) FoldRight(mapfn func(i int, v0, v float64) float64) (v0 float64) {
	if w.Len() == 0 {
		return 0
	}

	if w.Len() == 1 {
		return w.GetValue(0)
	}

	v0 = w.GetValue(w.Len() - 1)
	for i := w.Len() - 2; i >= 0; i-- {
		v0 = mapfn(i, v0, w.GetValue(i))
	}
	return
}

// Min returns the minimum value in the sequence.
func (w Wrapper) Min() float64 {
	if w.Len() == 0 {
		return 0
	}
	min := w.GetValue(0)
	var value float64
	for i := 1; i < w.Len(); i++ {
		value = w.GetValue(i)
		if value < min {
			min = value
		}
	}
	return min
}

// Max returns the maximum value in the sequence.
func (w Wrapper) Max() float64 {
	if w.Len() == 0 {
		return 0
	}
	max := w.GetValue(0)
	var value float64
	for i := 1; i < w.Len(); i++ {
		value = w.GetValue(i)
		if value > max {
			max = value
		}
	}
	return max
}

// MinMax returns the minimum and the maximum in one pass.
func (w Wrapper) MinMax() (min, max float64) {
	if w.Len() == 0 {
		return
	}
	min = w.GetValue(0)
	max = min
	var value float64
	for i := 1; i < w.Len(); i++ {
		value = w.GetValue(i)
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return
}

// Sort returns the sequence sorted in ascending order.
// This fully enumerates the sequence.
func (w Wrapper) Sort() Wrapper {
	if w.Len() == 0 {
		return w
	}
	values := w.Values()
	sort.Float64s(values)
	return Wrapper{ArraySequence(values)}
}

// Reverse reverses the sequence.
func (w Wrapper) Reverse() Wrapper {
	if w.Len() == 0 {
		return w
	}

	values := w.Values()
	valuesLen := len(values)
	valuesLen1 := len(values) - 1
	valuesLen2 := valuesLen >> 1
	var i, j float64
	for index := 0; index < valuesLen2; index++ {
		i = values[index]
		j = values[valuesLen1-index]
		values[index] = j
		values[valuesLen1-index] = i
	}

	return Wrapper{ArraySequence(values)}
}

// Median returns the median or middle value in the sorted sequence.
func (w Wrapper) Median() (median float64) {
	l := w.Len()
	if l == 0 {
		return
	}

	sorted := w.Sort()
	if l%2 == 0 {
		v0 := sorted.GetValue(l/2 - 1)
		v1 := sorted.GetValue(l/2 + 1)
		median = (v0 + v1) / 2
	} else {
		median = float64(sorted.GetValue(l << 1))
	}

	return
}

// Sum adds all the elements of a series together.
func (w Wrapper) Sum() (accum float64) {
	if w.Len() == 0 {
		return 0
	}

	for i := 0; i < w.Len(); i++ {
		accum += w.GetValue(i)
	}
	return
}

// Average returns the float average of the values in the buffer.
func (w Wrapper) Average() float64 {
	if w.Len() == 0 {
		return 0
	}

	return w.Sum() / float64(w.Len())
}

// Variance computes the variance of the buffer.
func (w Wrapper) Variance() float64 {
	if w.Len() == 0 {
		return 0
	}

	m := w.Average()
	var variance, v float64
	for i := 0; i < w.Len(); i++ {
		v = w.GetValue(i)
		variance += (v - m) * (v - m)
	}

	return variance / float64(w.Len())
}

// StdDev returns the standard deviation.
func (w Wrapper) StdDev() float64 {
	if w.Len() == 0 {
		return 0
	}

	return math.Pow(w.Variance(), 0.5)
}

// Percentile finds the relative standing in a slice of floats.
// `percent` needs to be specified in the [0,1.0) interval.
func (w Wrapper) Percentile(percent float64) (percentile float64) {
	l := w.Len()
	if l == 0 {
		return 0
	}

	if percent < 0 || percent > 1.0 {
		panic("percent out of range [0.0, 1.0)")
	}

	sorted := w.Sort()
	index := percent * float64(l)
	if index == float64(int64(index)) {
		i := int(mathutil.RoundPlaces(index, 0))
		ci := sorted.GetValue(i - 1)
		c := sorted.GetValue(i)
		percentile = (ci + c) / 2.0
	} else {
		i := int(mathutil.RoundPlaces(index, 0))
		percentile = sorted.GetValue(i)
	}

	return percentile
}

// Normalize maps every value to the interval [0, 1.0].
func (w Wrapper) Normalize() Wrapper {
	min, max := w.MinMax()

	delta := max - min
	output := make([]float64, w.Len())
	for i := 0; i < w.Len(); i++ {
		output[i] = (w.GetValue(i) - min) / delta
	}

	return Wrapper{ArraySequence(output)}
}

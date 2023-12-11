package sequence

import (
	"sort"
	"time"
)

// Assert types implement interfaces.
var (
	_ Sequence       = (*TimeSequence)(nil)
	_ sort.Interface = (*TimeSequence)(nil)
)

// TimeSequence are an array of times.
// It wraps the array with methods that implement the `Sequence` interface.
type TimeSequence []time.Time

// Array returns the times to an array.
func (t TimeSequence) Array() []time.Time {
	return []time.Time(t)
}

// Len returns the length of the array.
func (t TimeSequence) Len() int {
	return len(t)
}

// GetValue returns a value at an index as a time.
func (t TimeSequence) GetValue(index int) float64 {
	return ToFloat64(t[index])
}

// Swap implements sort.Interface.
func (t TimeSequence) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Less implements sort.Interface.
func (t TimeSequence) Less(i, j int) bool {
	return t[i].Before(t[j])
}

// ToFloat64 returns a float64 representation of a time.
func ToFloat64(t time.Time) float64 {
	return float64(t.UnixNano())
}

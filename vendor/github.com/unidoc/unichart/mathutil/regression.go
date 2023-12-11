package mathutil

import "errors"

var (
	// ErrPolyRegArraysSameLength is a common error.
	ErrPolyRegArraysSameLength = errors.New("polynomial array inputs must be the same length")
)

// PolyRegression returns the polynomial regression of a given degree over
// the given values.
func PolyRegression(xvalues, yvalues []float64, degree int) ([]float64, error) {
	if len(xvalues) != len(yvalues) {
		return nil, ErrPolyRegArraysSameLength
	}

	m := len(yvalues)
	n := degree + 1
	y := NewMatrix(m, 1, yvalues...)
	x := ZeroMatrix(m, n)

	for i := 0; i < m; i++ {
		ip := float64(1)
		for j := 0; j < n; j++ {
			x.Set(i, j, ip)
			ip *= xvalues[i]
		}
	}

	q, r := x.QR()
	qty, err := q.Transpose().Times(y)
	if err != nil {
		return nil, err
	}

	c := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		c[i] = qty.Get(i, 0)
		for j := i + 1; j < n; j++ {
			c[i] -= c[j] * r.Get(i, j)
		}
		c[i] /= r.Get(i, i)
	}

	return c, nil
}

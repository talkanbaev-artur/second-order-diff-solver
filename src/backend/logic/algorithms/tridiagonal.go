package algorithms

import "errors"

func solveTridiagonal(a, b, c, d []float64) ([]float64, error) {
	var n = len(b)
	if len(a) != n || len(c) != n || len(d) != n {
		return nil, errors.New("Wrong length")
	}
	var alph, bt = make([]float64, n-1), make([]float64, n-1)
	alph[0] = -c[0] / b[0]
	bt[0] = d[0] / b[0]

	for i := 1; i < n-1; i++ {
		alph[i] = -c[i] / (b[i] - alph[i-1]*-a[i])
		bt[i] = (d[i] + bt[i-1]*-a[i]) / (b[i] - alph[i-1]*-a[i])
	}
	var u = make([]float64, n)
	u[n-1] = (d[n-1] + bt[n-2]*-a[n-1]) / (b[n-1] - alph[n-2]*-a[n-1])
	for i := n - 2; i >= 0; i-- {
		u[i] = alph[i]*u[i+1] + bt[i]
	}
	return u, nil
}

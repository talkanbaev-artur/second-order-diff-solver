package functions

import (
	"math"

	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic/model"
)

func ProblemFactoryMethod(problemName string, eps float64) model.ODE {
	switch problemName {
	case "1":
		{
			return model.ODE{
				A:    func(f float64) float64 { return -1 },
				B:    func(f float64) float64 { return 0 },
				F:    func(f float64) float64 { return f * f * f },
				Xi1:  1,
				Xi2:  1,
				Eta1: 0,
				Eta2: 0,
				Phi1: func(f float64) float64 { return 0 },
				Phi2: func(f float64) float64 { return 1 },
			}
		}
	case "2":
		{
			return model.ODE{
				A:    func(f float64) float64 { return -1 },
				B:    func(f float64) float64 { return 0 },
				F:    func(f float64) float64 { return f * f * f },
				Xi1:  1,
				Xi2:  1,
				Eta1: 0,
				Eta2: 1,
				Phi1: func(f float64) float64 { return 0 },
				Phi2: func(f float64) float64 { return 1 },
			}
		}
	case "3":
		{
			return model.ODE{
				A: func(f float64) float64 {
					return 3*math.Pow(1+f, 2) - (2*eps)/(1+f)
				},
				B: func(f float64) float64 { return 0 },
				F: func(f float64) float64 {
					return (1.5 * eps / (math.Pow(1+f, 2))) - (1.5 * (1 + f))
				},
				Xi1:  1,
				Xi2:  1,
				Eta1: 1.0 / 3.0,
				Eta2: 0,
				Phi1: func(f float64) float64 { return eps/6 - 1/(1-math.Pow(math.E, -7/eps)) },
				Phi2: func(f float64) float64 { return 1 - math.Ln2/2 },
			}
		}
	default:
		panic("method not found")
	}
}

var problemSolutions = map[string]model.EPSRF{
	"3": func(eps float64) model.RF {
		return func(x float64) float64 {
			up := 1 - math.Pow(math.E, (1-math.Pow(1+x, 3))/eps)
			down := 1 - math.Pow(math.E, -7/eps)
			right := math.Log(1+x) / 2
			return up/down - right
		}
	},
	"1": func(eps float64) model.RF {
		return func(x float64) float64 {
			eVal := math.Pow(math.E, 1/eps)
			eps3 := 24 * eps * eps * eps
			x2 := x * x
			x3 := x * x2
			eps2 := 12 * eps * eps
			c1 := x2*x2 + 4*x3*eps + eps2*x2 - x*eVal*(x3+4*x2*eps+x*eps2+eps3)
			c3 := eps3 + eps2 + 4*eps + 5
			return (1 / (4 * (eVal - 1))) * (c1 + x*eps3 + c3*(math.Pow(math.E, x/eps)-1))
		}
	},
	"2": func(eps float64) model.RF {
		return func(x float64) float64 {
			eVal := math.Pow(math.E, 1/eps)
			eps1 := 4 * eps
			eps2 := 12 * eps * eps
			eps3 := 24 * eps * eps * eps
			x2 := x * x
			x3 := x * x2
			c1 := x2*x2 + eps1*x3 + eps2*x2 - 2*x*eVal*(x3+x2*eps1+x*eps2+eps3)
			c3 := eps3*eps + 2*eps3 + 2*eps2 + 2*eps1 + 5
			return (1 / (8*eVal - 4)) * (c1 + x*eps3 + c3*(math.Pow(math.E, x/eps)-1))
		}
	},
}

func Evaluate(problemName string, eps float64) model.SolutionData {
	n := 50000 //grid size
	factory, ok := problemSolutions[problemName]
	if !ok {
		panic("method not found")
	}
	f := factory(eps)
	var sol model.SolutionData
	h := 1.0 / float64(n)
	for i := 0; i < n; i++ {
		sol.XValues = append(sol.XValues, h*float64(i))
		sol.YValues = append(sol.YValues, f(sol.XValues[i]))
	}
	return sol
}

func EvaluateOnGrid(problemName string, eps float64, grid []float64) []float64 {
	factory, ok := problemSolutions[problemName]
	if !ok {
		panic("method not found")
	}
	f := factory(eps)
	var sol []float64
	for i := 0; i < len(grid); i++ {
		sol = append(sol, f(grid[i]))
	}
	return sol
}

func GetUniformGrid(n int) []float64 {
	h := 1.0 / float64(n)
	var x []float64
	for i := 0; i <= n; i++ {
		x = append(x, h*float64(i))
	}
	return x
}

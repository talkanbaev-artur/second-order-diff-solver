package functions

import (
	"math"

	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic/model"
)

func ProblemFactoryMethod(problemName string, eps float64) model.ODE {
	switch problemName {
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

func GetUniformGrid(n int) []float64 {
	h := 1.0 / float64(n)
	var x []float64
	for i := 0; i <= n; i++ {
		x = append(x, h*float64(i))
	}
	return x
}

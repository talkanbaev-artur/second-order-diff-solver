package functions

import (
	"math"

	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic/model"
)

func ProblemFactoryMethod(problemName string) model.ODE {
	return model.ODE{}
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

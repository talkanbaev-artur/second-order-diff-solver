package algorithms

import (
	"math"

	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic/functions"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic/model"
)

type solver struct {
	Original   model.ODE
	Inputs     model.Boundary2ndODEInputs
	a, b, c, f []float64
}

type viscosityParam model.RF

func getViscosityParam(scheme string) viscosityParam {
	switch scheme {
	case "central":
		return func(f float64) float64 {
			return 1
		}
	case "directional":
		return func(f float64) float64 {
			return 1 + math.Abs(f)
		}
	default:
		panic("unknown function")
	}
}

func Solve2nOrderODE(original model.ODE, inputs model.Boundary2ndODEInputs) (model.NumericalSolution, error) {
	s := solver{Original: original, Inputs: inputs}
	err := s.Precalc()
	if err != nil {
		return model.NumericalSolution{}, err
	}
	data := s.Solve()
	return data, nil
}

func (s *solver) Precalc() error {
	visc := getViscosityParam(s.Inputs.DifferenceScheme)
	mesh := functions.GetUniformGrid(s.Inputs.GridSize)
	n, h, eps := s.Inputs.GridSize+1, 1.0/float64(s.Inputs.GridSize), s.Inputs.EpsilonParam
	var a, b, c, f = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
	a[0], c[n-1] = 0, 0
	b[0], c[0], f[0] = s.Original.Xi1+s.Original.Eta1*eps/h, s.Original.Eta1*eps/h*-1, s.Original.Phi1(eps)
	a[n-1], b[n-1], f[n-1] = s.Original.Eta2*eps/h*-1, s.Original.Xi2+s.Original.Eta2*eps/h, s.Original.Phi2(eps)
	for i := 1; i < n-1; i++ {
		f[i] = s.Original.F(mesh[i])
		a_o, b_o := s.Original.A(mesh[i]), s.Original.B(mesh[i])
		rVal := a_o * h / (2 * eps)
		eVal, hSq := 2*eps*visc(rVal), math.Pow(h, 2)
		a[i] = (a_o*h + eVal) / 2 * hSq
		b[i] = -1 * (eVal/hSq + b_o)
		c[i] = (eVal + a_o*h) / (-2 * hSq)
	}
	s.a, s.b, s.c, s.f = a, b, c, f
	return nil
}

func (s *solver) Solve() model.NumericalSolution {
	mesh := functions.GetUniformGrid(s.Inputs.GridSize)
	vals, err := solveTridiagonal(s.a, s.b, s.c, s.f)
	if err != nil {
		panic(err)
	}
	return model.NumericalSolution{SolutionData: model.SolutionData{XValues: mesh, YValues: vals}}
}

package logic

import (
	"context"
	"errors"

	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic/algorithms"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic/functions"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic/model"
)

type APIService interface {
	GetNumericals(ctx context.Context) []model.Numericals
	GetAvailableTasks(ctx context.Context) ([]model.Task, error)
	GetRealSolution(ctx context.Context, taskName string, eps float64) (model.SolutionData, error)
	Calculate2ODE(ctx context.Context, inp model.Boundary2ndODEInputs) (model.NumericalSolution, error)
}

func NewAPIService() APIService {
	return service{}
}

type service struct {
}

func (s service) GetNumericals(ctx context.Context) []model.Numericals {
	nums := []model.Numericals{
		{MethodID: 1, MethodTitle: "Numerical methods for 2nd-order ODE Boundary problems"},
	}

	return nums
}

func (service) GetAvailableTasks(ctx context.Context) ([]model.Task, error) {
	return nil, errors.New("not implemented")
}

func (service) GetRealSolution(ctx context.Context, taskName string, eps float64) (model.SolutionData, error) {
	sol := functions.Evaluate(taskName, eps)
	return sol, nil
}

func (service) Calculate2ODE(ctx context.Context, inp model.Boundary2ndODEInputs) (model.NumericalSolution, error) {
	ode := functions.ProblemFactoryMethod(inp.TaskName, inp.EpsilonParam)
	solution, err := algorithms.Solve2nOrderODE(ode, inp)
	return solution, err
}

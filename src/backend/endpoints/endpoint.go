package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic/model"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/util"
)

type Endpoints struct {
	GetAnalyticalSolution endpoint.Endpoint
	GetNumericalSolution  endpoint.Endpoint
	GetNumericals         endpoint.Endpoint
}

func CreateEndpoints(s logic.APIService, lg log.Logger) Endpoints {
	es := Endpoints{}
	es.GetNumericals = MakeGetNumericalsEndpoint(s, lg)
	es.GetAnalyticalSolution = MakeGetAnalyticalSolutionEndpoint(s, lg)
	es.GetNumericalSolution = MakeGetNumericalSolutionEndpoint(s, lg)
	return es
}

func MakeGetNumericalsEndpoint(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data := s.GetNumericals(ctx)
		return data, nil
	}
	e = util.LoggingMiddleware(lg)(e)
	return e
}

type GetAnalyticalSolutionRequest struct {
	Eps      float64 `json:"eps"`
	TaskName string  `json:"task"`
}

func MakeGetAnalyticalSolutionEndpoint(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetAnalyticalSolutionRequest)
		data, err := s.GetRealSolution(ctx, req.TaskName, req.Eps)
		return data, err
	}
	e = util.LoggingMiddleware(lg)(e)
	return e
}

type GetNumericalSolutionRequest struct {
	GridSize     int     `json:"n"`
	EpsilonValue float64 `json:"eps"`
	Scheme       string  `json:"scheme"`
	Task         string  `json:"task"`
}

func MakeGetNumericalSolutionEndpoint(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetNumericalSolutionRequest)
		data, err := s.Calculate2ODE(ctx, model.Boundary2ndODEInputs{GridSize: req.GridSize,
			EpsilonParam: req.EpsilonValue, DifferenceScheme: req.Scheme, TaskName: req.Task})
		return data, err
	}
	e = util.LoggingMiddleware(lg)(e)
	return e
}

package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/util"
)

type Endpoints struct {
	GetNumericals endpoint.Endpoint
}

func CreateEndpoints(s logic.APIService, lg log.Logger) Endpoints {
	es := Endpoints{}
	es.GetNumericals = MakeGetNumericalsEndpoint(s, lg)
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

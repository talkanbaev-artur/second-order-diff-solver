package logic

import (
	"context"

	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic/model"
)

type APIService interface {
	GetNumericals(ctx context.Context) []model.Numericals
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

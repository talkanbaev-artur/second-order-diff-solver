package server

import (
	"context"
	"encoding/json"
	"net/http"

	kitlog "github.com/go-kit/kit/log/logrus"
	"github.com/go-kit/kit/transport"
	transp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/endpoints"
	"github.com/talkanbaev-artur/auca-numericals-template/src/backend/logic"
)

type errMsg struct {
	Err string `json:"error"`
}

func MakeMuxRoutes(s logic.APIService, r *mux.Router, lg *logrus.Logger) {
	log := kitlog.NewLogger(lg)
	options := []transp.ServerOption{
		transp.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(errMsg{err.Error()})
		}),
		transp.ServerErrorHandler(transport.NewLogErrorHandler(log)),
	}

	ends := endpoints.CreateEndpoints(s, log)

	r.Methods("GET").Path("/numericals").Handler(
		transp.NewServer(
			ends.GetNumericals,
			decodeGetNumericalsRequest,
			encodeResponse,
			options...,
		),
	)
}

func decodeGetNumericalsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	r.Methods("POST").Path("/analytical").Handler(
		transp.NewServer(
			ends.GetAnalyticalSolution,
			decodeGetAnalyticalSolutionReqeust,
			encodeResponse,
			options...,
		),
	)

	r.Methods("POST").Path("/solve").Handler(
		transp.NewServer(
			ends.GetNumericalSolution,
			decodeGetNumericalSolutionReqeust,
			encodeResponse,
			options...,
		),
	)
}

func decodeGetNumericalsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeGetAnalyticalSolutionReqeust(_ context.Context, r *http.Request) (any, error) {
	var req endpoints.GetAnalyticalSolutionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func decodeGetNumericalSolutionReqeust(_ context.Context, r *http.Request) (any, error) {
	var req endpoints.GetNumericalSolutionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func AttachSPA(r *mux.Router, base string, index string) {
	h := spaHandler{base: base, index: index}
	r.PathPrefix("/").Handler(h)
}

type spaHandler struct {
	base  string
	index string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	path = filepath.Join(h.base, path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.base, h.index))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.base)).ServeHTTP(w, r)
}

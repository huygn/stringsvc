package stringsvc

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
func MakeHTTPHandler(ctx context.Context, s Service) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/uppercase").Handler(httptransport.NewServer(
		ctx,
		MakeUppercaseEndpoint(s),
		decodeUppercaseRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/count").Handler(httptransport.NewServer(
		ctx,
		MakeCountEndpoint(s),
		decodeCountRequest,
		encodeResponse,
	))
	return r
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

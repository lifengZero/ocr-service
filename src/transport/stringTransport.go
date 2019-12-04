package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"ocr-service/src/dto"
)

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request dto.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EcodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

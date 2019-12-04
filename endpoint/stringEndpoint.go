package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"ocr-service/dto"
	"ocr-service/service"
)

func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return dto.UppercaseResponse{v, err.Error()}, nil
		}
		return dto.UppercaseResponse{v, ""}, nil
	}
}

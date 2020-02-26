package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"ocr-service/config"
	"ocr-service/dto"
	"ocr-service/service"
``
)

func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			config.Logger.Error(err)
			return dto.UppercaseResponse{v, err.Error()}, nil
		}
		config.Logger.Info("返回参数:", v)
		return dto.UppercaseResponse{v, ""}, nil
	}
}

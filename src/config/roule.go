package config

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"ocr-service/src/endpoint"
	"ocr-service/src/service/impl"
	"ocr-service/src/transport"
)

func MakeHttpHandler() http.Handler {
	mux := http.NewServeMux()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}
	svc := impl.StringService{}

	mux.Handle("/uppercase", kithttp.NewServer(
		endpoint.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EcodeResponse,
		options...,
	))


	return mux
}


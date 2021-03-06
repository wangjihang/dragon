package endpoint

import (
	"context"

	"dragon/pb"
	"dragon/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Set struct {
	SumEndpoint    endpoint.Endpoint
	ConcatEndpoint endpoint.Endpoint
	HelloEndpoint  endpoint.Endpoint
}

func New(svc service.AddService, logger log.Logger) Set {
	var sumEndpoint endpoint.Endpoint
	{
		sumEndpoint = MakeSumEndpoint(svc)
		sumEndpoint = LoggingMiddleware(log.With(logger, "method", "Sum"))(sumEndpoint)
	}

	var concatEndpoint endpoint.Endpoint
	{
		concatEndpoint = MakeConcatEndpoint(svc)
		concatEndpoint = LoggingMiddleware(log.With(logger, "method", "Concat"))(concatEndpoint)
	}

	var helloEndpoint endpoint.Endpoint
	{
		helloEndpoint = MakeHelloEndpoint(svc)
		helloEndpoint = LoggingMiddleware(log.With(logger, "method", "Hello"))(helloEndpoint)
	}

	return Set{
		SumEndpoint:    sumEndpoint,
		ConcatEndpoint: concatEndpoint,
		HelloEndpoint:  helloEndpoint,
	}
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func MakeSumEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.SumReq)
		return svc.Sum(ctx, req)

	}
}

func MakeConcatEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ConcatReq)
		return svc.Concat(ctx, req)
	}
}

func MakeHelloEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.HelloReq)
		return svc.Hello(ctx, req)
	}
}

package metrics

import (
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	pb "github.com/alvarowolfx/smart-street-sensor/api"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc"
)

type service struct {
	metricsRepository Repository
}

// ensure we implement the gRPC service
var _ pb.MetricServiceServer = &service{}

// NewService create a new metrics service
func NewService(metricsRepository Repository) (kit.Service, error) {
	return &service{
		metricsRepository: metricsRepository,
	}, nil
}

func (s *service) HTTPOptions() []httptransport.ServerOption {
	return nil
}

func (s *service) HTTPRouterOptions() []kit.RouterOption {
	return []kit.RouterOption{
		kit.RouterSelect("stdlib"),
	}
}

func (s *service) HTTPMiddleware(h http.Handler) http.Handler {
	return h
}

func (s *service) Middleware(ep endpoint.Endpoint) endpoint.Endpoint {
	return ep
}

// declare the endpoints for the HTTP server
func (s *service) HTTPEndpoints() map[string]map[string]kit.HTTPEndpoint {
	return map[string]map[string]kit.HTTPEndpoint{
		"/metrics": {
			"POST": {
				Endpoint: s.sendMetric,
				Decoder:  s.decodeSendMetricRequest,
			},
		},
		"/batch-metrics": {
			"POST": {
				Endpoint: s.sendBatchMetric,
				Decoder:  s.decodeSendBatchMetricRequest,
			},
		},
	}
}

// tracing RPC requests
func (s *service) RPCMiddleware() grpc.UnaryServerInterceptor {
	/*if s.tracer != nil {
		return s.tracer.GRPCServerInterceptor()
	}*/
	return nil
}

func (s *service) RPCServiceDesc() *grpc.ServiceDesc {
	return &pb.MetricService_serviceDesc
}

func (s *service) RPCOptions() []grpc.ServerOption {
	return nil
}

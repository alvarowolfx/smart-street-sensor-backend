package metrics

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	pb "github.com/alvarowolfx/smart-street-sensor/api"
)

// gRpc
func (s service) SendMetric(stream pb.MetricService_SendMetricServer) error {
	//ctx := stream.Context()
	packets := 0
	tx, err := s.metricsRepository.BeginTransaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for {
		metric, err := stream.Recv()
		if err == io.EOF {
			if err := tx.Commit(); err != nil {
				return nil
			}

			return stream.SendAndClose(&pb.MetricResponse{
				Status:          pb.Status_SUCCESS,
				InsertedMetrics: int32(packets),
			})
		}

		if err != nil {
			return err
		}

		apiMetric := Metric{Metric: metric}
		err = s.metricsRepository.SaveMetricWithTx(tx, apiMetric)
		//_, err = s.sendMetric(ctx, metric)

		if err != nil {
			return err
		}

		packets++
	}
}

// Shared biz for sendMetric
func (s service) sendMetric(ctx context.Context, req interface{}) (interface{}, error) {
	metric := req.(*pb.Metric)

	apiMetric := Metric{Metric: metric}

	err := s.metricsRepository.SaveMetric(apiMetric)

	return apiMetric, err
}

func (s service) decodeSendMetricRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return nil, kit.NewJSONStatusResponse(
			&pb.StatusResponse{Status: pb.Status_FAILED, Message: err.Error()},
			http.StatusBadRequest)
	}

	if err := r.Body.Close(); err != nil {
		return nil, kit.NewJSONStatusResponse(
			&pb.StatusResponse{Status: pb.Status_FAILED, Message: err.Error()},
			http.StatusBadRequest)
	}

	var metric pb.Metric
	err = json.Unmarshal(body, &metric)
	if err != nil {
		return nil, kit.NewJSONStatusResponse(
			&pb.StatusResponse{Status: pb.Status_FAILED, Message: err.Error()},
			http.StatusBadRequest)
	}

	return &metric, err
}

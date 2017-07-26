package metrics

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	pb "github.com/alvarowolfx/smart-street-sensor/api"
)

// gRpc

// Shared biz for sendMetric
func (s service) sendBatchMetric(ctx context.Context, req interface{}) (interface{}, error) {
	metrics := req.([]pb.Metric)
	packets := 0

	tx, err := s.metricsRepository.BeginTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	for _, metric := range metrics {
		apiMetric := Metric{Metric: &metric}
		err = s.metricsRepository.SaveMetricWithTx(tx, apiMetric)
		if err != nil {
			return nil, err
		}
		packets++
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.MetricResponse{
		InsertedMetrics: int32(packets),
		Status:          pb.Status_SUCCESS,
	}, nil
}

func (s service) decodeSendBatchMetricRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
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

	metrics := make([]pb.Metric, 0)
	err = json.Unmarshal(body, &metrics)
	if err != nil {
		return nil, kit.NewJSONStatusResponse(
			&pb.StatusResponse{Status: pb.Status_FAILED, Message: err.Error()},
			http.StatusBadRequest)
	}

	return metrics, err
}

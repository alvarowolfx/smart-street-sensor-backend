package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	pb "github.com/alvarowolfx/smart-street-sensor/api"
	"github.com/alvarowolfx/smart-street-sensor/backend/util"
)

func sendMetric(metric *pb.Metric) error {
	uri := fmt.Sprintf("%s:%d/metrics", *serverAddr, *serverPort)
	jsonStr, err := json.Marshal(metric)
	if err != nil {
		return err
	}
	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)

	return nil
}

func sendAllMetricsIndividually(metrics []*pb.Metric) error {
	for _, metric := range metrics {
		err := sendMetric(metric)
		if err != nil {
			return err
		}
	}

	return nil
}

func sendAllMetricsBatch(metrics []*pb.Metric) error {
	uri := fmt.Sprintf("%s:%d/batch-metrics", *serverAddr, *serverPort)
	jsonStr, err := json.Marshal(metrics)
	if err != nil {
		return err
	}
	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)

	return nil
}

func jsonMetricSize() int {
	metric := util.GenerateMetric()
	jsonStr, _ := json.Marshal(metric)
	return len(jsonStr)
}

func BenchmarkBatchHTTPServer10(b *testing.B)     { benchmarkBatchHTTPServer(10, b) }
func BenchmarkBatchHTTPServer100(b *testing.B)    { benchmarkBatchHTTPServer(100, b) }
func BenchmarkBatchHTTPServer1000(b *testing.B)   { benchmarkBatchHTTPServer(1000, b) }
func BenchmarkBatchHTTPServer10000(b *testing.B)  { benchmarkBatchHTTPServer(10000, b) }
func BenchmarkBatchHTTPServer100000(b *testing.B) { benchmarkBatchHTTPServer(100000, b) }

func benchmarkBatchHTTPServer(packetsCount int, b *testing.B) {
	len := int64(jsonMetricSize() * packetsCount)
	b.SetBytes(len)
	//fmt.Printf("%d bytes in json\n", len)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		runBatchHTTPTest(packetsCount, b)
	}
}

func runBatchHTTPTest(packetsCount int, b *testing.B) {
	b.StopTimer()
	metrics := util.GenerateMetrics(packetsCount)
	b.StartTimer()

	err := sendAllMetricsBatch(metrics)
	if err != nil {
		b.Fatalf("SendTelemetry(_) = _, %v", err)
	}
}

func BenchmarkHTTPServer10(b *testing.B)     { benchmarkHTTPServer(10, b) }
func BenchmarkHTTPServer100(b *testing.B)    { benchmarkHTTPServer(100, b) }
func BenchmarkHTTPServer1000(b *testing.B)   { benchmarkHTTPServer(1000, b) }
func BenchmarkHTTPServer10000(b *testing.B)  { benchmarkHTTPServer(10000, b) }
func BenchmarkHTTPServer100000(b *testing.B) { benchmarkHTTPServer(100000, b) }

func benchmarkHTTPServer(packetsCount int, b *testing.B) {
	len := int64(jsonMetricSize() * packetsCount)
	b.SetBytes(len)
	//fmt.Printf("%d bytes in json\n", len)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		runHTTPTest(packetsCount, b)
	}
}

func runHTTPTest(packetsCount int, b *testing.B) {
	b.StopTimer()
	metrics := util.GenerateMetrics(packetsCount)
	b.StartTimer()

	err := sendAllMetricsIndividually(metrics)
	if err != nil {
		b.Fatalf("SendTelemetry(_) = _, %v", err)
	}
}

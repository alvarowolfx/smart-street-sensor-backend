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

func sendPacket(packet *pb.SyncPacket) error {
	sendPacketURI := fmt.Sprintf("%s/send-telemetry", *serverAddr)
	jsonStr, err := json.Marshal(packet)
	if err != nil {
		return err
	}
	resp, err := http.Post(sendPacketURI, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)

	return nil
}

func sendAllPacketsIndividually(packets []*pb.SyncPacket) error {
	for _, packet := range packets {
		err := sendPacket(packet)
		if err != nil {
			return err
		}
	}

	return nil
}

func sendAllPacketsBatch(packets []*pb.SyncPacket) error {
	uri := fmt.Sprintf("%s/send-batch-telemetry", *serverAddr)
	jsonStr, err := json.Marshal(packets)
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

func jsonPacketSize() int {
	syncPacket := util.GeneratePacket()
	jsonStr, _ := json.Marshal(syncPacket)
	return len(jsonStr)
}

func BenchmarkBatchHttpServer10(b *testing.B)     { benchmarkBatchHTTPServer(10, b) }
func BenchmarkBatchHttpServer100(b *testing.B)    { benchmarkBatchHTTPServer(100, b) }
func BenchmarkBatchHttpServer1000(b *testing.B)   { benchmarkBatchHTTPServer(1000, b) }
func BenchmarkBatchHttpServer10000(b *testing.B)  { benchmarkBatchHTTPServer(10000, b) }
func BenchmarkBatchHttpServer100000(b *testing.B) { benchmarkBatchHTTPServer(100000, b) }

func benchmarkBatchHTTPServer(packetsCount int, b *testing.B) {
	len := int64(jsonPacketSize() * packetsCount)
	b.SetBytes(len)
	//fmt.Printf("%d bytes in json\n", len)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		runBatchHTTPTest(packetsCount, b)
	}
}

func runBatchHTTPTest(packetsCount int, b *testing.B) {
	b.StopTimer()
	packets := util.GeneratePackets(packetsCount)
	b.StartTimer()

	err := sendAllPacketsBatch(packets)
	if err != nil {
		b.Fatalf("SendTelemetry(_) = _, %v", err)
	}
}

func BenchmarkHttpServer10(b *testing.B)     { benchmarkHTTPServer(10, b) }
func BenchmarkHttpServer100(b *testing.B)    { benchmarkHTTPServer(100, b) }
func BenchmarkHttpServer1000(b *testing.B)   { benchmarkHTTPServer(1000, b) }
func BenchmarkHttpServer10000(b *testing.B)  { benchmarkHTTPServer(10000, b) }
func BenchmarkHttpServer100000(b *testing.B) { benchmarkHTTPServer(100000, b) }

func benchmarkHTTPServer(packetsCount int, b *testing.B) {
	len := int64(jsonPacketSize() * packetsCount)
	b.SetBytes(len)
	//fmt.Printf("%d bytes in json\n", len)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		runHTTPTest(packetsCount, b)
	}
}

func runHTTPTest(packetsCount int, b *testing.B) {
	b.StopTimer()
	packets := util.GeneratePackets(packetsCount)
	b.StartTimer()

	err := sendAllPacketsIndividually(packets)
	if err != nil {
		b.Fatalf("SendTelemetry(_) = _, %v", err)
	}
}

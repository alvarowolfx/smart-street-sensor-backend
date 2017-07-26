package util

import (
	pb "github.com/alvarowolfx/smart-street-sensor/api"
	"github.com/wawandco/fako"
)

/*
func RandomPacket(r *rand.Rand) *pb.SyncPacket {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &pb.SyncPacket{
		Location: &pb.Location{Latitude: lat, Longitude: long},
	}
}

func GeneratePackets(n int) []*pb.SyncPacket {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var packets []*pb.SyncPacket
	for i := 0; i < n; i++ {
		packets = append(packets, RandomPacket(r))
	}

	return packets
}
*/

func GenerateMetric() *pb.Metric {
	metric := &pb.Metric{}
	fako.Fuzz(metric)
	return metric
}

func GenerateMetrics(n int) []*pb.Metric {
	var metrics []*pb.Metric
	for i := 0; i < n; i++ {
		metrics = append(metrics, GenerateMetric())
	}
	return metrics
}

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

func GeneratePacket() *pb.SyncPacket {
	packet := &pb.SyncPacket{
		Location:    &pb.Location{},
		AccData:     &pb.AccelerometerData{},
		GyroData:    &pb.GyroscopeData{},
		MagData:     &pb.MagnetometerData{},
		ClimateData: &pb.ClimateData{},
	}
	fako.Fuzz(packet.Location)
	fako.Fuzz(packet.GyroData)
	fako.Fuzz(packet.AccData)
	fako.Fuzz(packet.MagData)
	fako.Fuzz(packet.ClimateData)
	return packet
}

func GeneratePackets(n int) []*pb.SyncPacket {
	var packets []*pb.SyncPacket
	for i := 0; i < n; i++ {
		packets = append(packets, GeneratePacket())
	}
	return packets
}

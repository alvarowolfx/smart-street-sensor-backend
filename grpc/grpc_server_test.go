package grpc

import (
	"context"
	"flag"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "github.com/alvarowolfx/smart-street-sensor-backend/syncservice"
	"github.com/alvarowolfx/smart-street-sensor-backend/util"
	"github.com/golang/protobuf/proto"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "testdata/ca.pem", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func connectToServer() *grpc.ClientConn {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		var sn string
		if *serverHostOverride != "" {
			sn = *serverHostOverride
		}
		var creds credentials.TransportCredentials
		if *caFile != "" {
			var err error
			creds, err = credentials.NewClientTLSFromFile(*caFile, sn)
			if err != nil {
				grpclog.Fatalf("Failed to create TLS credentials %v", err)
			}
		} else {
			creds = credentials.NewClientTLSFromCert(nil, sn)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	return conn
}

func sendAllPackets(client pb.SyncServiceClient, packets []*pb.SyncPacket) (*pb.SyncResponse, error) {
	stream, err := client.SendTelemetry(context.Background())
	if err != nil {
		return nil, err
	}

	for _, packet := range packets {
		if err := stream.Send(packet); err != nil {
			return nil, err
		}
	}

	syncResponse, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return syncResponse, nil
}

func BenchmarkGrpcServer10(b *testing.B)     { benchmarkGrpcServer(10, b) }
func BenchmarkGrpcServer100(b *testing.B)    { benchmarkGrpcServer(100, b) }
func BenchmarkGrpcServer1000(b *testing.B)   { benchmarkGrpcServer(1000, b) }
func BenchmarkGrpcServer10000(b *testing.B)  { benchmarkGrpcServer(10000, b) }
func BenchmarkGrpcServer100000(b *testing.B) { benchmarkGrpcServer(100000, b) }

func protobufPacketSize() int {
	syncPacket := util.GeneratePacket()
	protoStr, _ := proto.Marshal(syncPacket)
	return len(protoStr)
}

func benchmarkGrpcServer(packetsCount int, b *testing.B) {
	len := int64(protobufPacketSize() * packetsCount)
	b.SetBytes(len)
	//fmt.Printf("%d bytes in protof\n", len)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		runGrpcTest(packetsCount, b)
	}
}

func runGrpcTest(packetsCount int, b *testing.B) {
	b.StopTimer()
	conn := connectToServer()
	defer conn.Close()
	client := pb.NewSyncServiceClient(conn)
	packets := util.GeneratePackets(packetsCount)
	b.StartTimer()
	_, err := sendAllPackets(client, packets)
	if err != nil {
		b.Fatalf("%v.SendTelemetry(_) = _, %v", client, err)
	}
}

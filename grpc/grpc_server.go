package grpc

import (
	"io"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"

	pb "github.com/alvarowolfx/smart-street-sensor-backend/syncservice"
)

type SyncServiceGrpcServer struct {
	db *bolt.DB
}

func (s SyncServiceGrpcServer) SendTelemetry(stream pb.SyncService_SendTelemetryServer) error {
	packets := 0
	tx, err := s.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bkt, err := tx.CreateBucketIfNotExists([]byte("telemetry"))
	if err != nil {
		return err
	}

	for {
		syncPacket, err := stream.Recv()
		if err == io.EOF {
			if err := tx.Commit(); err != nil {
				return err
			}
			return stream.SendAndClose(&pb.SyncResponse{
				Status:          pb.Status_SUCCESS,
				InsertedPackets: int32(packets),
			})
		}
		if err != nil {
			return err
		}

		id, err := bkt.NextSequence()
		if err != nil {
			return err
		}

		syncPacket.Timestamp = int32(id)

		if buf, err := proto.Marshal(syncPacket); err != nil {
			return err
		} else if err := bkt.Put([]byte(strconv.FormatUint(id, 10)), buf); err != nil {
			return err
		}

		packets++
	}
}

func (s SyncServiceGrpcServer) GetLastJobSummary(ctx context.Context, requestSummary *pb.RequestSummary) (*pb.SyncResponse, error) {
	return nil, nil
}

func NewGrpcServer(db *bolt.DB) *SyncServiceGrpcServer {
	s := new(SyncServiceGrpcServer)
	s.db = db
	return s
}

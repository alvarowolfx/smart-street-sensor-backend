package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/boltdb/bolt"

	pb "github.com/alvarowolfx/smart-street-sensor-backend/syncservice"
)

type SyncServiceHttpServer struct {
	db *bolt.DB
}

func sendError(w http.ResponseWriter, err error) {
	data := struct {
		message string
	}{
		err.Error(),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(data)
}

func (s SyncServiceHttpServer) beginTransaction() (*bolt.Tx, *bolt.Bucket, error) {
	tx, err := s.db.Begin(true)

	if err != nil {
		return nil, nil, err
	}

	bkt, err := tx.CreateBucketIfNotExists([]byte("telemetry"))
	if err != nil {
		return nil, nil, err
	}

	return tx, bkt, nil
}

func (s SyncServiceHttpServer) savePacket(tx *bolt.Tx, bkt *bolt.Bucket, syncPacket *pb.SyncPacket) error {
	id, err := bkt.NextSequence()
	if err != nil {
		return err
	}

	syncPacket.Timestamp = int32(id)

	if buf, err := json.Marshal(syncPacket); err != nil {
		return err
	} else if err := bkt.Put([]byte(strconv.FormatUint(id, 10)), buf); err != nil {
		return err
	}

	return nil
}

func (s SyncServiceHttpServer) SendTelemetry(w http.ResponseWriter, r *http.Request) {
	tx, bkt, err := s.beginTransaction()
	if err != nil {
		sendError(w, err)
	}
	defer tx.Rollback()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		sendError(w, err)
	}

	if err := r.Body.Close(); err != nil {
		sendError(w, err)
	}

	var syncPacket pb.SyncPacket
	err = json.Unmarshal(body, &syncPacket)
	if err != nil {
		sendError(w, err)
	}

	err = s.savePacket(tx, bkt, &syncPacket)
	if err != nil {
		sendError(w, err)
	}

	if err := tx.Commit(); err != nil {
		sendError(w, err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pb.SyncResponse{
		Status:          pb.Status_SUCCESS,
		InsertedPackets: 1,
	})
}

func (s SyncServiceHttpServer) SendBatchTelemetry(w http.ResponseWriter, r *http.Request) {
	tx, bkt, err := s.beginTransaction()
	if err != nil {
		sendError(w, err)
		return
	}
	defer tx.Rollback()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendError(w, err)
		return
	}

	if err := r.Body.Close(); err != nil {
		sendError(w, err)
		return
	}

	packets := make([]pb.SyncPacket, 0)
	err = json.Unmarshal(body, &packets)
	if err != nil {
		sendError(w, err)
		return
	}

	for _, packet := range packets {
		err = s.savePacket(tx, bkt, &packet)
		if err != nil {
			sendError(w, err)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		sendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pb.SyncResponse{
		Status:          pb.Status_SUCCESS,
		InsertedPackets: int32(len(packets)),
	})
}

func NewHttpServer(db *bolt.DB) *SyncServiceHttpServer {
	s := new(SyncServiceHttpServer)
	s.db = db
	return s
}

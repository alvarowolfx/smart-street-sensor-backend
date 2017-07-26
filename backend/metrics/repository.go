package metrics

import (
	"strconv"

	pb "github.com/alvarowolfx/smart-street-sensor/api"
	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
)

type Transaction interface {
	Commit() error
	Rollback() error
}

//Repository repository with metrics
type Repository interface {
	SaveMetric(Metric) error
	SaveMetricWithTx(Transaction, Metric) error
	BeginTransaction() (Transaction, error)
	Commit(Transaction) error
}

type Metric struct {
	Metric *pb.Metric
}

var bucketName = []byte("metrics")

type metricsRepository struct {
	db *bolt.DB
}

type metricsTransaction struct {
	tx *bolt.Tx
}

func NewRepository(db *bolt.DB) (Repository, error) {
	return &metricsRepository{
		db: db,
	}, nil
}

func (mr *metricsRepository) BeginTransaction() (Transaction, error) {
	tx, err := mr.db.Begin(true)

	if err != nil {
		return nil, err
	}

	return metricsTransaction{tx: tx}, err
}

func (mr *metricsRepository) SaveMetric(metric Metric) error {
	tx, err := mr.BeginTransaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = mr.SaveMetricWithTx(tx, metric)

	if err := mr.Commit(tx); err != nil {
		return err
	}

	return nil
}

func (mr *metricsRepository) SaveMetricWithTx(tx Transaction, metric Metric) error {
	bkt, err := tx.(metricsTransaction).tx.CreateBucketIfNotExists(bucketName)
	if err != nil {
		return err
	}

	id, err := bkt.NextSequence()
	if err != nil {
		return err
	}

	metric.Metric.Timestamp = int32(id)

	if buf, err := proto.Marshal(metric.Metric); err != nil {
		return err
	} else if err := bkt.Put([]byte(strconv.FormatUint(id, 10)), buf); err != nil {
		return err
	}

	return nil
}

func (mr *metricsRepository) Commit(tx Transaction) error {
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (mt metricsTransaction) Commit() error {
	if err := mt.tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (mt metricsTransaction) Rollback() error {
	if err := mt.tx.Rollback(); err != nil {
		return err
	}
	return nil
}

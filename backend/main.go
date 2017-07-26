package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "github.com/alvarowolfx/smart-street-sensor/api"
	"github.com/alvarowolfx/smart-street-sensor/backend/metrics"
	"github.com/alvarowolfx/smart-street-sensor/backend/server"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "testdata/server1.pem", "The TLS cert file")
	keyFile    = flag.String("key_file", "testdata/server1.key", "The TLS key file")
	port       = flag.Int("port", 10000, "The server port")
	serverType = flag.String("type", "grpc", "The server type: grpc or http")
)

func startGrpcServer(db *bolt.DB) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			logrus.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterMetricServiceServer(grpcServer, server.NewGrpcServer(db))

	logrus.Infof("GRPC Server initialized at port %d", *port)

	grpcServer.Serve(lis)

	grpclog.Infoln("End")
}

func startHTTPServer(db *bolt.DB) {

	httpServer := server.NewHTTPServer(db)

	http.HandleFunc("/send-telemetry", httpServer.SendMetric)
	http.HandleFunc("/send-batch-telemetry", httpServer.SendBatchTelemetry)

	httpPort := *port + 1
	portString := fmt.Sprintf(":%d", httpPort)

	logrus.Infof("HTTP Server initialized at port %d", httpPort)

	http.ListenAndServe(portString, nil)
}

func main() {
	flag.Parse()
	db, err := bolt.Open("my.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logrus.Infof("Starting servers")
	//if *serverType == "grpc" {
	//go startGrpcServer(db)
	//} else {
	//startHTTPServer(db)
	//}

	metricsRepository, err := metrics.NewRepository(db)
	if err != nil {
		log.Fatal(err)
	}
	svc, err := metrics.NewService(metricsRepository)
	if err != nil {
		log.Fatal(err)
	}

	kit.Run(svc)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/boltdb/bolt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	grpcSyncServer "github.com/alvarowolfx/smart-street-sensor-backend/grpc"
	httpSyncServer "github.com/alvarowolfx/smart-street-sensor-backend/http"
	pb "github.com/alvarowolfx/smart-street-sensor-backend/syncservice"
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
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterSyncServiceServer(grpcServer, grpcSyncServer.NewGrpcServer(db))

	grpclog.Infof("GRPC Server initialized at port %d", *port)

	grpcServer.Serve(lis)

	grpclog.Infoln("End")
}

func startHttpServer(db *bolt.DB) {

	httpServer := httpSyncServer.NewHttpServer(db)

	http.HandleFunc("/send-telemetry", httpServer.SendTelemetry)
	http.HandleFunc("/send-batch-telemetry", httpServer.SendBatchTelemetry)

	portString := fmt.Sprintf(":%d", *port)

	grpclog.Infof("HTTP Server initialized at port %d", *port)

	http.ListenAndServe(portString, nil)
}

func main() {
	flag.Parse()
	db, err := bolt.Open("my.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if *serverType == "grpc" {
		startGrpcServer(db)
	} else {
		startHttpServer(db)
	}
}

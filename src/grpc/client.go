package grpc

import (
	"log"

	"time"

	pb "./proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	grpcServerURL = "localhost:9999"
)

func RequestMonitoringJob(cateogry, targetUrl string) (bool, string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(grpcServerURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMonitoringJobClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendJobInfo(ctx, &pb.JobRequest{Category: cateogry, TargetUrl: targetUrl})
	if err != nil {
		log.Fatalf("could not connect to grpc servers: %v", err)
	}
	//log.Printf("Result: %v - %s", r.Status, r.Result)

	return r.Status, r.Result
}

package main

import (
	"fmt"
	"runtime"
	"time"

	grpc "../grpc"
)

func monitorHTTP() {
	t := time.Now()
	status, result := grpc.RequestMonitoringJob("http", "https://google.com/")
	fmt.Printf("%v  [%v] %s\n", t.UTC().Format("2006/1/2 15:04:05"), status, result)
}
func monitorDNS1() {
	t := time.Now()
	status, result := grpc.RequestMonitoringJob("dns", "https://google.com/")
	fmt.Printf("%v  [%v] %s\n", t.UTC().Format("2006/1/2 15:04:05"), status, result)
}
func monitorDNS2() {
	t := time.Now()
	status, result := grpc.RequestMonitoringJob("dns", "https://www.yahoo.co.jp/")
	fmt.Printf("%v  [%v] %s\n", t.UTC().Format("2006/1/2 15:04:05"), status, result)
}

func main() {
	// Run every 5 seconds but not now.
	//Every(2).Seconds().NotImmediately().Run(monitorHTTP)

	// Run every 5 seconds but not now.
	Every(2).Seconds().NotImmediately().Run(monitorDNS1)
	Every(9).Seconds().NotImmediately().Run(monitorDNS2)

	// Keep the program from not exiting.
	runtime.Goexit()
}

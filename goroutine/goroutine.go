package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Worker shutting down...")
			return
		default:
			fmt.Println("Worker processing...")
			time.Sleep(1 * time.Second)
		}
		
		
	}
}

func main() {
	stopChan := make(chan struct{})
	go worker(stopChan)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan

	fmt.Println("\nStopping worker...")

	close(stopChan)
	time.Sleep(1 * time.Second)
	fmt.Println("Worker gracefully stopped")
}
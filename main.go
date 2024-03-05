package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

type Watcher struct {
	watchList map[string]chan int
}

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"192.168.1.8:9092"},
		Topic:     "fullfillment.db0.user",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	// Create context to handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-sigCh:
			cancel()
		case <-ctx.Done():
		}
	}()

	// Start reading messages from Kafka topic
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}
		data, err := json.Marshal(string(msg.Value))
		if err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Println(data)
	}

	// Close Kafka reader
	reader.Close()
}

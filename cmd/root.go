package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/ripple-mq/go-client/api"
)

func Execute() {
	client := api.NewClient[string](api.Config{
		Topic:   "topic-X",
		Bucket:  "topic-Y",
		Brokers: []string{":8080"},
	})

	client.Create()

	prodCh := client.Writer()
	consCh := client.Reader()

	time.Sleep(4 * time.Second)
	fmt.Println("Producing")

	go func() {
		for i := 0; i < 1000100; i++ {
			prodCh <- strings.Repeat("x", 1024) // 1 KB message
		}
	}()

	time.Sleep(1 * time.Minute)
	fmt.Println("Producing Done")

	go func() {
		const messageCount = 10000
		const messageSizeKB = 1 // each message is 1 KB
		start := time.Now()

		for i := 0; i < messageCount; i++ {
			<-consCh
		}

		elapsed := time.Since(start).Seconds()
		totalMB := float64(messageCount*messageSizeKB) / 1024.0
		throughput := totalMB / elapsed

		fmt.Printf("Consumed %d messages of %d KB each in %.3f seconds\n", messageCount, messageSizeKB, elapsed)
		fmt.Printf("Consumer throughput: %.2f MB/s\n", throughput)
	}()

	select {}
}

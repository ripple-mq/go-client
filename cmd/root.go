package cmd

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ripple-mq/go-client/internal"
)

func Execute() {

	broker := internal.NewBroker(":8891")
	topic, bukcet := uuid.NewString(), uuid.NewString()
	broker.Create(topic, bukcet)

	prodCh := broker.RegisterProducer(topic, bukcet)
	consCh := broker.RegisterConsumer(topic, bukcet)
	time.Sleep(4 * time.Second)
	fmt.Println("Producing")
	i := 0
	go func() {
		for i := range 1000100 {
			prodCh <- fmt.Sprintf("MESSAGE: %d ", i)
		}
		fmt.Println(i)
	}()

	time.Sleep(1 * time.Minute)

	fmt.Println("Producing Done")
	// consCh := make(chan int, 100)
	// go func() {
	// 	for i := range 10000000 {
	// 		consCh <- i
	// 	}
	// }()

	go func() {
		start := time.Now()
		for range 10000 {
			d := <-consCh
			fmt.Println(d)
		}
		elapsed := time.Since(start)
		fmt.Printf("Time taken: %s\n", elapsed)
	}()

	select {}
}

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
	go func() {
		for i := range 1000000 {
			prodCh <- fmt.Sprintf("MESSAGE: %d ", i)
		}
	}()

	go func() {
		for range 1000000 {
			d := <-consCh
			fmt.Println(d)
		}
	}()

	select {}
}

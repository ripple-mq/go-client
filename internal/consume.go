package internal

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/charmbracelet/log"
	pb "github.com/ripple-mq/go-client/internal/proto"
	"github.com/ripple-mq/go-client/pkg/p2p/encoder"
	"github.com/ripple-mq/go-client/pkg/p2p/transport/tcp"
)

type AskQuery struct {
	Count int
	ID    string
}

func (t *Broker) RegisterConsumer(topic string, bucket string) <-chan any {
	resp, err := t.Client.GetConsumerConnection(context.Background(), &pb.GetConsumerConnnectionReq{Topic: topic, Bucket: bucket})
	if err != nil {
		fmt.Println("Consumer: failed to get response: ", err)
	}
	ch := make(chan any, 100)
	go consumeLoop(ch, ":8901", resp.ConsumerId)
	return ch
}

func consumeLoop(ch chan any, addr string, id string) {
	cons, _ := tcp.NewTransport(":8001", func(conn net.Conn, message []byte) {})
	cons.Listen()

	if err := cons.SendToAsync(addr, id, "0", AskQuery{Count: 400, ID: strconv.Itoa(0)}); err != nil {
		fmt.Printf("ConsumerLoop: failed to send read req: %v", err)
	}
	for {
		var msg []Payload
		_, err := cons.Consume(encoder.GOBDecoder{}, &msg)
		if err != nil {
			log.Warnf("error: %v", err)
		}
		for _, i := range msg {
			var m string
			err := encoder.GOBDecoder{}.Decode(bytes.NewBuffer(i.Data), &m)
			if err != nil {
				fmt.Printf("failed to decode: %v", err)
			}
			ch <- m
		}
	}
}

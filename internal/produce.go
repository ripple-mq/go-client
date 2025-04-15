package internal

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/ripple-mq/go-client/internal/proto"
	"github.com/ripple-mq/go-client/pkg/p2p/encoder"
	"github.com/ripple-mq/go-client/pkg/p2p/transport/tcp"
	"google.golang.org/grpc"
)

type Payload struct {
	Id   int32  // Unique identifier for the payload
	Data []byte // Actual data of the payload
}

type Broker struct {
	Client pb.BootstrapServerClient
}

func NewBroker(port string) *Broker {
	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error connecting to ripple serve, ", err)
		return nil
	}
	client := pb.NewBootstrapServerClient(conn)
	return &Broker{
		Client: client,
	}
}

func (t *Broker) RegisterProducer(topic string, bucket string) chan<- any {
	resp, err := t.Client.GetProducerConnection(context.Background(), &pb.GetProducerConnectionReq{Topic: topic, Bucket: bucket})
	if err != nil {
		fmt.Println("ProdcuerLoop: failed to get response: ", err)
	}
	ch := make(chan any, 100)
	go producerLoop(ch, ":8901", resp.ProducerId)
	return ch
}

func producerLoop(ch chan any, addr string, id string) {
	prod, _ := tcp.NewTransport(":8000", func(conn net.Conn, message []byte) {})
	prod.Listen()
	i := 0
	for data := range ch {
		var buff bytes.Buffer
		err := encoder.GOBEncoder{}.Encode(data, &buff)
		if err != nil {
			fmt.Printf("ProdcuerLoop: failed to encode: %v", err)
		}
		if err := prod.SendToAsync(addr, id, struct{}{}, Payload{Id: int32(i), Data: buff.Bytes()}); err != nil {
			fmt.Printf("ProdcuerLoop: failed to send: %v", err)
		}
		time.Sleep(1 * time.Millisecond)
	}
}

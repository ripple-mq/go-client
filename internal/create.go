package internal

import (
	"context"
	"fmt"

	pb "github.com/ripple-mq/go-client/internal/proto"
)

func (t *Broker[T]) Create(topic string, bucket string) {
	_, err := t.Client.CreateBucket(context.Background(), &pb.CreateBucketReq{Topic: topic, Bucket: bucket})
	if err != nil {
		fmt.Println("failed to get the response: ", err)
	}
}

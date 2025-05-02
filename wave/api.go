package wave

import "github.com/ripple-mq/go-client/internal"

type Config struct {
	Brokers       []string
	Bucket        string
	Topic         string
	ReadBatchSize int
}

type Client[T any] struct {
	Conf   Config
	broker internal.Broker[T]
}

// TODO try random connection attempts with all provided broker address
func NewClient[T any](conf Config) *Client[T] {
	return &Client[T]{Conf: conf, broker: *internal.NewBroker[T](conf.Brokers[0])}
}

func (t Client[T]) Create() {
	t.broker.Create(t.Conf.Topic, t.Conf.Bucket)
}

func (t Client[T]) Writer() chan<- any {
	prodCh := t.broker.RegisterProducer(t.Conf.Topic, t.Conf.Bucket)
	return prodCh
}

func (t Client[T]) Reader() <-chan any {
	consCh := t.broker.RegisterConsumer(t.Conf.Topic, t.Conf.Bucket, t.Conf.ReadBatchSize)
	return consCh
}

package kafka

import (
	"context"
	"time"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type KafkaDatabase struct {
	conn    *kafka.Conn
	url     string
	topic   string
	timeout time.Duration
}

func NewKafkaDatabase(URL, topic string, timeout int) (*KafkaRepository, error) {}

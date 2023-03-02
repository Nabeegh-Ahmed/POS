package kafkaHandler

import (
	"github.com/segmentio/kafka-go"
)

// KafkaReader returns a new kafka reader
func KafkaReader(kafkaURL, groupID, topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 1,    // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

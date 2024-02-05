package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce(topic string, partitionID int, message []byte) {
	// to produce messages
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", topic, partitionID)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: message},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func ListTopics() {
	conn, err := kafka.Dial("tcp", "localhost:29092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}

// func Consumer() {
// 	// to consume messages
// 	topic := "my-topic"
// 	partition := 0

// 	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", topic, partition)
// 	if err != nil {
// 		log.Fatal("failed to dial leader:", err)
// 	}

// 	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
// 	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

// 	b := make([]byte, 10e3) // 10KB max per message
// 	for {
// 		n, err := batch.Read(b)
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(string(b[:n]))
// 	}

// 	if err := batch.Close(); err != nil {
// 		log.Fatal("failed to close batch:", err)
// 	}

// 	if err := conn.Close(); err != nil {
// 		log.Fatal("failed to close connection:", err)
// 	}
// }

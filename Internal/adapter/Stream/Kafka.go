package Stream

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"secondchallange/Internal/Models"
	"secondchallange/config"
	"time"
)

func NewProduce(transaction *Models.Transaction, configurations config.Configurations) {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", configurations.Kafka.URL, configurations.Kafka.Topic, 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	obj, _ := json.Marshal(&transaction)
	conn.WriteMessages(kafka.Message{Value: []byte(obj)})
}
func NewConsumer(configurations config.Configurations) {
	config := kafka.ReaderConfig{
		Brokers:  []string{configurations.Kafka.URL},
		Topic:    configurations.Kafka.Topic,
		MaxBytes: 1e6,
	}
	reader := kafka.NewReader(config)
	for {
		message, error := reader.ReadMessage(context.Background())
		if error != nil {
			fmt.Println(time.Now().String()+":: Error happened during calling kafka server %v", error)
			continue
		}
		fmt.Println(time.Now().String() + "::message of transaction consumed:: " + string(message.Value))
	}
}

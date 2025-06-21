package kafkaconsumer

import (
	"context"
	"encoding/json"
	"github.com/Sanchir01/gw-notification/internal/domain/models"
	"github.com/Sanchir01/gw-notification/internal/feature/transaction"
	"github.com/segmentio/kafka-go"
	"log"
	"log/slog"
)

type KafkaConsumer struct {
	reader *kafka.Reader
	log    *slog.Logger
	trsrv  *transaction.Service
}

func NewConsumer(topic, broker, groupid string, trsrv *transaction.Service, log *slog.Logger) (*KafkaConsumer, error) {
	cfg := kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupid,
	}
	reader := kafka.NewReader(cfg)

	return &KafkaConsumer{
		reader: reader,
		log:    log,
		trsrv:  trsrv,
	}, nil
}

func (kc *KafkaConsumer) Run(ctx context.Context) error {
	defer kc.reader.Close()

	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 Kafka consumer shutdown")
			return nil

		default:
			msg, err := kc.reader.FetchMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return nil
				}
				kc.log.Error("❌ fetch error: %v", err)
				continue
			}

			var kafkaMsg models.KafkaMessage
			if err := json.Unmarshal(msg.Value, &kafkaMsg); err != nil {
				kc.log.Error("❌ failed to unmarshal Kafka message: %v", err)
				continue
			}

			var payload models.KafkaPayload
			if err := json.Unmarshal([]byte(kafkaMsg.Payload), &payload); err != nil {
				kc.log.Error("❌ failed to unmarshal Payload: %v", err)
				continue
			}
			kc.log.Info("payload: ", payload)
			if err := kc.trsrv.SetTransaction(ctx, payload.UserId, payload.Amount, kafkaMsg.Type); err != nil {
				kc.log.Error("failed to set transaction: %v", err)
				continue
			}
			if err := kc.reader.CommitMessages(ctx, msg); err != nil {
				kc.log.Error("⚠️ commit error: %v", err)
			}
		}
	}
}

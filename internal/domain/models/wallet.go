package models

import "github.com/google/uuid"

type KafkaMessage struct {
	ID         string `json:"ID"`
	Type       string `json:"Type"`
	ReservedTo string `json:"ReservedTo"`
	Payload    string `json:"Payload"`
}

type KafkaPayload struct {
	Amount       float32            `json:"amount"`
	UserId       uuid.UUID          `json:"user_id"`
	BalanceAfter map[string]float32 `json:"balance_after"`
}

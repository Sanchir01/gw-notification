package transaction

import "time"

type Transaction struct {
	ID        string    `bson:"id"`
	WalletId  string    `bson:"wallet_id"`
	Amount    float32   `bson:"amount"`
	Type      string    `bson:"type"`
	CreatedAt time.Time `bson:"created_at"`
}

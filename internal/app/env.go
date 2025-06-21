package app

import (
	"context"
	"github.com/Sanchir01/gw-notification/internal/config"
	kafkaconsumer "github.com/Sanchir01/gw-notification/internal/kafka"
	"github.com/Sanchir01/gw-notification/pkg/db"
	"github.com/Sanchir01/gw-notification/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type App struct {
	MongoCl       *mongo.Client
	Log           *slog.Logger
	Cfg           *config.Config
	KafkaConsumer *kafkaconsumer.KafkaConsumer
}

func NewApp(ctx context.Context) (*App, error) {
	cfg := config.InitConfig()
	l := logger.SetupLogger(cfg.Env)
	clientmonfog, err := db.NewMongoClient(ctx, cfg.DB.Mongo.Host, cfg.DB.Mongo.Port, cfg.DB.Mongo.Dbname)
	if err != nil {
		return nil, err
	}

	repos := NewRepositories(clientmonfog)
	srv := NewServices(repos)
	kafkacous, err := kafkaconsumer.NewConsumer(cfg.Kafka.Consumer.Topic[0], cfg.Kafka.Consumer.Brokers[0], cfg.Kafka.Consumer.GroupID, srv.TransactionService, l)
	if err != nil {
		return nil, err
	}
	return &App{
		MongoCl:       clientmonfog,
		Log:           l,
		Cfg:           cfg,
		KafkaConsumer: kafkacous,
	}, nil
}

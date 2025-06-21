package main

import (
	"context"
	"github.com/Sanchir01/gw-notification/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Signal(syscall.SIGINT), syscall.SIGTERM)
	defer cancel()
	env, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}
	env.Log.Info("starting app")
	go func() {
		if err := env.KafkaConsumer.Run(ctx); err != nil {
			env.Log.Error("kafka_consumer run error", "error", err.Error())
			return
		}
	}()
	<-ctx.Done()
	env.Log.Info("shutting down")

	if err := env.MongoCl.Disconnect(ctx); err != nil {
		env.Log.Error("mongo disconnect error", "error", err.Error())
	}
}

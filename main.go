package main

import (
	"github.com/corporateanon/barker-worker/pkg/config"
	"github.com/corporateanon/barker-worker/pkg/restclient"
	"github.com/corporateanon/barker-worker/pkg/telegramsender"
	"github.com/corporateanon/barker-worker/pkg/worker"
	"github.com/corporateanon/barker/pkg/client"
	"go.uber.org/fx"
)

func start(worker worker.Worker) {
	worker.Loop()
}

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			worker.NewWorkerImpl,
			telegramsender.NewSenderImplTelegram,
			client.NewBotDaoImplResty,
			client.NewDeliveryDaoImplResty,
			restclient.New,
		),
		fx.Invoke(start),
	)

	app.Run()
}

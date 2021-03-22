package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ITheCorgi/EscrumGolang/db"
	"github.com/ITheCorgi/EscrumGolang/internal/config"
)

func terminationHandler() {
	var terminate = make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-terminate
}

func Run(path string) error {
	terminationHandler()

	cfg, err := config.GetConfig(path)
	if err != nil {
		return err
	}

	mongoClient, err := db.NewDbInstance(cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		return err
	}

	//FIXME: introduce repo at a next task
	// ...

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		return err
	}

	return nil
}

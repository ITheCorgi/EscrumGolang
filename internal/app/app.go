package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ITheCorgi/EscrumGolang/db"
	"github.com/ITheCorgi/EscrumGolang/internal/config"
)

//terminateHandler checks for key combo to terminate server connection
func setTerminateHandler() {
	var keyboardInterupt = make(chan os.Signal, 1)
	signal.Notify(keyboardInterupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-keyboardInterupt
		os.Exit(1)
	}()
}

func Run(configFilePath string) {
	setTerminateHandler()

	cfg, err := config.GetConfig(configFilePath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	mongoClient, err := db.NewDbInstance(cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		log.Println("Error is occured during connecting to Mongo database")
		return
	}

	db := mongoClient.Database(cfg.Mongo.Name)

	//for debugging purpose
	log.Printf("Mongo URI: %s, User: %s, Password: %s", cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	log.Println(mongoClient)
	log.Println(db)
}

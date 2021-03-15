package server

import (
	"log"

	"github.com/ITheCorgi/EscrumGolang/app/config"
	"github.com/ITheCorgi/EscrumGolang/db"
)

func Run(configFilePath string) {
	cfg, err := config.GetConfig(configFilePath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	mongoClient, err := db.NewDbInstance(cfg.MongoDb.URI, cfg.MongoDb.User, cfg.MongoDb.Password)
	if err != nil {
		log.Println("Error is occured during connecting to Mongo database")
		return
	}

	db := mongoClient.Database(cfg.MongoDb.Name)
	//for debugging purpose
	log.Printf("Mongo URI: %s, User: %s, Password: %s", cfg.MongoDb.URI, cfg.MongoDb.User, cfg.MongoDb.Password)
	log.Println(mongoClient)
	log.Println(db)
}

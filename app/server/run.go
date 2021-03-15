package server

import (
	"log"

	"github.com/ITheCorgi/EscrumGolang/app/config"
	"github.com/ITheCorgi/EscrumGolang/db"
)

func Run(configFilePath string) {
	cfg, err := config.GetConfig(configFilePath)
	if err != nil {
		log.Println("Error is occured during reading configuration file")
		return
	}

	mongoClient, err := db.NewDbInstance(cfg.MongoDbAuth.URI, cfg.MongoDbAuth.User, cfg.MongoDbAuth.Password)
	if err != nil {
		log.Println("Error is occured during connecting to Mongo database")
		return
	}

	db := mongoClient.Database(cfg.MongoDbAuth.Name)

}

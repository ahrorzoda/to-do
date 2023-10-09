package main

import (
	"github.com/ahrorzoda/to-do/internal/composites"
	"github.com/ahrorzoda/to-do/pkg/config"
	logger "github.com/ahrorzoda/to-do/pkg/logger"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// entry point and implementation for all composites
	logger.Info.Println("config initializing")
	setConfig := config.InitJsonConfigs().Postgres
	postgresConfig, err := composites.NewConfigComposite(setConfig)
	if err != nil {
		logger.Error.Println("config composite failed")
		return
	}

	logger.Info.Printf("postgres composite initializing")
	postgresDBC, err := composites.NewPostgresDBComposite(postgresConfig)
	if err != nil {
		logger.Error.Println("postgres composite failed")
		return
	}

	logger.Info.Printf("List composite initializing")
	ListComposite, err := composites.NewListComposite(postgresDBC)
	if err != nil {
		logger.Error.Println("List composite failed")
		return
	}

	logger.Info.Println("router initializing")
	router := gin.Default()
	ListComposite.Handler.Register(router)
	if err := router.Run(":8081"); err != nil {
		log.Println("Error in run project ==> ", err.Error())
		return
	}
}

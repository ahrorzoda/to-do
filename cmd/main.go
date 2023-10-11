package main

import (
	"github.com/ahrorzoda/to-do/internal/composites"
	"github.com/ahrorzoda/to-do/pkg/config"
	logger "github.com/ahrorzoda/to-do/pkg/logger"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	setConfig := config.InitJsonConfigs().Postgres
	postgresConfig, err := composites.NewConfigComposite(setConfig)
	if err != nil {
		logger.Error.Println("config composite failed")
		return
	}
	postgresDBC, err := composites.NewPostgresDBComposite(postgresConfig)
	if err != nil {
		logger.Error.Println("postgres composite failed")
		return
	}
	listComposite, err := composites.NewListComposite(postgresDBC)
	if err != nil {
		logger.Error.Println("List composite failed")
		return
	}
	listComposite.Handler.Register(router)
	if err := router.Run(":8080"); err != nil {
		log.Println("Error in run project ==> ", err.Error())
		return
	}
}

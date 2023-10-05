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
	log.Println("hello")

	logger.Info.Println("config initializing")
	cfg := config.InitJsonConfigs().Postgres

	logger.Info.Printf("postgres composite initializing")
	postgresDBC, err := composites.NewPostgresDBComposite(&cfg)
	if err != nil {
		logger.Error.Println("postgres composite failed")
		return
	}

	logger.Info.Printf("list composite initializing")
	listComposite, err := composites.NewListComposite(postgresDBC)
	if err != nil {
		logger.Error.Println("list composite failed")
		return
	}

	logger.Info.Println("router initializing")
	router := gin.Default()

	listComposite.Handler.Register(router)
}

package composites

import (
	"github.com/ahrorzoda/to-do/pkg/config"
	"github.com/ahrorzoda/to-do/pkg/database"
	lg "github.com/ahrorzoda/to-do/pkg/logger"
	"gorm.io/gorm"
)

type PostgresDBComposite struct {
	db *gorm.DB
}

func NewPostgresDBComposite(config *config.Postgres) (*PostgresDBComposite, error) {
	conn, err := database.Connection(config)
	if err != nil {
		lg.Error.Printf("Error in connect db ==> %s", err.Error())
		return nil, err
	}
	return &PostgresDBComposite{db: conn}, nil
}

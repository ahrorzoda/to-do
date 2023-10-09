package composites

import (
	"github.com/ahrorzoda/to-do/pkg/database"
	lg "github.com/ahrorzoda/to-do/pkg/logger"
	"gorm.io/gorm"
)

type PostgresDBComposite struct {
	db *gorm.DB
}

func NewPostgresDBComposite(compositeCfg *ConfigComposite) (*PostgresDBComposite, error) {
	conn, err := database.Connection(&compositeCfg.config)
	if err != nil {
		lg.Error.Printf("Error in connect db ==> %s", err.Error())
		return nil, err
	}
	return &PostgresDBComposite{db: conn}, nil
}

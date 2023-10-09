package composites

import (
	"github.com/ahrorzoda/to-do/pkg/config"
)

type ConfigComposite struct {
	config config.Postgres
}

func NewConfigComposite(config config.Postgres) (*ConfigComposite, error) {
	return &ConfigComposite{config: config}, nil
}

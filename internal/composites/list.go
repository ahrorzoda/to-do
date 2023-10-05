package composites

import (
	"github.com/ahrorzoda/to-do/internal/adapters/api"
	list2 "github.com/ahrorzoda/to-do/internal/adapters/api/list"
	list3 "github.com/ahrorzoda/to-do/internal/adapters/db/list"
	"github.com/ahrorzoda/to-do/internal/domain/list"
)

type ListComposite struct {
	Storage list.Storage
	Service list2.Service
	Handler api.Handler
}

func NewListComposite(postgresComposite *PostgresDBComposite) (*ListComposite, error) {
	storage := list3.NewStorage(postgresComposite.db)
	service := list.NewService(storage)
	handler := list2.NewHandler(service)
	return &ListComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}

package list

import (
	"context"
	"github.com/ahrorzoda/to-do/internal/domain/list"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *list.List
	GetAll(ctx context.Context, limit, offset int) []*list.List
	Create(ctx context.Context, dto *list.CreateList) *list.List
}

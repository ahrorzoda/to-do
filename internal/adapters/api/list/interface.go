package list

import (
	"context"
	"github.com/ahrorzoda/to-do/internal/domain/list"
)

type Service interface {
	GetListByUUID(ctx context.Context, uuid string) *list.List
	GetLimOffList(ctx context.Context, limit, offset int) []*list.List
	CreateLists(ctx context.Context, dto *list.CreateList) *list.List
	UpdateList(ctx context.Context, uuid string) *list.List
	DeleteList(ctx context.Context, uuid string) error
}

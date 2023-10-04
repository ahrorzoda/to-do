package list

import (
	"context"
	"github.com/ahrorzoda/to-do/internal/adapters/api/list"
)

type Service interface {
	GetListByUUID(ctx context.Context, uuid string) *List
	GetLimOffList(ctx context.Context, limit, offset int) []*List
	CreateList(ctx context.Context, dto *CreateList) *List
}

type service struct {
	storage Storage
}

func NewService(storage Storage) list.Service {
	return &service{storage: storage}
}

func (s *service) GetListByUUID(ctx context.Context, uuid string) *List {
	return s.storage.GetListByUUID(uuid)
}

func (s *service) GetLimOffList(ctx context.Context, limit, offset int) []*List {
	return s.storage.GetLimOffList(limit, offset)
}

func (s *service) CreateLists(ctx context.Context, dto *CreateList) *List {

	return s.storage.CreateLists(dto)
}

func (s *service) UpdateList(ctx context.Context, uuid string) *List {
	return s.storage.UpdateList(uuid)
}

func (s *service) DeleteList(ctx context.Context, uuid string) error {
	return s.storage.DeleteList(uuid)
}

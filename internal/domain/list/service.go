package list

import (
	"github.com/ahrorzoda/to-do/internal/adapters/api/list"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetListByUUID(ctx *gin.Context, uuid string) *List
	GetLimOffList(ctx *gin.Context, limit, offset int) []*List
	CreateList(ctx *gin.Context, dto *CreateList) *List
}

type service struct {
	storage Storage
}

func NewService(storage Storage) list.Service {
	return &service{storage: storage}
}

func (s *service) GetListByUUID(ctx *gin.Context, uuid string) *List {
	return s.storage.GetListByUUID(uuid)
}

func (s *service) GetLimOffList(ctx *gin.Context, limit, offset int) []*List {
	return s.storage.GetLimOffList(limit, offset)
}

func (s *service) CreateLists(ctx *gin.Context, dto *CreateList) *List {
	return s.storage.CreateList(dto)
}

func (s *service) UpdateList(ctx *gin.Context, uuid string) *List {
	return s.storage.UpdateList(uuid)
}

func (s *service) DeleteList(ctx *gin.Context, uuid string) error {
	return s.storage.DeleteList(uuid)
}

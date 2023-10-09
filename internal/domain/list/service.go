package list

import (
	logger "github.com/ahrorzoda/to-do/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service interface {
	GetListByUUID(ctx *gin.Context)
	GetLimOffList(ctx *gin.Context)
	CreateList(ctx *gin.Context)
	UpdateList(ctx *gin.Context)
	DeleteList(ctx *gin.Context)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) GetListByUUID(ctx *gin.Context) {
	//s.storage.GetListByUUIDStorage(uuid)
	return
}

func (s *service) GetLimOffList(ctx *gin.Context) {
	//s.storage.GetLimOffList(limit, offset)
	return
}

func (s *service) CreateList(ctx *gin.Context) {
	var lst Task
	if err := ctx.ShouldBind(&lst); err != nil {
		logger.Error.Println("CreateList(): Error in Bind json ==> %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list := s.storage.CreateList(ctx, &lst)
	ctx.JSON(http.StatusOK, gin.H{"message": "Успешно добавлено!", "list": list})
}

func (s *service) UpdateList(ctx *gin.Context) {
	//s.storage.UpdateList(uuid)
	return
}

func (s *service) DeleteList(ctx *gin.Context) {
	//s.storage.DeleteList(uuid)
	return
}

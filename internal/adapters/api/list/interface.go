package list

import (
	"github.com/ahrorzoda/to-do/internal/domain/list"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetListByUUID(ctx *gin.Context, uuid string) *list.List
	GetLimOffList(ctx *gin.Context, limit, offset int) []*list.List
	CreateLists(ctx *gin.Context, dto *list.CreateList) *list.List
	UpdateList(ctx *gin.Context, uuid string) *list.List
	DeleteList(ctx *gin.Context, uuid string) error
}

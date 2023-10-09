package List

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetListByUUID(ctx *gin.Context)
	GetLimOffList(ctx *gin.Context)
	CreateList(ctx *gin.Context)
	UpdateList(ctx *gin.Context)
	DeleteList(ctx *gin.Context)
}

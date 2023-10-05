package list

import (
	"github.com/ahrorzoda/to-do/internal/adapters/api"
	"github.com/gin-gonic/gin"
)

type handler struct {
	listService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{listService: service}
}

func (h *handler) Register(rout *gin.Engine) {
	rout.GET("/list/get/all", h.listService.GetListByUUID)
	panic("implement me")
}

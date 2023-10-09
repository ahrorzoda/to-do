package List

import (
	"github.com/ahrorzoda/to-do/internal/adapters/api"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ListService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{ListService: service}
}

func (h *handler) Register(rout *gin.Engine) {
	//NOTE: calling all routes
	rout.POST("/todo", h.ListService.CreateList)
}

package list

import "github.com/gin-gonic/gin"

type Storage interface {
	GetListByUUIDStorage(uuid string) *Task
	GetLimOffList(limit, offset int) []*Task
	CreateList(ctx *gin.Context, list *Task) *Task
	UpdateList(uuid string) *Task
	DeleteList(uuid string) error
}

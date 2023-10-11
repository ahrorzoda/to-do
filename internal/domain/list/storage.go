package list

import "github.com/gin-gonic/gin"

type Storage interface {
	GetListByUUIDStorage(uuid string) *Task
	GetLists(ctx *gin.Context) []*Task
	CreateList(ctx *gin.Context, list *Task) *Task
	UpdateList(uuid string) *Task
	DeleteList(uuid string) error
}

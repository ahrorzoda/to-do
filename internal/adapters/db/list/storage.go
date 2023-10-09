package List

import (
	"github.com/ahrorzoda/to-do/internal/domain/list"
	logger "github.com/ahrorzoda/to-do/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type ListStorage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) list.Storage {
	return &ListStorage{db: db}
}

func (ls *ListStorage) GetListByUUIDStorage(uuid string) *list.Task {
	return nil
}

func (ls *ListStorage) GetLimOffList(limit, offset int) []*list.Task {
	return nil
}

//NOTE: добавление новых задач

func (ls *ListStorage) CreateList(ctx *gin.Context, lst *list.Task) *list.Task {
	err := ls.db.Create(&lst).Error
	if err != nil {
		logger.Error.Println("CreateList(): Error in create list ==> %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	return lst
}

func (ls *ListStorage) UpdateList(uuid string) *list.Task {
	//TODO implement me
	panic("implement me")
	return nil
}

func (ls *ListStorage) DeleteList(uuid string) error {
	return nil
}

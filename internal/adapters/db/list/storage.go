package list

import (
	"github.com/ahrorzoda/to-do/internal/domain/list"
	"gorm.io/gorm"
)

type listStorage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) list.Storage {
	return &listStorage{db: db}
}

func (ls *listStorage) GetListByUUID(uuid string) *list.List {
	//TODO implement me
	panic("implement me")
	return nil
}

func (ls *listStorage) GetLimOffList(limit, offset int) []*list.List {
	return nil
}

func (ls *listStorage) CreateList(list *list.CreateList) *list.List {
	//TODO implement me
	panic("implement me")
	return nil
}

func (ls *listStorage) UpdateList(uuid string) *list.List {
	//TODO implement me
	panic("implement me")
	return nil
}

func (ls *listStorage) DeleteList(uuid string) error {
	return nil
}

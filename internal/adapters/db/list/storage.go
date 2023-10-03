package list

import "github.com/ahrorzoda/to-do/internal/domain/list"

type listStorage struct {
}

func NewStorage() list.List {
	return &listStorage{}
}

func (ls *listStorage) GetOneList(uuid string) *list.List {
	return nil
}

func (ls *listStorage) GetLimOffList(limit, offset int) []*list.List {
	return nil
}

func (ls *listStorage) CreateList(list *list.List) *list.List {
	return nil
}

func (ls *listStorage) DeleteList(list *list.List) error {
	return nil
}

package list

type Storage interface {
	GetListByUUID(uuid string) *List
	GetLimOffList(limit, offset int) []*List
	CreateList(list *CreateList) *List
	UpdateList(uuid string) *List
	DeleteList(uuid string) error
}

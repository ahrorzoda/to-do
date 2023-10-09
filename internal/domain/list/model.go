package list

import "time"

type List struct {
	Title       string `gorm:"column:title" json:"title,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	Status      bool   `gorm:"column:status" json:"status,omitempty"`
}

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id,omitempty"`
	Title       string    `gorm:"column:title" json:"title,omitempty"`
	Description string    `gorm:"column:description" json:"description,omitempty"`
	Status      bool      `gorm:"column:status" json:"status,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at; default:true" json:"created_at,omitempty"`
}

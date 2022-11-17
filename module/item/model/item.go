package todomodel

import (
	"errors"
)

var (
	ErrTitleCannotBeBlank       = errors.New("title can not be blank")
	ErrItemNotFound             = errors.New("item not found")
	ErrCannotUpdateFinishedItem = errors.New("can not update finished item")
)

// @Description User account information
// @Description with user id and username
type ToDoItem struct {
	Id        string `json:"id" gorm:"column:id;"`
	Title     string `json:"title" gorm:"column:title;"`
	Completed bool   `json:"status" gorm:"column:completed;"`
	IsShow    bool   `json:"is_show" gorm:"column:is_show;"`
}

func (ToDoItem) TableName() string { return "todolist_list" }

func (item ToDoItem) Validate() error {
	if item.Title == "" {
		return ErrTitleCannotBeBlank
	}

	return nil
}

type DataPaging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *DataPaging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
}

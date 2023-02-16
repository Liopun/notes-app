package notes

import (
	"fmt"
)

const (
	validationError = "%s did not provide required value(s)"
)

type NotesList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type NotesItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Archived    bool   `json:"archived" db:"archived"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (inp UpdateListInput) Validate() (err error) {
	if inp.Title == nil && inp.Description == nil {
		err = fmt.Errorf(validationError, "update list input")
	}

	return
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Archived    *bool   `json:"archived"`
}

func (inp UpdateItemInput) Validate() (err error) {
	if inp.Title == nil && inp.Description == nil && inp.Archived == nil {
		err = fmt.Errorf(validationError, "update item input")
	}

	return
}

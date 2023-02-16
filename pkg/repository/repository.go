package repository

import (
	"github.com/Liopun/notes-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user notes.User) (int, error)
	GetUser(username, password string) (notes.User, error)
}

type NotesList interface {
	Create(userId int, list notes.NotesList) (int, error)
	GetAll(userId int) ([]notes.NotesList, error)
	GetById(userId, listId int) (notes.NotesList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, inp notes.UpdateListInput) error
}

type NotesItem interface {
	Create(userId int, list notes.NotesItem) (int, error)
	GetAll(userId, listId int) ([]notes.NotesItem, error)
	GetById(userId, itemId int) (notes.NotesItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, inp notes.UpdateItemInput) error
}

type Repository struct {
	Authorization
	NotesList
	NotesItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		// NotesList: ,
		// NotesItem: ,
	}
}

package service

import (
	"time"

	"github.com/Liopun/notes-app"
	"github.com/Liopun/notes-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user notes.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type NotesList interface {
	Create(userId int, list notes.NotesList) (int, error)
	GetAll(userId int) ([]notes.NotesList, error)
	GetById(userId, listId int) (notes.NotesList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, inp notes.UpdateListInput) error
}

type NotesItem interface {
	Create(userId, listId int, item notes.NotesItem) (int, error)
	GetAll(userId, listId int) ([]notes.NotesItem, error)
	GetById(userId, itemId int) (notes.NotesItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, inp notes.UpdateItemInput) error
}

type Service struct {
	Authorization
	NotesList
	NotesItem
}

type Deps struct {
	Repos        *repository.Repository
	PasswordSalt string
	TokenTTL     time.Duration
	SigningKey   string
}

func NewService(deps Deps) *Service {
	authService := NewAuthService(deps.Repos.Authorization, deps.PasswordSalt, deps.SigningKey, deps.TokenTTL)
	notesListService := NewNotesListService(deps.Repos.NotesList)
	notesItemService := NewNotesItemService(deps.Repos.NotesItem, deps.Repos.NotesList)

	return &Service{
		Authorization: authService,
		NotesList:     notesListService,
		NotesItem:     notesItemService,
	}
}

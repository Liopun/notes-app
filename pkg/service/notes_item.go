package service

import (
	"github.com/Liopun/notes-app"
	"github.com/Liopun/notes-app/pkg/repository"
)

type NotesItemService struct {
	repo     repository.NotesItem
	listRepo repository.NotesList
}

func NewNotesItemService(repo repository.NotesItem, listRepo repository.NotesList) *NotesItemService {
	return &NotesItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *NotesItemService) Create(userId, listId int, item notes.NotesItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return -1, err
	}

	return s.repo.Create(listId, item)
}

func (s *NotesItemService) GetAll(userId, listId int) ([]notes.NotesItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *NotesItemService) GetById(userId, itemId int) (notes.NotesItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *NotesItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *NotesItemService) Update(userId, itemId int, inp notes.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, inp)
}

package service

import (
	"github.com/Liopun/notes-app"
	"github.com/Liopun/notes-app/pkg/repository"
)

type NotesListService struct {
	repo repository.NotesList
}

func NewNotesListService(repo repository.NotesList) *NotesListService {
	return &NotesListService{repo: repo}
}

func (s *NotesListService) Create(userId int, list notes.NotesList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *NotesListService) GetAll(userId int) ([]notes.NotesList, error) {
	return s.repo.GetAll(userId)
}

func (s *NotesListService) GetById(userId, listId int) (notes.NotesList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *NotesListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *NotesListService) Update(userId, listId int, inp notes.UpdateListInput) error {
	if err := inp.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, inp)
}

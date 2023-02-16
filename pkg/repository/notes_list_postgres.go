package repository

import (
	"fmt"
	"strings"

	"github.com/Liopun/notes-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type NotesListPostgres struct {
	db *sqlx.DB
}

func NewNotesListPostgres(db *sqlx.DB) *NotesListPostgres {
	return &NotesListPostgres{db: db}
}

func (r *NotesListPostgres) Create(userId int, list notes.NotesList) (int, error) {
	var id int

	tx, err := r.db.Begin()
	if err != nil {
		return -1, err
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", notesListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return -1, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	return id, tx.Commit()
}

func (r *NotesListPostgres) GetAll(userId int) ([]notes.NotesList, error) {
	var lists []notes.NotesList

	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		notesListsTable,
		usersListsTable,
	)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *NotesListPostgres) GetById(userId, listId int) (notes.NotesList, error) {
	var list notes.NotesList

	query := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		notesListsTable,
		usersListsTable,
	)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *NotesListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		notesListsTable,
		usersListsTable,
	)

	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *NotesListPostgres) Update(userId, listId int, inp notes.UpdateListInput) error {
	qValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inp.Title != nil {
		qValues = append(qValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *inp.Title)
		argId++
	}

	if inp.Description != nil {
		qValues = append(qValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *inp.Description)
		argId++
	}

	qString := strings.Join(qValues, ", ")

	query := fmt.Sprintf(
		"UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		notesListsTable,
		qString,
		usersListsTable,
		argId,
		argId+1,
	)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}

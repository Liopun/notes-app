package repository

import (
	"fmt"
	"strings"

	"github.com/Liopun/notes-app"
	"github.com/jmoiron/sqlx"
)

type NotesItemPostgres struct {
	db *sqlx.DB
}

func NewNotesItemPostgres(db *sqlx.DB) *NotesItemPostgres {
	return &NotesItemPostgres{db: db}
}

func (r *NotesItemPostgres) Create(userId int, item notes.NotesItem) (int, error) {
	var itemId int

	tx, err := r.db.Begin()
	if err != nil {
		return -1, err
	}

	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", notesItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return -1, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, userId, itemId)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	return itemId, tx.Commit()
}

func (r *NotesItemPostgres) GetAll(userId, listId int) ([]notes.NotesItem, error) {
	var items []notes.NotesItem

	query := fmt.Sprintf(
		`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		notesItemsTable,
		listsItemsTable,
		usersListsTable,
	)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *NotesItemPostgres) GetById(userId, itemId int) (notes.NotesItem, error) {
	var item notes.NotesItem

	query := fmt.Sprintf(
		`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		notesItemsTable,
		listsItemsTable,
		usersListsTable,
	)

	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *NotesItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		notesItemsTable,
		listsItemsTable,
		usersListsTable,
	)

	_, err := r.db.Exec(query, userId, itemId)

	return err
}

func (r *NotesItemPostgres) Update(userId, itemId int, inp notes.UpdateItemInput) error {
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

	if inp.Archived != nil {
		qValues = append(qValues, fmt.Sprintf("archived=$%d", argId))
		args = append(args, *inp.Archived)
		argId++
	}

	qString := strings.Join(qValues, ", ")

	query := fmt.Sprintf(
		`UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		notesItemsTable,
		qString,
		listsItemsTable,
		usersListsTable,
		argId,
		argId+1,
	)

	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)

	return err
}

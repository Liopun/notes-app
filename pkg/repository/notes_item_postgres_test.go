package repository

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Liopun/notes-app"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestNotesItemPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewNotesItemPostgres(sqlxDb)

	type args struct {
		listId int
		item   notes.NotesItem
	}

	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		input   args
		mock    mockBehavior
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			input: args{
				listId: 1,
				item: notes.NotesItem{
					Title:       "test title",
					Description: "test description",
				},
			},
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)

				mock.ExpectQuery("INSERT INTO notes_items").WithArgs(args.item.Title, args.item.Description).WillReturnRows(rows)
				mock.ExpectExec("INSERT INTO lists_items").WithArgs(args.listId, id).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			want: 2,
		},
		{
			name: "Empty Fields",
			input: args{
				listId: 1,
				item: notes.NotesItem{
					Title:       "",
					Description: "",
				},
			},
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id).RowError(0, errors.New("insert error"))
				mock.ExpectQuery("INSERT INTO notes_items").WithArgs(args.item.Title, args.item.Description).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "2nd Insert Failure",
			input: args{
				listId: 1,
				item: notes.NotesItem{
					Title:       "title",
					Description: "description",
				},
			},
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO notes_items").WithArgs(args.item.Title, args.item.Description).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO lists_items").WithArgs(args.listId, id).WillReturnError(errors.New("insert error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			got, err := r.Create(tt.input.listId, tt.input.item)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestNotesItemPostgres_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewNotesItemPostgres(sqlxDb)

	type args struct {
		userId int
		listId int
	}

	tests := []struct {
		name    string
		input   args
		mock    func()
		want    []notes.NotesItem
		wantErr bool
	}{
		{
			name:  "OK",
			input: args{userId: 1, listId: 1},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "archived"}).
					AddRow(1, "title1", "description1", false).
					AddRow(2, "title2", "description2", true).
					AddRow(3, "title3", "description3", true)

				mock.ExpectQuery("SELECT (.+) FROM notes_items ti INNER JOIN lists_items li on (.+) INNER JOIN users_lists ul on (.+) WHERE (.+)").WithArgs(1, 1).WillReturnRows(rows)
			},
			want: []notes.NotesItem{
				{Id: 1, Title: "title1", Description: "description1", Archived: false},
				{Id: 2, Title: "title2", Description: "description2", Archived: true},
				{Id: 3, Title: "title3", Description: "description3", Archived: true},
			},
		},
		{
			name:  "No Items",
			input: args{userId: 1, listId: 1},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "archived"})
				mock.ExpectQuery("SELECT (.+) FROM notes_items ti INNER JOIN lists_items li on (.+) INNER JOIN users_lists ul on (.+) WHERE (.+)").WithArgs(1, 1).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetAll(tt.input.userId, tt.input.listId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestNotesItemPostgres_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewNotesItemPostgres(sqlxDb)

	type args struct {
		userId int
		itemId int
	}

	tests := []struct {
		name    string
		input   args
		mock    func()
		want    notes.NotesItem
		wantErr bool
	}{
		{
			name:  "OK",
			input: args{userId: 1, itemId: 1},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "archived"}).AddRow(1, "title1", "description1", false)

				mock.ExpectQuery("SELECT (.+) FROM notes_items ti INNER JOIN lists_items li on (.+) INNER JOIN users_lists ul on (.+) WHERE (.+)").WithArgs(1, 1).WillReturnRows(rows)
			},
			want: notes.NotesItem{Id: 1, Title: "title1", Description: "description1", Archived: false},
		},
		{
			name:  "No Item",
			input: args{userId: -1, itemId: -1},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "archived"})
				mock.ExpectQuery("SELECT (.+) FROM notes_items ti INNER JOIN lists_items li on (.+) INNER JOIN users_lists ul on (.+) WHERE (.+)").WithArgs(-1, -1).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetById(tt.input.userId, tt.input.itemId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestNotesItemPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewNotesItemPostgres(sqlxDb)

	type args struct {
		userId int
		itemId int
	}

	tests := []struct {
		name    string
		input   args
		mock    func()
		want    notes.NotesItem
		wantErr bool
	}{
		{
			name:  "OK",
			input: args{userId: 1, itemId: 1},
			mock: func() {
				mock.ExpectExec("DELETE FROM notes_items ti USING lists_items li, users_lists ul WHERE (.+)").
					WithArgs(1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

			},
		},
		{
			name:  "No Item",
			input: args{userId: -1, itemId: -1},
			mock: func() {
				mock.ExpectExec("DELETE FROM notes_items ti USING lists_items li, users_lists ul WHERE (.+)").
					WithArgs(-1, -1).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.Delete(tt.input.userId, tt.input.itemId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestNotesItemPostgres_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewNotesItemPostgres(sqlxDb)

	type args struct {
		userId int
		itemId int
		input  notes.UpdateItemInput
	}

	tests := []struct {
		name    string
		input   args
		mock    func()
		wantErr bool
	}{
		{
			name: "OK_Archived",
			input: args{
				userId: 1,
				itemId: 1,
				input: notes.UpdateItemInput{
					Title:       stringPointer("updated title"),
					Description: stringPointer("updated desc"),
					Archived:    boolPointer(true),
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE notes_items ti SET (.+) FROM lists_items li, users_lists ul WHERE (.+)").
					WithArgs("updated title", "updated desc", true, 1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

			},
		},
		{
			name: "OK_NotArchived",
			input: args{
				userId: 1,
				itemId: 1,
				input: notes.UpdateItemInput{
					Title:       stringPointer("updated title"),
					Description: stringPointer("updated desc"),
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE notes_items ti SET (.+) FROM lists_items li, users_lists ul WHERE (.+)").
					WithArgs("updated title", "updated desc", 1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

			},
		},
		{
			name: "OK_NoDoneAndNoDescription",
			input: args{
				userId: 1,
				itemId: 1,
				input: notes.UpdateItemInput{
					Title: stringPointer("updated title"),
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE notes_items ti SET (.+) FROM lists_items li, users_lists ul WHERE (.+)").
					WithArgs("updated title", 1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

			},
		},
		{
			name:  "OK_NoInput",
			input: args{userId: 1, itemId: 1},
			mock: func() {
				mock.ExpectExec("UPDATE notes_items ti SET FROM lists_items li, users_lists ul WHERE (.+)").
					WithArgs(1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.Update(tt.input.userId, tt.input.itemId, tt.input.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func stringPointer(s string) *string {
	return &s
}

func boolPointer(b bool) *bool {
	return &b
}

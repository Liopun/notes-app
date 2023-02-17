package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Liopun/notes-app"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestNotesListPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewNotesListPostgres(sqlxDb)

	type args struct {
		userId int
		item   notes.NotesList
	}

	tests := []struct {
		name    string
		input   args
		mock    func()
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			input: args{
				userId: 1,
				item: notes.NotesList{
					Title:       "test title",
					Description: "test description",
				},
			},
			mock: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

				mock.ExpectQuery("INSERT INTO notes_lists").WithArgs("test title", "test description").WillReturnRows(rows)
				mock.ExpectExec("INSERT INTO users_lists").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			input: args{
				userId: 1,
				item: notes.NotesList{
					Title:       "",
					Description: "",
				},
			},
			mock: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO notes_lists").WithArgs("", "").WillReturnRows(rows)

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.Create(tt.input.userId, tt.input.item)
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

func TestNotesListPostgres_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewNotesListPostgres(sqlxDb)

	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		input   args
		mock    func()
		want    []notes.NotesList
		wantErr bool
	}{
		{
			name: "OK",
			input: args{
				userId: 1,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description"}).
					AddRow(1, "title1", "description1").
					AddRow(2, "title2", "description2").
					AddRow(3, "title3", "description3")

				mock.ExpectQuery("SELECT (.+) FROM notes_lists tl INNER JOIN users_lists ul on (.+) WHERE (.+)").WithArgs(1).WillReturnRows(rows)
			},
			want: []notes.NotesList{
				{Id: 1, Title: "title1", Description: "description1"},
				{Id: 2, Title: "title2", Description: "description2"},
				{Id: 3, Title: "title3", Description: "description3"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetAll(tt.input.userId)
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

func TestNotesListPostgres_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewNotesListPostgres(sqlxDb)

	type args struct {
		listId int
		userId int
	}

	tests := []struct {
		name    string
		input   args
		mock    func()
		want    notes.NotesList
		wantErr bool
	}{
		{
			name: "OK",
			input: args{
				userId: 1,
				listId: 1,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description"}).AddRow(1, "title1", "description1")

				mock.ExpectQuery("SELECT (.+) FROM notes_lists tl INNER JOIN users_lists ul on (.+) WHERE (.+)").WithArgs(1, 1).WillReturnRows(rows)
			},
			want: notes.NotesList{Id: 1, Title: "title1", Description: "description1"},
		},
		{
			name: "No List",
			input: args{
				userId: -1,
				listId: -1,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description"})

				mock.ExpectQuery("SELECT (.+) FROM notes_lists tl INNER JOIN users_lists ul on (.+) WHERE (.+)").WithArgs(-1, -1).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetById(tt.input.userId, tt.input.listId)
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

func TestNotesListPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewNotesListPostgres(sqlxDb)

	type args struct {
		listId int
		userId int
	}

	tests := []struct {
		name    string
		input   args
		mock    func()
		wantErr bool
	}{
		{
			name: "OK",
			input: args{
				userId: 1,
				listId: 1,
			},
			mock: func() {
				mock.ExpectExec("DELETE FROM notes_lists tl USING users_lists ul WHERE (.+)").
					WithArgs(1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "No List",
			input: args{
				userId: -1,
				listId: -1,
			},
			mock: func() {
				mock.ExpectExec("DELETE FROM notes_lists tl USING users_lists ul WHERE (.+)").
					WithArgs(-1, -1).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.Delete(tt.input.userId, tt.input.listId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestNotesListPostgres_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewNotesListPostgres(sqlxDb)

	type args struct {
		userId int
		listId int
		input  notes.UpdateListInput
	}

	tests := []struct {
		name    string
		input   args
		mock    func()
		wantErr bool
	}{
		{
			name: "OK",
			input: args{
				userId: 1,
				listId: 1,
				input: notes.UpdateListInput{
					Title:       stringPointer("updated title"),
					Description: stringPointer("updated descr"),
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE notes_lists tl SET (.+) FROM users_lists ul WHERE (.+)").
					WithArgs("updated title", "updated descr", 1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "OK_NoDescription",
			input: args{
				userId: 1,
				listId: 1,
				input: notes.UpdateListInput{
					Title: stringPointer("updated title"),
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE notes_lists tl SET (.+) FROM users_lists ul WHERE (.+)").
					WithArgs("updated title", 1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "OK_NoTitle",
			input: args{
				userId: 1,
				listId: 1,
				input: notes.UpdateListInput{
					Description: stringPointer("updated desc"),
				},
			},
			mock: func() {
				mock.ExpectExec("UPDATE notes_lists tl SET (.+) FROM users_lists ul WHERE (.+)").
					WithArgs("updated desc", 1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "OK_NoInput",
			input: args{
				userId: 1,
				listId: 1,
			},
			mock: func() {
				mock.ExpectExec("UPDATE notes_lists tl SET FROM users_lists ul WHERE (.+)").
					WithArgs(1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.Update(tt.input.userId, tt.input.listId, tt.input.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

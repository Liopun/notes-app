package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Liopun/notes-app"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestAuthPostgres_CreateUse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewAuthPostgres(sqlxDb)

	tests := []struct {
		name    string
		mock    func()
		input   notes.User
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO users").WithArgs("Test", "test", "password").WillReturnRows(rows)
			},
			input: notes.User{
				Name:     "Test",
				Username: "test",
				Password: "password",
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").WithArgs("Test", "test", "").WillReturnRows(rows)
			},
			input: notes.User{
				Name:     "Test",
				Username: "test",
				Password: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.CreateUser(tt.input)
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

func TestAuthPostgres_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured '%s' was not expected for stub database connection", err)
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	defer sqlxDb.Close()

	r := NewAuthPostgres(sqlxDb)

	type args struct {
		username string
		password string
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		want    notes.User
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "username", "password"}).AddRow(1, "Test", "test", "password")
				mock.ExpectQuery("SELECT (.+) FROM users").WithArgs("test", "password").WillReturnRows(rows)
			},
			input: args{"test", "password"},
			want: notes.User{
				Id:       1,
				Name:     "Test",
				Username: "test",
				Password: "password",
			},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "username", "password"})
				mock.ExpectQuery("SELECT (.+) FROM users").WithArgs("not", "found").WillReturnRows(rows)
			},
			input:   args{"not", "found"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetUser(tt.input.username, tt.input.password)
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

package mocksql

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMock(t *testing.T) (gdb *gorm.DB, mock sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	gdb, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))
	assert.Nil(t, err)
	return
}

func TestRepository_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		isError bool
		rows    *sqlmock.Rows
	}{
		{
			name:    "GetAll failed",
			query:   "^SELECT (.+)FROM (.+)user",
			isError: true,
		},
		{
			name:    "GetAll success",
			query:   "^SELECT (.+)FROM (.+)user",
			isError: false,
			rows:    sqlmock.NewRows([]string{"id", "name", "age", "email"}).AddRow(int64(10), "Enigma", int64(21), "enigma@enigmacamp.com"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gdb, mock := NewMock(t)
			r := NewUserRepository(gdb)

			if tt.isError {
				mock.ExpectQuery(tt.query).WillReturnError(errors.New("general-error"))
			} else {
				mock.ExpectQuery(tt.query).WillReturnRows(tt.rows)
			}

			users, err := r.GetAll()
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, users)
			}
		})
	}

	// t.Run("failed", func(t *testing.T) {
	// 	gdb, mock := NewMock(t)
	// 	r := NewUserRepository(gdb)

	// 	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user"`)).WillReturnError(errors.New("general-error"))

	// 	_, err := r.GetAll()
	// 	assert.Error(t, err)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, "general-error", err.Error())
	// })
	// t.Run("success", func(t *testing.T) {
	// 	gdb, mock := NewMock(t)
	// 	r := NewUserRepository(gdb)

	// 	rows := sqlmock.NewRows([]string{"id", "name", "age", "email"}).
	// 		AddRow(int64(10), "Enigma", int64(21), "enigma@enigmacamp.com").
	// 		AddRow(int64(11), "Enigma 11", int64(22), "enigma11@enigmacamp.com").
	// 		AddRow(int64(12), "Enigma 12", int64(23), "enigma12@enigmacamp.com")

	// 	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user"`)).WillReturnRows(rows)

	// 	users, err := r.GetAll()
	// 	assert.NoError(t, err)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, 3, len(users))
	// 	assert.Equal(t, "Enigma 11", users[1].Name)
	// })

	// t.Run("success", func(t *testing.T) {
	// 	gdb, mock := NewMock(t)
	// 	r := NewUserRepository(gdb)

	// 	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user"`)).
	// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "email"}).
	// 			AddRow(1, "test", 24, "test@email.com"))

	// 	entities, err := r.GetAll()
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, 1, len(entities))

	// })
}

func TestRepository_GetByID(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		isError bool
		rows    *sqlmock.Rows
		paramID int64
	}{
		{
			name:    "GetByID failed",
			query:   "^SELECT (.+)FROM (.+)user",
			isError: true,
			paramID: 10,
		},
		{
			name:    "GetByID success",
			query:   "^SELECT (.+)FROM (.+)user",
			isError: false,
			rows:    sqlmock.NewRows([]string{"id", "name", "age", "email"}).AddRow(int64(10), "Enigma", int64(21), "enigma@enigmacamp.com"),
			paramID: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gdb, mock := NewMock(t)
			r := NewUserRepository(gdb)

			if tt.isError {
				mock.ExpectQuery(tt.query).WithArgs(int64(10)).WillReturnError(errors.New("general-error"))
			} else {
				mock.ExpectQuery(tt.query).WithArgs(int64(10)).WillReturnRows(tt.rows)
			}

			user, err := r.GetByID(tt.paramID)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, user)
			}
		})
	}
}

func TestRepository_Save(t *testing.T) {
	entity := &User{
		Name:  "Test",
		Age:   24,
		Email: "test@email.com",
	}

	t.Run("failed", func(t *testing.T) {
		gdb, mock := NewMock(t)
		r := NewUserRepository(gdb)

		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO (.+)user").
			WithArgs("Test", int64(24), "test@email.com").
			WillReturnError(errors.New("connection error"))
		mock.ExpectRollback()

		err := r.Save(entity)
		assert.Error(t, err)
		assert.NotNil(t, err)
		assert.Equal(t, "connection error", err.Error())
	})

	t.Run("success", func(t *testing.T) {
		gdb, mock := NewMock(t)
		r := NewUserRepository(gdb)

		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO (.+)user").
			WithArgs("Test", int64(24), "test@email.com").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(1))
		mock.ExpectCommit()

		// expectExec untuk update
		//mock.ExpectExec(regexp.QuoteMeta(`UPDATE "user" SET "name" = $1, "age" = $2, "email" = $3 WHERE id = $4)`)).
		// WithArgs("name baru", age baru, email baru, id).WillReturnResult(sqlmock.NewResult(0, 1))

		err := r.Save(entity)
		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}

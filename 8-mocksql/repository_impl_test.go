package mocksql

import (
	"errors"
	"regexp"
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
	t.Run("failed", func(t *testing.T) {
		gdb, mock := NewMock(t)
		r := NewUserRepository(gdb)

		mock.ExpectQuery("^SELECT (.+)FROM (.+)user").
			WillReturnError(errors.New("connection error"))

		_, err := r.GetAll()
		assert.Error(t, err)
		assert.NotNil(t, err)
		assert.Equal(t, "connection error", err.Error())
	})

	t.Run("success", func(t *testing.T) {
		gdb, mock := NewMock(t)
		r := NewUserRepository(gdb)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user"`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "email"}).
				AddRow(1, "test", 24, "test@email.com"))

		entities, err := r.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(entities))

	})
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

		err := r.Save(entity)
		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}

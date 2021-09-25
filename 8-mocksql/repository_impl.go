package mocksql

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() (users []User, err error) {
	err = r.db.Find(&users).Error
	return
}

func (r *repository) GetByID(id int64) (user User, err error) {
	err = r.db.First(&user, id).Error
	return
}

func (r *repository) Save(user *User) (err error) {
	err = r.db.Save(user).Error
	return
}

func (r *repository) Delete(user *User) (err error) {
	err = r.db.Delete(user).Error
	return
}

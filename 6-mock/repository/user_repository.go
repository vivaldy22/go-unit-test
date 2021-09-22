package repository

type UserEntity struct {
	ID   int64
	Name string
	Age  int64
}

type UserRepository interface {
	GetByID(id int64) *UserEntity
}

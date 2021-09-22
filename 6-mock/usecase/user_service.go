package usecase

type UserService interface {
	GetByID(id int64) (User, error)
}

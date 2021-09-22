package repository

type userRepo struct {
	//db
}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (u *userRepo) GetByID(id int64) (entity *UserEntity) {
	// kode repo

	entity = &UserEntity{
		ID:   1,
		Name: "User Satu",
		Age:  21,
	}

	return
}

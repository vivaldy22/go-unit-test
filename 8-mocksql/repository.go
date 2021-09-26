package mocksql

type User struct {
	ID    int64 `gorm:"primaryKey;autoIncrement"` //struct tag gorm
	Name  string
	Age   int64
	Email string
}

func (u *User) TableName() string {
	return "user"
}

type Repository interface {
	GetAll() ([]User, error)
	GetByID(id int64) (User, error)
	Save(*User) error
	Delete(*User) error
}

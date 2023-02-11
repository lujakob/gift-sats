package user

type IUserStore interface {
	GetAll() ([]User, int64, error)
	GetByEmail(string) (*User, error)
	Create(*User) error
}

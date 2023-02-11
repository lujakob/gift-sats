package user

type IUserStore interface {
	GetAll() ([]User, int64, error)
	GetByUsername(string) (*User, error)
	Create(*User) error
}

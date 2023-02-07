package user

type IUserStore interface {
	GetAll() ([]User, error)
	GetByUsername(string) (*User, error)
	Create(*User) error
}

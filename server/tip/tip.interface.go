package tip

type ITipStore interface {
	GetAll() ([]Tip, int64, error)
	Create(*Tip) error
}

package customer

type Repository interface {
	Delete(id string) error
	GetAll() ([]*Customer, error)
	GetByID(id string) (*Customer, error)
	Store(c *Customer) error
	Update(c *Customer) error
}
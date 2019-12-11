package customer

type Service interface {
	Delete(id string) error
	GetAll() ([]*Customer, error)
	GetByID(id string) (*Customer, error)
	Store(u *Customer) error
	Update(u *Customer) error
}

type customerService struct {
	repo Repository
}

func NewCustomerService(repo Repository) Service {
	return &customerService{
		repo: repo,
	}
}

func (svc *customerService) Delete(id string) error { return svc.repo.Delete(id) }

func (svc *customerService) GetAll() ([]*Customer, error) { return svc.repo.GetAll() }

func (svc *customerService) GetByID(id string) (*Customer, error) { return svc.repo.GetByID(id) }

func (svc *customerService) Store(u *Customer) error { return svc.repo.Store(u) }

func (svc *customerService) Update(u *Customer) error { return svc.repo.Update(u) }
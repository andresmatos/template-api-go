package orm

import (
"errors"
	"template-api-go/pkg/customer"

	"github.com/jinzhu/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) customer.Repository {
	return &customerRepo{db}
}

func (p *customerRepo) Delete(id string) error {
	if p.db.Delete(&customer.Customer{}, "customer_id = ?", id).Error != nil {
		return errors.New("error while deleting the customer")
	}
	return nil
}

func (p *customerRepo) GetAll() ([]*customer.Customer, error) {
	customers := make([]*customer.Customer, 0)
	err := p.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (p *customerRepo) GetByID(id string) (*customer.Customer, error) {
	customer := &customer.Customer{}
	if p.db.Where("customer_id = ?", id).First(&customer).Error != nil {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

func (p *customerRepo) Store(pt *customer.Customer) error {
	if p.db.Create(&pt).Error != nil {
		return errors.New("error while creating the customer")
	}
	return nil
}

func (p *customerRepo) Update(pt *customer.Customer) error {
	err := p.db.Model(&pt).Updates(customer.Customer{FirstName: pt.FirstName, LastName: pt.LastName, Email: pt.Email,Address: pt.Address}).Error
	if err != nil {
		return errors.New("error while updating the customer")
	}
	return nil
}
package customer

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Customer struct {
	ID             string    `json:"customer_id" db:"customer_id" gorm:"column:customer_id;primary_key;type:uuid"`
	FirstName      string    `json:"first_name" db:"first_name" gorm:"varchar(50);column:first_name;not null"`
	LastName       string    `json:"last_name" db:"last_name" gorm:"varchar(50);column:last_name"`
	Address        string    `json:"address" db:"address" gorm:"varchar(250)"`
	Email          string    `json:"email" db:"email" gorm:"varchar(200)"`

}

func (Customer) TableName() string {
	return "customers"
}

func (c *Customer) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
}

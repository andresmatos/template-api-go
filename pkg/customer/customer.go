package customer

type Customer struct {
	ID             string    `json:"customer_id" db:"customer_id" gorm:"column:customer_id;primary_key;type:uuid"`
	FirstName      string    `json:"first_name" db:"first_name" gorm:"varchar(50);column:first_name;not null"`
	LastName       string    `json:"last_name" db:"last_name" gorm:"varchar(50);column:last_name"`
	Address        string    `json:"address" db:"address" gorm:"varchar(250)"`
	Age            int8      `json:"age" db:"age"`
	Email          string    `json:"email" db:"email" gorm:"varchar(200)"`
	Gender         int8      `json:"gender" db:"gender" gorm:"not null"`
}

func (Customer) TableName() string {
	return "customers"
}
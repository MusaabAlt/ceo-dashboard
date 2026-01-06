package domain

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name            string    `gorm:"type:varchar(255);not null" json:"name"`
	DefaultCurrency string    `gorm:"type:varchar(3);default:'USD'" json:"default_currency"`
	Timezone        string    `gorm:"type:varchar(50);default:'UTC'" json:"timezone"`
	CreatedAt       time.Time `json:"created_at"`
}

func (Company) TableName() string {
	return "companies"
}

package domain

import "github.com/google/uuid"

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
}

func (Role) TableName() string {
	return "roles"
}

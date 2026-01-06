package repository

import (
	"go-admin/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Company").Preload("Roles").Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Company").Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserRepository) UpdateLastLogin(userID uuid.UUID) error {
	return r.db.Model(&domain.User{}).Where("id = ?", userID).Update("last_login_at", gorm.Expr("NOW()")).Error
}

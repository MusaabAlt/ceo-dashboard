package repository

import (
	"go-admin/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) Create(company *domain.Company) error {
	return r.db.Create(company).Error
}

func (r *CompanyRepository) FindByID(id uuid.UUID) (*domain.Company, error) {
	var company domain.Company
	err := r.db.Where("id = ?", id).First(&company).Error
	return &company, err
}

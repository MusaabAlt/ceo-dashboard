package service

import (
	"errors"
	"go-admin/internal/config"
	"go-admin/internal/domain"
	"go-admin/internal/dto"
	"go-admin/internal/repository"
	"go-admin/internal/util"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo    *repository.UserRepository
	companyRepo *repository.CompanyRepository
	db          *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		userRepo:    repository.NewUserRepository(db),
		companyRepo: repository.NewCompanyRepository(db),
		db:          db,
	}
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	if !util.IsValidEmail(req.Email) {
		return nil, errors.New("invalid email")
	}
	if !util.IsValidPassword(req.Password) {
		return nil, errors.New("password too short")
	}

	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil && existingUser.ID != uuid.Nil {
		return nil, errors.New("email exists")
	}

	var company domain.Company
	var user domain.User

	err := s.db.Transaction(func(tx *gorm.DB) error {
		company = domain.Company{Name: req.CompanyName, DefaultCurrency: "USD", Timezone: "UTC"}
		if err := tx.Create(&company).Error; err != nil {
			return err
		}
		user = domain.User{CompanyID: company.ID, Email: req.Email, FullName: req.FullName, IsActive: true}
		if err := user.SetPassword(req.Password); err != nil {
			return err
		}
		return tx.Create(&user).Error
	})

	if err != nil {
		return nil, err
	}

	token, _ := util.GenerateJWT(user.ID, company.ID, user.Email, config.AppConfig.JWTSecret)
	return &dto.AuthResponse{
		Token:   token,
		User:    dto.UserInfo{ID: user.ID, Email: user.Email, FullName: user.FullName, IsActive: user.IsActive},
		Company: dto.CompanyInfo{ID: company.ID, Name: company.Name, DefaultCurrency: company.DefaultCurrency, Timezone: company.Timezone},
	}, nil
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if !user.IsActive {
		return nil, errors.New("account disabled")
	}
	if err := user.ComparePassword(req.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}
	s.userRepo.UpdateLastLogin(user.ID)
	token, _ := util.GenerateJWT(user.ID, user.CompanyID, user.Email, config.AppConfig.JWTSecret)
	return &dto.AuthResponse{
		Token:   token,
		User:    dto.UserInfo{ID: user.ID, Email: user.Email, FullName: user.FullName, IsActive: user.IsActive},
		Company: dto.CompanyInfo{ID: user.Company.ID, Name: user.Company.Name, DefaultCurrency: user.Company.DefaultCurrency, Timezone: user.Company.Timezone},
	}, nil
}

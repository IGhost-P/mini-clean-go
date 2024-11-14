package service

import (
    "github.com/IGhost-p/mini-clean-go/internal/model"
    "github.com/IGhost-p/mini-clean-go/internal/repository"
)

// UserService는 사용자 관련 비즈니스 로직을 처리하는 인터페이스입니다.
type UserService interface {
    CreateUser(user *model.User) error
    GetAllUsers() ([]*model.User, error)
}

// userService는 UserService 인터페이스의 구현체입니다.
type userService struct {
    userRepo repository.UserRepository
}

// NewUserService는 새로운 UserService를 생성합니다.
func NewUserService(userRepo repository.UserRepository) UserService {
    return &userService{
        userRepo: userRepo,
    }
}

// CreateUser는 새로운 사용자를 생성합니다.
func (s *userService) CreateUser(user *model.User) error {
    return s.userRepo.Create(user)
}

// GetAllUsers는 모든 사용자를 조회합니다.
func (s *userService) GetAllUsers() ([]*model.User, error) {
    return s.userRepo.FindAll()
}
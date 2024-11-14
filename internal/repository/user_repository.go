package repository

import (
    "sync"

		"github.com/IGhost-p/mini-clean-go/internal/model"
)

// UserRepository는 사용자 데이터 접근을 위한 인터페이스입니다.
type UserRepository interface {
    Create(user *model.User) error
    FindAll() ([]*model.User, error)
}

// memoryUserRepository는 메모리에 데이터를 저장하는 구현체입니다.
type memoryUserRepository struct {
    sync.RWMutex // 동시성 처리를 위한 mutex
    users        map[string]*model.User
}

// NewMemoryUserRepository는 새로운 메모리 저장소를 생성합니다.
func NewMemoryUserRepository() UserRepository {
    return &memoryUserRepository{
        users: make(map[string]*model.User),
    }
}

// Create는 새로운 사용자를 저장합니다.
func (r *memoryUserRepository) Create(user *model.User) error {
    r.Lock()
    defer r.Unlock()
    
    r.users[user.ID] = user
    return nil
}

// FindAll은 모든 사용자를 조회합니다.
func (r *memoryUserRepository) FindAll() ([]*model.User, error) {
    r.RLock()
    defer r.RUnlock()
    
    users := make([]*model.User, 0, len(r.users))
    for _, user := range r.users {
        users = append(users, user)
    }
    
    return users, nil
}
package handler

import (
    "encoding/json"
    "net/http"

    "github.com/IGhost-p/mini-clean-go/internal/model"
    "github.com/IGhost-p/mini-clean-go/internal/service"
)

// UserHandler는 HTTP 요청을 처리하는 핸들러입니다.
type UserHandler struct {
    userService service.UserService
}

// NewUserHandler는 새로운 UserHandler를 생성합니다.
func NewUserHandler(userService service.UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
    }
}

// CreateUser는 새로운 사용자를 생성하는 핸들러입니다.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.userService.CreateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// GetUsers는 모든 사용자를 조회하는 핸들러입니다.
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.userService.GetAllUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
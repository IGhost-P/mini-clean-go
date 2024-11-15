package handler

import (
    "encoding/json"
    "net/http"

    "github.com/IGhost-p/mini-clean-go/internal/logger"  // logger 패키지 추가
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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User object"
// @Success 201 {string} string "Created"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [post]
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

    // 사용자 생성 후 로그 기록
    logger.LogUserActivity(user, "create_user")

    w.WriteHeader(http.StatusCreated)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [get]
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.userService.GetAllUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

package main

import (
    "log"
    "net/http"

    "github.com/IGhost-p/mini-clean-go/internal/handler"
    "github.com/IGhost-p/mini-clean-go/internal/repository"
    "github.com/IGhost-p/mini-clean-go/internal/service"
)

func main() {
    // 포트 변경
    const port = ":8081"  // 8080 -> 8081로 변경
    
    log.Println("Starting server...")
    
    // 의존성 초기화
    userRepo := repository.NewMemoryUserRepository()
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    // 라우터 설정
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            userHandler.CreateUser(w, r)
        case http.MethodGet:
            userHandler.GetUsers(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    // 서버 시작
    log.Printf("Server is running on %s\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatal(err)
    }
}
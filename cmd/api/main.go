// @title Mini Clean Go API
// @version 1.0
// @description This is a sample server for a clean architecture Go application.
// @host localhost:8081
// @BasePath /

package main

import (
    "net/http"
		
		httpSwagger "github.com/swaggo/http-swagger"  // 별칭 사용
    _ "github.com/IGhost-p/mini-clean-go/docs"

    "github.com/IGhost-p/mini-clean-go/internal/handler"
    "github.com/IGhost-p/mini-clean-go/internal/middleware"
    "github.com/IGhost-p/mini-clean-go/internal/repository"
    "github.com/IGhost-p/mini-clean-go/internal/service"
    customLogger "github.com/IGhost-p/mini-clean-go/internal/logger"
)

func main() {
    logger := customLogger.GetLogger()
    logger.Info("Starting server...")
    
    userRepo := repository.NewMemoryUserRepository()
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

		  // Swagger UI 경로 설정
			http.HandleFunc("/swagger/*", httpSwagger.Handler(
        httpSwagger.URL("http://localhost:8081/swagger/doc.json"),
    ))

    // 미들웨어를 적용한 핸들러 등록
    http.HandleFunc("/users", middleware.LoggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            userHandler.CreateUser(w, r)
        case http.MethodGet:
            userHandler.GetUsers(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }))

    const port = ":8081"
    logger.Infof("Server is running on %s", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        logger.Fatal(err)
    }
}
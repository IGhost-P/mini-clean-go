.PHONY: build run test docker-build docker-run clean

# Go 명령어
build:
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

# Docker 명령어
docker-build:
	docker build -t mini-clean-go .

docker-run:
	docker run -p 8081:8081 mini-clean-go

# Docker Compose 명령어
docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

# 청소
clean:
	rm -rf bin/
	go clean

# 개발 도구 설치
setup:
	go install github.com/golang/mock/mockgen@v1.6.0
	go install golang.org/x/tools/cmd/goimports@latest
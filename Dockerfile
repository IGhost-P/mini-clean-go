# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 소스 코드 복사
COPY . .

# 의존성 다운로드 및 빌드
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# builder 스테이지에서 빌드된 바이너리 복사
COPY --from=builder /app/bin/api .

# 포트 노출
EXPOSE 8081

# 실행
CMD ["./api"]
# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

# 의존성 다운로드를 위한 파일들만 먼저 복사
COPY go.mod .
COPY go.sum .
RUN go mod download

# 소스 코드 복사
COPY . .

# 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api cmd/api/main.go

# Final stage
FROM alpine:3.14

WORKDIR /app

# builder 스테이지에서 빌드된 바이너리 복사
COPY --from=builder /app/bin/api .

# 포트 노출
EXPOSE 8081

# 실행
CMD ["./api"]
package logger

import (
    "context" // context 패키지 추가
    "os"
    "time"

    "github.com/olivere/elastic/v7" // ElasticSearch client 라이브러리
    "github.com/sirupsen/logrus"
 		"github.com/IGhost-p/mini-clean-go/internal/model"
)

var (
    log      *logrus.Logger
    esClient *elastic.Client
)

func init() {
    log = logrus.New()
    log.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: time.RFC3339,
    })
    log.SetOutput(os.Stdout)

    // ElasticSearch 클라이언트 초기화
    initElasticSearch()
}

func initElasticSearch() {
    var err error
    esClient, err = elastic.NewClient(
        elastic.SetURL("http://elasticsearch:9200"), // 변경된 ElasticSearch URL
        elastic.SetSniff(false),
    )
    if err != nil {
        log.Fatalf("Error creating the ElasticSearch client: %s", err)
    }
}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
    return log
}

// Info logs info level messages
func Info(msg string, fields map[string]interface{}) {
    logEntry := log.WithFields(fields)
    logEntry.Info(msg)

    if esClient != nil {
        go logToElastic("info", msg, fields)
    }
}

// Error logs error level messages to both Logrus and ElasticSearch
func Error(msg string, err error, fields map[string]interface{}) {
    if fields == nil {
        fields = make(map[string]interface{})
    }
    fields["error"] = err
    logEntry := log.WithFields(fields)
    logEntry.Error(msg)

    if esClient != nil {
        go logToElastic("error", msg, fields)
    }
}

// logToElastic logs the message to ElasticSearch
func logToElastic(level string, msg string, fields map[string]interface{}) {
    if esClient == nil {
        return
    }

    logData := map[string]interface{}{
        "level":     level,
        "message":   msg,
        "timestamp": time.Now(),
        "fields":    fields,
    }

    _, err := esClient.Index().
        Index("logs").   // ElasticSearch 인덱스 이름
        BodyJson(logData).
        Do(context.Background())

    if err != nil {
        log.Errorf("Failed to log to ElasticSearch: %s", err)
    }
}

// LogUserActivity logs a user action, sending it to both Logrus and ElasticSearch
func LogUserActivity(user model.User, action string) {
    fields := map[string]interface{}{
        "user_id":   user.ID,
        "user_name": user.Name,
        "action":    action,
        "timestamp": time.Now(),
    }
    Info("User action logged", fields)
}

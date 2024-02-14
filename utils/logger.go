package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Message represents a Telegram message.
type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

type TelegramHook struct {
}

var Log = logrus.New()

func InitLogger() {
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.InfoLevel)

	path := "logs"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}

	location, _ := time.LoadLocation("Asia/Jakarta")
	d := time.Now().In(location)
	formatDate := d.Format("2006-01-02")
	logFilename := fmt.Sprintf("%s/log-%s.log", path, formatDate)

	file, _ := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	Log.SetOutput(file)
	// Log.AddHook(&TelegramHook{})
}

func (s *TelegramHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.WarnLevel}
}

func (s *TelegramHook) Fire(entry *logrus.Entry) error {
	token := os.Getenv("TELEGRAM_API_TOKEN")
	appName := os.Getenv("APP_NAME")
	chatIdStr := os.Getenv("TELEGRAM_CHAT_GROUP_ID")
	urlTelegram := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	payload, _ := json.Marshal(entry.Data)

	// <timestamp> - [loglevel] - <message> - <payload>
	text := fmt.Sprintf("%s [%s] - %s - %s - %s", appName, entry.Time.Format("2006-01-02 15:04:05"), entry.Level, entry.Message, payload)

	message := &TelegramMessage{
		ChatID: chatIdStr,
		Text:   text,
	}

	json, _ := json.Marshal(message)
	resp, err := http.Post(urlTelegram, "application/json", bytes.NewBuffer(json))

	if err != nil {
		Log.Info("Failed to send telegram", err.Error())
	}

	defer resp.Body.Close()
	return nil
}

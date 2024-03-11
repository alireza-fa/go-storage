package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	newEvent = "/api/events/raw"
)

var once sync.Once

var seqLogger *Log

type Log struct {
	ApiKey  string
	Host    string
	Port    string
	BaseUrl string
}

type EventProperty struct {
	Level           string                   `json:"Level"`
	MessageTemplate string                   `json:"MessageTemplate"`
	Timestamp       string                   `json:"Timestamp"`
	Properties      map[ExtraKey]interface{} `json:"Properties"`
}

func NewSeqLog(cfg *Config) *Log {
	once.Do(func() {
		log := &Log{
			ApiKey:  cfg.Seq.ApiKey,
			Host:    cfg.Seq.BaseUrl,
			Port:    cfg.Seq.Port,
			BaseUrl: fmt.Sprintf("http://%s:%s", cfg.Seq.BaseUrl, cfg.Seq.Port),
		}

		seqLogger = log
	})

	return seqLogger
}

func (log Log) Init() {}

func (log Log) Debug(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	go log.createNewEvent(cat, sub, DebugLevel, message, extra)
}

func (log Log) Info(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	go log.createNewEvent(cat, sub, InfoLevel, message, extra)
}

func (log Log) Warn(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	go log.createNewEvent(cat, sub, WarnLevel, message, extra)
}

func (log Log) Error(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	go log.createNewEvent(cat, sub, ErrorLevel, message, extra)
}

func (log Log) Fatal(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	go log.createNewEvent(cat, sub, FatalLevel, message, extra)
}

func (log Log) createNewEvent(cat Category, sub SubCategory, level string, message string, extra map[ExtraKey]interface{}) {
	timestamp := time.Now().Format(time.RFC3339)

	if extra != nil {
		extra = map[ExtraKey]interface{}{
			"Category":    cat,
			"SubCategory": sub,
		}
	} else {
		extra["Category"] = cat
		extra["SubCategory"] = sub
	}

	events := struct {
		Events []EventProperty
	}{
		Events: []EventProperty{
			{
				Level:           level,
				MessageTemplate: message,
				Timestamp:       timestamp,
				Properties:      extra,
			},
		},
	}

	eventsJson, err := json.Marshal(events)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(http.MethodPost, log.BaseUrl+newEvent, bytes.NewReader(eventsJson))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Api-Key", log.ApiKey)

	_, err = http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
}

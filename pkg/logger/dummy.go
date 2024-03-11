package logger

import (
	log2 "log"
)

type LogDummy struct{}

func NewDummyLog() *LogDummy {
	return &LogDummy{}
}

func (log LogDummy) Init() {}

func (log LogDummy) Debug(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	log2.Println(cat, sub, message, extra)
}

func (log LogDummy) Info(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	log2.Println(cat, sub, message, extra)
}

func (log LogDummy) Warn(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	log2.Println(cat, sub, message, extra)
}

func (log LogDummy) Error(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	log2.Println(cat, sub, message, extra)
}

func (log LogDummy) Fatal(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	log2.Println(cat, sub, message, extra)
}

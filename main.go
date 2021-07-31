package main

import (
	"log"
	"os"
)

var InfoLog *log.Logger
var WarnLog *log.Logger
var ErrLog *log.Logger

func init() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("error when open file ", err)
	}

	InfoLog = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLog = log.New(file, "WARN ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrLog = log.New(file, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
}

const (
	FlagInfo = 1 + iota
	FlagWarn
	FlagError
)

func StdLog(msg string, flag int) {
	switch flag {
	case FlagInfo:
		InfoLog.Println(msg)
	case FlagWarn:
		WarnLog.Println(msg)
	case FlagError:
		ErrLog.Println(msg)
	default:
		log.Println(msg)
	}
}

func main() {
	StdLog("something happened...", FlagInfo)
	StdLog("please be warning", FlagWarn)
	StdLog("error occurs", FlagError)
}

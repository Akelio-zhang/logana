package main

import (
	"regexp"
	"strings"
)

// const
var (
	levelPat   = regexp.MustCompile(`(DEBUG|INFO|WARN|ERROR|FATAL)`)
	classPat   = regexp.MustCompile(`[a-zA-z.]* -`)
	exptionPat = regexp.MustCompile(`[a-zA-z.]*Exception`)
)

type ParseHandler struct {
	SerialNum uint32

	LineNum  int
	FileName string
	rawData  string

	logInfoSlice []LogInfo
}

func (handler *ParseHandler) addLine(line string) {
	if line != "" {
		handler.rawData += line + "\n"
	}
}

func (handler *ParseHandler) parse() {
	str := handler.rawData
	if str == "" {
		return
	}
	time := str[0:23]

	level := strings.Replace(levelPat.FindString(str), " ", "", -1)
	class := classPat.FindString(str)
	class = class[:len(class)-2]
	exception := exptionPat.FindString(str)
	logInfo := LogInfo{Time: time, Level: level, Class: class, Exception: exception}
	handler.logInfoSlice = append(handler.logInfoSlice, logInfo)
}

func (handler *ParseHandler) clear() {
	handler.SerialNum++
	handler.rawData = ""
}

package main

import (
	"bufio"
	"flag"
	"os"
	"path/filepath"
	"strings"
)

type LogInfo struct {
	Time      string
	Level     string
	Class     string
	Msg       string
	Exception string
}

func main() {
	dir := flag.String("dir", ".", "dir to scan")
	flag.Parse()

	filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".log") {
			execute(path)
		}
		return nil
	})
}

func execute(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()

	in := bufio.NewScanner(f)
	handler := &ParseHandler{SerialNum: 0, logInfoSlice: make([]LogInfo, 0)}

	for in.Scan() {
		// 1. read line
		line := in.Text()
		// 2. if encounter a new line, parse line and clear handler.
		if strings.HasPrefix(line, "20") {
			handler.parse()
			handler.clear()
		}
		// 3. add line to handler.
		handler.addLine(line)
	}
	handler.parse()

	output(handler.logInfoSlice)
}

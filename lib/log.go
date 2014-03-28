package lib

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func init() {
	fmt.Println("Initial log model")
}

const (
	LOG_DEBUG = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR

	_formatter = `[%s] [%s] %s:%d - %s`
)

var (
	_level_str = []string{"DEBUG", "INFO ", "WARN ", "ERROR"}
)

type Log struct {
	Level int
	Inner *log.Logger
	file  *os.File
}

func NewLog(fs string, lv int) (*Log, error) {
	f, err := os.OpenFile(fs, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return nil, err
	}
	l := log.New(f, "", 0644)
	return &Log{Level: lv, Inner: l, file: f}, nil
}

func (self *Log) Close() {
	self.file.Close()
}

func (self *Log) Debug(msg string, a ...interface{}) {
	if self.Level >= LOG_DEBUG {
		self._write(fmt.Sprintf(msg, a...), LOG_DEBUG)
	}
}

func (self *Log) Info(msg string, a ...interface{}) {
	if self.Level >= LOG_DEBUG {
		self._write(fmt.Sprintf(msg, a...), LOG_INFO)
	}
}

func (self *Log) Warn(msg string, a ...interface{}) {
	if self.Level >= LOG_DEBUG {
		self._write(fmt.Sprintf(msg, a...), LOG_WARN)
	}
}

func (self *Log) Error(msg string, a ...interface{}) {
	if self.Level >= LOG_DEBUG {
		self._write(fmt.Sprintf(msg, a...), LOG_ERROR)
	}
}

func (self *Log) _write(msg string, lv int) {
	var (
		file string
		line int
		ok   bool
	)
	_, file, line, ok = runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	t := time.Now().Format(time.RubyDate)
	self.Inner.Println(fmt.Sprintf(_formatter, _level_str[lv], t, file, line, msg))
}

package logger

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"service-all/library/file"
	"time"
)

const (
	// LevelError 错误
	LevelError = iota
	// LevelWarning 警告
	LevelWarning
	// LevelInformational 提示
	LevelInformational
	// LevelDebug 除错
	LevelDebug
)

var DefaultCallerDepth = 3
var GloablLogger *Logger
var Level = LevelDebug
var F *os.File
var DefaultPrefix = ""
var logger *log.Logger

// Logger 日志
type Logger struct {
	level int
}

// 日志颜色
var colors = map[string]func(a ...interface{}) string{
	"Warning": color.New(color.FgYellow).Add(color.Bold).SprintFunc(),
	"Panic":   color.New(color.BgRed).Add(color.Bold).SprintFunc(),
	"Error":   color.New(color.FgRed).Add(color.Bold).SprintFunc(),
	"Info":    color.New(color.FgCyan).Add(color.Bold).SprintFunc(),
	"Debug":   color.New(color.FgWhite).Add(color.Bold).SprintFunc(),
}

// 不同级别前缀与时间的间隔，保持宽度一致
var spaces = map[string]string{
	"Warning": "",
	"Panic":   "  ",
	"Error":   "  ",
	"Info":    "   ",
	"Debug":   "  ",
}

func Init() {
	filePath := fmt.Sprintf("%s", "logs/")
	fileName := fmt.Sprintf("%s.%s",
		time.Now().Format("20060102"), "log",
	)
	F, err := file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Println 打印
func (ll *Logger) Println(prefix string, msg string) {
	// TODO Release时去掉
	// color.NoColor = false
	_, file, line, _ := runtime.Caller(DefaultCallerDepth)
	//log.Printf("test file %s line %d", filepath.Base(file), line)
	c := color.New()
	_, _ = c.Printf(
		"%s%s [%s:%d] %s %s\n",
		colors[prefix]("["+prefix+"]"),
		spaces[prefix],
		filepath.Base(file),
		line,
		time.Now().Format("2006-01-02 15:04:05"),
		msg,
	)
	logger.Println(msg)
}

// Panic 极端错误
func (ll *Logger) Panic(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	ll.Println("Panic", msg)
	panic(msg)
}

// Error 错误
func (ll *Logger) Error(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	ll.Println("Error", msg)
}

// Warning 警告
func (ll *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > ll.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	ll.Println("Warning", msg)
}

// Info 信息
func (ll *Logger) Info(format string, v ...interface{}) {
	if LevelInformational > ll.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	ll.Println("Info", msg)
}

// Debug 校验
func (ll *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > ll.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	ll.Println("Debug", msg)
}

//Print GORM 的 Logger实现
func (ll *Logger) Print(v ...interface{}) {
	if LevelDebug > ll.level {
		return
	}
	msg := fmt.Sprintf("[SQL] %s", v...)
	ll.Println("SQL", msg)
}

// BuildLogger 构建logger
func BuildLogger(level string) {
	intLevel := LevelError
	switch level {
	case "error":
		intLevel = LevelError
	case "warning":
		intLevel = LevelWarning
	case "info":
		intLevel = LevelInformational
	case "debug":
		intLevel = LevelDebug
	}
	l := Logger{
		level: intLevel,
	}
	GloablLogger = &l
}

// Log 返回日志对象
func Log() *Logger {
	if GloablLogger == nil {
		l := Logger{
			level: Level,
		}
		GloablLogger = &l
	}
	return GloablLogger
}
func RetuLog() Logger {
	if GloablLogger == nil {
		l := Logger{
			level: Level,
		}
		GloablLogger = &l
	}
	return *GloablLogger
}

package gol

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger

	logLevel = 0
	logOut   *os.File
	day      int
	logFile  string
)

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func SetLogLevel(level int) {
	logLevel = level
}

func SetLogFile(file string) error {
	logFile = file
	now := time.Now()
	var err error
	if logOut, err = os.OpenFile(logFile+now.Format("06-01-02.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664); err != nil {
		return err
	} else {
		debugLogger = log.New(logOut, "DEBUG ", log.Ltime)
		infoLogger = log.New(logOut, "INFO ", log.Ltime)
		warnLogger = log.New(logOut, "WARN ", log.Ltime)
		errorLogger = log.New(logOut, "ERROR ", log.Ltime)
		day = now.YearDay()
	}

	return nil
}

func switchLogfile() {
	now := time.Now()
	if now.YearDay() == day {
		return
	}

	logOut.Close()
	var err error
	if logOut, err = os.OpenFile(logFile+now.Format("2006-01-02"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664); err != nil {
		fmt.Printf("create log file %s failed %v\n", logFile, err)
		return
	} else {
		debugLogger.SetOutput(logOut)
		infoLogger.SetOutput(logOut)
		warnLogger.SetOutput(logOut)
		errorLogger.SetOutput(logOut)
		day = now.YearDay()
	}
}

func addPrefix() string {
	file, line := getLineNo()
	arr := strings.Split(file, "/")
	if len(arr) > 3 {
		arr = arr[len(arr)-3:]
	}
	return strings.Join(arr, "/") + ":" + strconv.Itoa(line) + " "
}

func Debugf(format string, v ...any) {
	if logLevel <= DebugLevel {
		switchLogfile()
		debugLogger.Printf(addPrefix()+format, v...)
	}
}

func Infof(format string, v ...any) {
	if logLevel <= InfoLevel {
		switchLogfile()
		infoLogger.Printf(addPrefix()+format, v...)
	}
}

func Warnf(format string, v ...any) {
	if logLevel <= WarnLevel {
		switchLogfile()
		warnLogger.Printf(addPrefix()+format, v...)
	}
}

func Errorf(format string, v ...any) {
	if logLevel <= ErrorLevel {
		switchLogfile()
		errorLogger.Printf(addPrefix()+format, v...)
	}
}

func Debugln(v ...any) {
	if logLevel <= DebugLevel {
		switchLogfile()
		debugLogger.Println(addPrefix() + fmt.Sprint(v...))
	}
}

func Infoln(v ...any) {
	if logLevel <= InfoLevel {
		switchLogfile()
		infoLogger.Println(addPrefix() + fmt.Sprint(v...))
	}
}

func Warnln(v ...any) {
	if logLevel <= WarnLevel {
		switchLogfile()
		warnLogger.Println(addPrefix() + fmt.Sprint(v...))
	}
}

func Errorln(v ...any) {
	if logLevel <= ErrorLevel {
		switchLogfile()
		errorLogger.Println(addPrefix() + fmt.Sprint(v...))
	}
}

func Debug(v ...any) {
	if logLevel <= DebugLevel {
		switchLogfile()
		debugLogger.Print(addPrefix() + fmt.Sprint(v...))
	}
}

func Info(v ...any) {
	if logLevel <= InfoLevel {
		switchLogfile()
		infoLogger.Print(addPrefix() + fmt.Sprint(v...))
	}
}

func Warn(v ...any) {
	if logLevel <= WarnLevel {
		switchLogfile()
		warnLogger.Print(addPrefix() + fmt.Sprint(v...))
	}
}

func Error(v ...any) {
	if logLevel <= ErrorLevel {
		switchLogfile()
		errorLogger.Print(addPrefix() + fmt.Sprint(v...))
	}
}

func getLineNo() (string, int) {
	_, file, line, ok := runtime.Caller(3)
	if ok {
		return file, line
	} else {
		return "", 0
	}
}

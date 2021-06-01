package log

import (
	"io"
	"log"
	"os"
)

// WriteLog 将日志内容 text 写入到 file 中
func WriteLog(file string, text string, prefix string) error {
	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	logger := getLogger(logFile, prefix, log.Llongfile|log.Lmicroseconds|log.Ldate)
	logger.Println(text)
	return nil
}

func getLogger(out io.Writer, prefix string, flag int) *log.Logger {
	return log.New(out, prefix, flag)
}

package log

import (
	"github.com/Node1650665999/Glib/common"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// WriteLog 将日志内容 text 写入到 file 中
func WriteLog(text interface{})  error {
	pc, file, line, _ := runtime.Caller(1)

	//项目根目录
	wd, _  := os.Getwd()

	//文件名
	funcname := filepath.Ext(runtime.FuncForPC(pc).Name())
	funcname = strings.TrimPrefix(funcname, ".")
	filename := fmt.Sprintf("%s-%s.log", funcname, time.Now().Format("2006-01-02"))

	//日志路径
	dirSplit := strings.Split(file, wd)
	subpath  := strings.Replace(dirSplit[len(dirSplit)-1], ".go", "", -1)
	logpath  := fmt.Sprintf("%s/runtime/logs/%s", wd, strings.Trim(subpath, "/"))
	if err   := common.MkDir(logpath); err != nil {
		return err
	}

	//完整的日志文件路径
	save := fmt.Sprintf("%s/%s", logpath, filename)
	fh, err := os.OpenFile(save, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer fh.Close()
	if err != nil {
		return err
	}

	//写入日志
	prefix := fmt.Sprintf("[info:%d]", line)
	logger :=  log.New(fh, prefix,  log.Lmicroseconds|log.Ldate)
	logger.Println(text)

	return nil
}

/*func WriteLog2(file string, text string, prefix string) error {
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
}*/

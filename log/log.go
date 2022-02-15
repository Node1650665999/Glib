package log

import (
	"fmt"
	"github.com/node1650665999/Glib/common"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// WriteLog 将日志内容 text 写入到 file 中
func WriteLog(text interface{}) error {
	pc, file, line, _ := runtime.Caller(1)

	//文件名
	funcname := filepath.Ext(runtime.FuncForPC(pc).Name())
	funcname = strings.TrimPrefix(funcname, ".")

	//日志路径
	wd, _   := os.Getwd()
	subpath := ""

	r, _ := regexp.Compile("@v.+")
	if r.MatchString(file) {
		//应用外调用
		match := r.FindString(file)
		index   := strings.Index(match, string(os.PathSeparator))
		subpath = match[index:]
	} else {
		//应用内调用
		/*dirSplit := strings.SplitAfter(file, wd)
		subpath  = dirSplit[len(dirSplit)-1]*/
		subpath  = strings.Replace(file, wd, "", -1)
	}

	prefix   := strings.Replace(path.Base(subpath),path.Ext(subpath),"", -1)
	filename := fmt.Sprintf("%s-%s.log", prefix, time.Now().Format("2006-01-02"))

	//创建日志目录
	logpath  := fmt.Sprintf("%s/runtime/logs/%s", wd, strings.Trim(subpath, "/"))
	if err   := common.MkDir(logpath); err != nil {
		return err
	}

	//完整的日志文件路径
	save := fmt.Sprintf("%s/%s", logpath, filename)

	//超过容量限制则重命名
	fileInfo,err := os.Stat(save)
	logSize  := int64(64) * 1024 * 1024
	if err == nil && fileInfo.Size() > logSize {
		renameSave := fmt.Sprintf("%s-%d", save, time.Now().Nanosecond())
		err := os.Rename(save, renameSave)
		if err != nil {
			return err
		}
	}

	fh, err := os.OpenFile(save, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer fh.Close()
	if err != nil {
		return err
	}

	//写入日志
	pfx := fmt.Sprintf("[info:%s:%d]",funcname,line)
	logger :=  log.New(fh, pfx,  log.Lmicroseconds|log.Ldate)
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


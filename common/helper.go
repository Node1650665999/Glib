package common

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

//InArray 判断切片是否包含某个元素
func InArray(need interface{}, haystack interface{}) bool {
	switch key := need.(type) {
	case int:
		for _, item := range haystack.([]int) {
			if item == key {
				return true
			}
		}
	case string:
		for _, item := range haystack.([]string) {
			if item == key {
				return true
			}
		}
	case int64:
		for _, item := range haystack.([]int64) {
			if item == key {
				return true
			}
		}
	case float64:
		for _, item := range haystack.([]float64) {
			if item == key {
				return true
			}
		}
	default:
		return false
	}
	return false
}

//ByteFormat 将字节格式化为指定的单位
// refer:https://blog.microdba.com/golang/2016/05/01/golang-byte-conv/
func ByteFormat(size uint64) string {
	sz   := float64(size)
	base := float64(1024)
	unit := []string{"B","KB","MB","GB","TB","EB"}
	i := 0

	for sz >= base {
		sz /= base
		i++
	}
	return fmt.Sprintf("%.2f%s",sz, unit[i])
}


//MkDir 创建一个目录
func MkDir(path string) error {
	_, err := os.Stat(path)

	// 权限问题
	if os.IsPermission(err) {
		return fmt.Errorf(" Permission denied src: %s", err)
	}

	// 已存在
	if os.IsExist(err) {
		return nil
	}

	// 创建目录
	return os.MkdirAll(path, os.ModePerm)
}

// Ext get the file ext
func Ext(fileName string) string {
	return path.Ext(fileName)
}

// BaseName 获取文件的 basename
func BaseName(filename string) string  {
	return path.Base(filename)
}

//FileExist 判断文件是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//StrJoin 用来拼接字符串
func StrJoin(args ...string) string {
	var buf bytes.Buffer
	for _, arg := range args {
		buf.WriteString(arg)
	}
	return buf.String()
}

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

//NumberUuid 生成数字类型的uuid
func NumberUuid() int64 {
	w := &Worker{
		timestamp: 0,
		workerId:  1,
		number:    0,
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}

	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}

//StrUuid 生成字符串格式的UUID
func StrUuid(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

//ExecCommand 运行系统命令和二进制文件
func ExecCommand(cmd string, params ...string) (string, error) {
	// Print Go Version
	cmdOutput, err := exec.Command(cmd, params...).Output()
	if err != nil {
		return "", err
	}
	return string(cmdOutput), err
}

//JsonEncode 实现Json编码
func JsonEncode(v interface{}) string {
	u,_ := json.Marshal(v)
	return string(u)
}

//JsonDecode 实现json解码
func JsonDecode(data string, v interface{}) error  {
	return json.Unmarshal([]byte(data), v)
}

//GetBuildAbsPath 获取编译后可执行文件的根目录
//注意： 如果以 go run 运行，则无法获取正确的根目录,因为 go run 生成的可执行文件位于/tmp目录下
func GetBuildAbsPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}

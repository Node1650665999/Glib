package io

import (
	"github.com/Node1650665999/Glib/common"
	mytime "github.com/Node1650665999/Glib/time"
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ReadFile 读取文件并将内容写入到一个 string 切片中的
func ReadFile(file string) ([]string, error) {
	inputFile, err := os.Open(file)
	defer inputFile.Close()

	if err != nil {
		return nil, err
	}

	inputReader := bufio.NewReader(inputFile)
	container := []string{}
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF || readerError != nil {
			break
		}
		container = append(container, inputString)
	}

	return container, nil
}

// WriteFile 往 file中写入 content
func WriteFile(file string, content string) (int, error) {
	outputFile, _ := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	n, err := outputWriter.WriteString(content)
	if err != nil {
		return n, err
	}

	err = outputWriter.Flush()
	return n, err
}

// WriteFileSimple 以非缓冲的方式写入文件
func WriteFileSimple(file string, content string) (int, error) {
	outputFile, _ := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	defer outputFile.Close()
	return outputFile.WriteString(content)
}

// ReadFileSimple 读取文件
func ReadFileSimple(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

// CopyFile 将文件 srcName 复制到 dstName
func CopyFile(dstName string, srcName string) (int64, error) {
	src, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

type FileUpload struct {
	AcceptType []string
	maxSize    int64
	fromFile   string
	savePath   string
	rename     bool
}

//NewFileUpload 构建文件上传对象
func NewFileUpload(fromFile string, savePath string) *FileUpload {
	return &FileUpload{
		AcceptType: []string{"image/png", "image/jpg", "image/jpeg", "image/gif", "image/webp"},
		maxSize:    20 * 1024 * 1024,
		fromFile:   fromFile,
		savePath:   savePath,
		rename:     true,
	}
}

// UploadHandler 处理文件上传
func (upload *FileUpload) UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) (dstfile string, err error) {
	r.Body = http.MaxBytesReader(w, r.Body, upload.maxSize)

	//文件大小限制
	if err := r.ParseMultipartForm(upload.maxSize); err != nil {
		return "", err
	}

	//读取文件
	file, headers, err := r.FormFile(upload.fromFile)
	if err != nil {
		return "", err
	}

	//上传类型校验
	fileType := strings.ToLower(headers.Header.Get("Content-Type"))
	if !common.InArray(fileType, upload.AcceptType) {
		return "", fmt.Errorf("%s 类型的文件不被允许上传", fileType)
	}

	//读取文件内容
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	ext := common.Ext(headers.Filename)

	//写入文件
	dstFile := ""
	if upload.rename {
		dstFile = upload.savePath + "/" + strconv.FormatInt(mytime.CurrentTimestamp(), 10) + "." + ext
	} else {
		dstFile = upload.savePath + "/" + headers.Filename
	}

	err = ioutil.WriteFile(dstFile, data, 0666)
	if err != nil {
		return "", err
	}
	return dstfile, err
}

//RemoteDownload 实现下载文件到本地,获得网络文件的输入流以及本地文件的输出流 ,然后将输入流读取到输出流中
func RemoteDownload(remote string,local string) error  {
	res, err := http.Get(remote)
	if err != nil {
		return fmt.Errorf("A error occurred: %v", err)
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32 * 1024)

	// 获得文件的writer对象
	file, err := os.Create(local)
	defer file.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}

	return nil
}

func ImportExcel() {

}

func ExportExcel()  {
}

func ImportCsv()  {

}

//ExportCsv 导出 csv 文件
func ExportCsv(filename string, data [][]string, w http.ResponseWriter) error {
	buf := &bytes.Buffer{}
	buf.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM,避免文件打开乱码

	writer := csv.NewWriter(buf)
	writer.WriteAll(data)

	// 设置下载的文件名
	w.Header().Add("Content-Type", "application/octet-stream")
	//w.Header().Add("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", filename))
	//输出数据
	w.Write(buf.Bytes())
	return nil
}

//SaveCsv 保存生成的csv文件到本地
func SaveCsv(filename string, data [][]string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	/*data := [][]string{
		{"1", "test1", "test1-1"},
		{"2", "test2", "test2-1"},
		{"3", "test3", "test3-1"},
	}*/

	return w.WriteAll(data)
}



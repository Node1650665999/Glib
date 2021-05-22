package io

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

// ReadFile 读取文件并将内容写入到一个 string 切片中的
func ReadFile (file string) ([]string,error) {
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
func WriteFile(file string, content string) (int, error)  {
	outputFile, _ := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)

	defer outputFile.Close()

	outputWriter  := bufio.NewWriter(outputFile)

	n, err := outputWriter.WriteString(content)
	if err != nil {
		return n, err
	}

	err = outputWriter.Flush()
	return  n, err
}

// WriteFileSimple 以非缓冲的方式写入文件
func WriteFileSimple(file string, content string) (int, error)  {
	outputFile, _ := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	defer outputFile.Close()
	return  outputFile.WriteString(content)
}

// ReadFileSimple 读取文件
func ReadFileSimple(file string) ([]byte, error)  {
	return  ioutil.ReadFile(file)
}

// CopyFile 将文件 srcName 复制到 dstName
func CopyFile( dstName string, srcName string) (int64, error)  {
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
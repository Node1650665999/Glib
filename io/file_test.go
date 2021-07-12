package io_test

import (
	"github.com/node1650665999/Glib/io"
	"net/http"
	"testing"
)

func TestExportCsv(t *testing.T) {
	http.HandleFunc("/export_csv", func(writer http.ResponseWriter, request *http.Request) {
		data := [][]string{
			{"编号", "姓名", "年龄"},
			{"1", "张三", "23"},
			{"2", "李四", "24"},
			{"3", "王五", "25"},
			{"4", "赵六", "26"},
		}
		filename := "data.csv"
		save     := "../data/" + filename
		io.SaveCsv(save, data)
		io.ExportCsv("data.csv", data , writer)
	})
	http.ListenAndServe(":9090", nil)
	t.Log("export finish !!!")
}

func TestImportCsv(t *testing.T) {
	http.HandleFunc("/import_csv", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseMultipartForm(32 << 20)
		file, handler,_ := request.FormFile("file")
		data,_ := io.ImportCsv(file)
		t.Log(data)
		t.Log(handler.Header)
		writer.Write([]byte("success"))
	})
	http.ListenAndServe(":9090", nil)
}

func TestExportExcel(t *testing.T) {
	data := [][]string{
		{"编号", "姓名", "年龄"},
		{"1", "张三", "23"},
		{"2", "李四", "24"},
		{"3", "王五", "25"},
		{"4", "赵六", "26"},
	}

	http.HandleFunc("/export_excel", func(writer http.ResponseWriter, request *http.Request) {
		filename := "data.xlsx"
		save     := "../data/" + filename
		io.SaveExcel(save, data)
		io.ExportExcel(filename, data, writer)
	})
	http.ListenAndServe(":9090", nil)
}

func TestImportExcel(t *testing.T) {
	http.HandleFunc("/import_excel", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseMultipartForm(32 << 20)
		file, handler,_ := request.FormFile("file")
		data,_ := io.ImportExcel(file)
		t.Log(data)
		t.Log(handler.Header)
		writer.Write([]byte("success"))
	})
	http.ListenAndServe(":9090", nil)
}

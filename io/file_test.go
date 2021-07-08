package io_test

import (
	"Glib/io"
	"testing"
)

func TestExportCsv(t *testing.T) {

	data := [][]string{
		{"1", "test1", "test1-1"},
		{"2", "test2", "test2-1"},
		{"3", "test3", "test3-1"},
	}

	io.ExportCsv("test.csv", data, nil)
}

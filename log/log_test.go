package log_test

import "testing"
import "Glib/log"

func TestWriteLog(t *testing.T) {
	data := "asd"
	err := log.WriteLog(data)
	if err != nil {
		t.Fatalf("some thing error:%s ", err)
	}
	t.Log("success")
}
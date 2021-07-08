package common_test

import (
	"github.com/Node1650665999/Glib/common"
	"testing"
)

func TestJwtExpire(t *testing.T) {
	//token := common.GenerateToken(3)
	token := "eyJhbGciOiJtZDUiLCJ0eXAiOiJqd3QifQ==.eyJpc3MiOiIiLCJzdWIiOiIiLCJhdWQiOiIiLCJleHAiOjE2MjM3N\nTk4MTAsIm5iZiI6MCwiaWF0IjoxNjIzNzU5ODA3LCJqdGkiOiJmMGE0NGJiNzEyZWNkMWYwZDFjZDRiYjg5Njg4MzU\nifQ==.5445479584556afc49d1aa52ed3f68e5"
	if err := common.VerifyToken(token); err != nil {
		t.Log(err.Error())
	} else {
		t.Log("pass")
	}
}


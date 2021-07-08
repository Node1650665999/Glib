package common_test

import "testing"
import "github.com/Node1650665999/Glib/common"

func TestExecCommand(t *testing.T) {
	res, err := common.ExecCommand("ps", "-aux")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}



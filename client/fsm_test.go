package ztpclient

import (
	"testing"
	"time"
)

func TestFSM(t *testing.T) {
    DEBUG = true
    for _,ID := range []string{"11111-11111","22222-22222", "33333-333333"} {
        zc := NewZtpClient(ID)
        go zc.StateMachine()
    }
    time.Sleep(time.Second * 10)
}

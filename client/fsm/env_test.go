package ztpclient

import (
	"os"
	"testing"
)

func TestDebugBool(t *testing.T) {
	if DEBUG {
		t.Fatal("DEBUG should default to false")
	}
	os.Setenv(EnvDebug, "1")
	envSetup()
	if !DEBUG {
		t.Fatal("DEBUG was not enabled when setting environment", EnvDebug)
	}
	DEBUG = true
	os.Setenv(EnvDebug, "0")
	envSetup()
	if DEBUG {
		t.Fatal("DEBUG was not disabled when setting environment", EnvDebug)
	}
}

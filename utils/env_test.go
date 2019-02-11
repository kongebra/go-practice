package utils

import (
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	pEnv := os.Getenv("PORT")

	var expectedPort string

	if pEnv == "" {
		expectedPort = ":8080"
	} else {
		expectedPort = pEnv
	}

	port := GetPort()

	if port == "" {
		t.Logf("\nexpected: %s\ngot: %s\n", expectedPort, port)
		t.Fail()
	}

	if port != expectedPort {
		t.Logf("\nexpected: %s\ngot: %s\n", expectedPort, port)
		t.Fail()
	}
}

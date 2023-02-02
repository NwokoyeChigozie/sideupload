package tests

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/vesicash/upload-ms/internal/config"
	"github.com/vesicash/upload-ms/pkg/repository/storage/awss3"
	"github.com/vesicash/upload-ms/utility"
)

func Setup() *utility.Logger {
	logger := utility.NewLogger()
	config.Setup(logger, "../../app")
	awss3.ConnectAws(logger)
	return logger
}

func ParseResponse(w *httptest.ResponseRecorder) map[string]interface{} {
	res := make(map[string]interface{})
	json.NewDecoder(w.Body).Decode(&res)
	return res
}

func AssertStatusCode(t *testing.T, got, expected int) {
	if got != expected {
		t.Errorf("handler returned wrong status code: got status %d expected status %d", got, expected)
	}
}

func AssertBool(t *testing.T, got, expected bool) {
	if got != expected {
		t.Errorf("handler returned wrong boolean: got %v expected %v", got, expected)
	}
}

func AssertResponseMessage(t *testing.T, got, expected string) {
	if got != expected {
		t.Errorf("handler returned wrong message: got message: %q expected: %q", got, expected)
	}
}

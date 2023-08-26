package logging_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/afiskon/promtail-client/promtail"
	"logging"
)

func TestInitLogger_Success(t *testing.T) {
	// Prepare test data
	format := "proto"
	labels := "test_labels"
	config := &promtail.ClientConfig{
		PushURL:            "http://localhost:3100/api/prom/push",
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}

	// Call the function
	err := logging.InitLogger(format, labels, config)

	// Check the result
	if err != nil {
		t.Errorf("InitLogger returned an unexpected error: %s", err)
	}
}

func TestInitLogger_InvalidFormat(t *testing.T) {
	// Prepare test data
	format := "invalid"
	labels := "test_labels"
	config := &promtail.ClientConfig{
		PushURL:            "http://localhost:3100/api/prom/push",
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}

	// Call the function
	err := logging.InitLogger(format, labels, config)

	// Check the result
	if err == nil {
		t.Error("InitLogger expected to return an error for invalid log format, but it didn't")
	}
	// Check the error message
	expectedErrMsg := fmt.Sprintf("invalid log format: %s", format)
	if err.Error() != expectedErrMsg {
		t.Errorf("InitLogger returned an unexpected error message.\nExpected: %s\nActual: %s", expectedErrMsg, err.Error())
	}
}

func TestNewLogger_Success(t *testing.T) {
	// Prepare test data
	format := "proto"
	labels := "test_labels"
	config := &promtail.ClientConfig{
		PushURL:            "http://localhost:3100/api/prom/push",
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}

	// Call the function
	loggerInstance, err := logging.NewLogger(format, labels, config)

	// Check the result
	if err != nil {
		t.Errorf("NewLogger returned an unexpected error: %s", err)
	}
	if loggerInstance == nil {
		t.Error("NewLogger returned a nil logger instance")
	}
	// TODO: Add additional checks for the loggerInstance if necessary
}

func TestNewLogger_DefaultConfig(t *testing.T) {
	// Prepare test data
	format := "json"
	labels := "test_labels"

	// Call the function without providing the config
	loggerInstance, err := logging.NewLogger(format, labels, nil)

	// Check the result
	if err != nil {
		t.Errorf("NewLogger returned an unexpected error: %s", err)
	}
	if loggerInstance == nil {
		t.Error("NewLogger returned a nil logger instance")
	}
	// TODO: Add additional checks for the loggerInstance and default config if necessary
}

// TODO: Add more unit tests for other functions in the logger package

func TestDebugf(t *testing.T) {
	// Prepare test data
	format := "proto"
	labels := "test_labels"
	config := &promtail.ClientConfig{
		PushURL:            "http://localhost:3100/api/prom/push",
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}
	_ = logging.InitLogger(format, labels, config)

	// Call the function
	logging.Debugf("Test debug message")

	// TODO: Add assertions or checks based on your logging implementation
}

func TestInfof(t *testing.T) {
	// Prepare test data
	format := "proto"
	labels := "test_labels"
	config := &promtail.ClientConfig{
		PushURL:            "http://localhost:3100/api/prom/push",
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}
	_ = logging.InitLogger(format, labels, config)

	// Call the function
	logging.Infof("Test info message")

	// TODO: Add assertions or checks based on your logging implementation
}

func TestWarnf(t *testing.T) {
	// Prepare test data
	format := "proto"
	labels := "test_labels"
	config := &promtail.ClientConfig{
		PushURL:            "http://localhost:3100/api/prom/push",
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}
	_ = logging.InitLogger(format, labels, config)

	// Call the function
	logging.Warnf("Test warning message")

	// TODO: Add assertions or checks based on your logging implementation
}

func TestErrorf(t *testing.T) {
	// Prepare test data
	format := "proto"
	labels := "test_labels"
	config := &promtail.ClientConfig{
		PushURL:            "http://localhost:3100/api/prom/push",
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}
	_ = logging.InitLogger(format, labels, config)

	// Call the function
	logging.Errorf("Test error message")

	// TODO: Add assertions or checks based on your logging implementation
}

func TestShutdown(t *testing.T) {
	// Prepare test data
	format := "proto"
	labels := "test_labels"
	config := &promtail.ClientConfig{
		PushURL:            "http://localhost:3100/api/prom/push",
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}
	_ = logging.InitLogger(format, labels, config)

	// Call the function
	logging.Shutdown()

	// TODO: Add assertions or checks based on your logging implementation
}

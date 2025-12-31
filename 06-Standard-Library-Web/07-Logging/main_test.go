package main

import (
	"bytes"
	"log"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestBasicLogging(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stdout)
	
	log.Println("Test message")
	
	if !strings.Contains(buf.String(), "Test message") {
		t.Error("Log message should be in output")
	}
}

func TestStructuredLogging(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, nil))
	
	logger.Info("Test message", "key", "value")
	
	output := buf.String()
	if !strings.Contains(output, "Test message") {
		t.Error("Log message should be in output")
	}
	if !strings.Contains(output, "key=value") {
		t.Error("Structured fields should be in output")
	}
}

func TestJSONLogging(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	
	logger.Info("Test message", "key", "value")
	
	output := buf.String()
	if !strings.Contains(output, "Test message") {
		t.Error("Log message should be in output")
	}
	if !strings.Contains(output, "\"key\"") {
		t.Error("JSON structure should be present")
	}
}

func TestLoggerWithContext(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, nil))
	logger = logger.With("request_id", "req-123")
	
	logger.Info("Message")
	
	output := buf.String()
	if !strings.Contains(output, "request_id=req-123") {
		t.Error("Context should be included in log")
	}
}


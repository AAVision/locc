package main

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer
	logger := NewLogger(LogLevelDebug, &out, &errOut)

	t.Run("Debug", func(t *testing.T) {
		out.Reset()
		logger.Debug("test %s", "debug")
		if !strings.Contains(out.String(), "[DEBUG] test debug") {
			t.Errorf("Expected log to contain '[DEBUG] test debug', got %q", out.String())
		}
	})

	t.Run("Info", func(t *testing.T) {
		out.Reset()
		logger.Info("test %s", "info")
		if !strings.Contains(out.String(), "[INFO] test info") {
			t.Errorf("Expected log to contain '[INFO] test info', got %q", out.String())
		}
	})

	t.Run("Warn", func(t *testing.T) {
		out.Reset()
		logger.Warn("test %s", "warn")
		if !strings.Contains(out.String(), "[WARN] test warn") {
			t.Errorf("Expected log to contain '[WARN] test warn', got %q", out.String())
		}
		if logger.GetWarnCount() != 1 {
			t.Errorf("Expected warn count 1, got %d", logger.GetWarnCount())
		}
	})

	t.Run("Error", func(t *testing.T) {
		errOut.Reset()
		logger.Error("test %s", "error")
		if !strings.Contains(errOut.String(), "[ERROR] test error") {
			t.Errorf("Expected log to contain '[ERROR] test error', got %q", errOut.String())
		}
		if logger.GetErrorCount() != 1 {
			t.Errorf("Expected error count 1, got %d", logger.GetErrorCount())
		}
	})

	t.Run("Level filtering", func(t *testing.T) {
		logger.SetLevel(LogLevelInfo)
		out.Reset()
		logger.Debug("should not appear")
		if out.Len() > 0 {
			t.Errorf("Debug log should have been filtered out, got %q", out.String())
		}

		logger.SetLevel(LogLevelSilent)
		out.Reset()
		logger.Error("should not appear")
		if errOut.Len() > 0 && strings.Contains(errOut.String(), "should not appear") {
			t.Errorf("Error log should have been filtered out when silent")
		}
	})
}

func TestCustomErrors(t *testing.T) {
	t.Run("FileError", func(t *testing.T) {
		err := NewFileError("file.go", errors.New("read error"))
		expected := "error processing file file.go: read error"
		if err.Error() != expected {
			t.Errorf("FileError.Error() = %q, want %q", err.Error(), expected)
		}
	})

	t.Run("DirectoryError", func(t *testing.T) {
		err := NewDirectoryError("dir", errors.New("access error"))
		expected := "error processing directory dir: access error"
		if err.Error() != expected {
			t.Errorf("DirectoryError.Error() = %q, want %q", err.Error(), expected)
		}
	})

	t.Run("PermissionError", func(t *testing.T) {
		err := NewPermissionError("secret", errors.New("denied"))
		expected := "permission denied: secret"
		if err.Error() != expected {
			t.Errorf("PermissionError.Error() = %q, want %q", err.Error(), expected)
		}
	})
}

func TestGlobalLogFunctions(t *testing.T) {
	SetLogLevel(LogLevelDebug)
	var out bytes.Buffer
	var errOut bytes.Buffer

	SetLogOutput(&out)
	SetLogErrorOutput(&errOut)

	// Reset to original outputs after test
	defer func() {
		SetLogOutput(os.Stdout)
		SetLogErrorOutput(os.Stderr)
	}()

	t.Run("LogDebug", func(t *testing.T) {
		out.Reset()
		LogDebug("global debug")
		if !strings.Contains(out.String(), "[DEBUG] global debug") {
			t.Errorf("Expected output to contain global debug, got %q", out.String())
		}
	})

	t.Run("LogInfo", func(t *testing.T) {
		out.Reset()
		LogInfo("global info")
		if !strings.Contains(out.String(), "[INFO] global info") {
			t.Errorf("Expected output to contain global info, got %q", out.String())
		}
	})

	t.Run("LogWarn", func(t *testing.T) {
		out.Reset()
		LogWarn("global warn")
		if !strings.Contains(out.String(), "[WARN] global warn") {
			t.Errorf("Expected output to contain global warn, got %q", out.String())
		}
	})

	t.Run("LogError", func(t *testing.T) {
		errOut.Reset()
		LogError("global error")
		if !strings.Contains(errOut.String(), "[ERROR] global error") {
			t.Errorf("Expected stderr to contain global error, got %q", errOut.String())
		}
	})
}

func TestLoggerSetOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(LogLevelInfo, os.Stdout, os.Stderr)
	logger.SetOutput(&buf)
	logger.Info("to buffer")
	if !strings.Contains(buf.String(), "to buffer") {
		t.Errorf("Expected buffer to contain 'to buffer', got %q", buf.String())
	}
}

func TestLogFileDirectoryErrors(t *testing.T) {
	SetLogLevel(LogLevelDebug)
	var out bytes.Buffer
	SetLogOutput(&out)
	SetLogErrorOutput(&out)
	defer func() {
		SetLogOutput(os.Stdout)
		SetLogErrorOutput(os.Stderr)
	}()

	t.Run("LogFileError", func(t *testing.T) {
		out.Reset()
		LogFileError("test.go", errors.New("fail"))
		if !strings.Contains(out.String(), "Failed to process file test.go") {
			t.Errorf("Unexpected output: %s", out.String())
		}
	})

	t.Run("LogDirectoryError", func(t *testing.T) {
		out.Reset()
		LogDirectoryError("testdir", errors.New("fail"))
		if !strings.Contains(out.String(), "Failed to process directory testdir") {
			t.Errorf("Unexpected output: %s", out.String())
		}
	})

	t.Run("Permission errors", func(t *testing.T) {
		out.Reset()
		LogFileError("secret.go", os.ErrPermission)
		LogDirectoryError("secretdir", os.ErrPermission)
		if !strings.Contains(out.String(), "Permission denied") {
			t.Errorf("Expected Permission denied in output, got %q", out.String())
		}
	})
}

func TestIsPermissionError(t *testing.T) {
	if IsPermissionError(os.ErrPermission) == false {
		t.Error("Expected IsPermissionError(os.ErrPermission) to be true")
	}
	if IsPermissionError(errors.New("other error")) == true {
		t.Error("Expected IsPermissionError(other error) to be false")
	}
}

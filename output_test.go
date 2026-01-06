package main

import (
	"errors"
	"strings"
	"testing"
)

func TestPrintResults(t *testing.T) {
	langStats := map[string]*LanguageStats{
		"Go": {
			Language:     "Go",
			FileCount:    1,
			BlankLines:   10,
			CommentLines: 20,
			CodeLines:    70,
			TotalLines:   100,
		},
	}
	total := &LanguageStats{
		Language:     "Total",
		FileCount:    1,
		BlankLines:   10,
		CommentLines: 20,
		CodeLines:    70,
		TotalLines:   100,
	}

	t.Run("Default format", func(t *testing.T) {
		output := captureStdout(func() {
			PrintResults(langStats, total, 1, 0, 0)
		})
		if !strings.Contains(output, "Go") || !strings.Contains(output, "Total") {
			t.Errorf("Output missing expected content: %s", output)
		}
	})

	t.Run("Formatted format", func(t *testing.T) {
		output := captureStdout(func() {
			PrintResultsFormatted(langStats, total, 1, 0, 0)
		})
		if !strings.Contains(output, "Go") || !strings.Contains(output, "Total") {
			t.Errorf("Output missing expected content: %s", output)
		}
	})

	t.Run("JSON format", func(t *testing.T) {
		output := captureStdout(func() {
			PrintJSON(langStats, total)
		})
		if !strings.Contains(output, "\"languages\"") || !strings.Contains(output, "\"Go\"") {
			t.Errorf("Output missing expected content: %s", output)
		}
	})

	t.Run("Compact format", func(t *testing.T) {
		output := captureStdout(func() {
			PrintCompact(total)
		})
		if !strings.Contains(output, "Code: 70") {
			t.Errorf("Output missing expected content: %s", output)
		}
	})

	t.Run("ByFiles format", func(t *testing.T) {
		output := captureStdout(func() {
			PrintByFiles(langStats, total, 1, 0, 0)
		})
		if !strings.Contains(output, "Go") {
			t.Errorf("Output missing expected content: %s", output)
		}
	})
}

func TestPrintErrors(t *testing.T) {
	errs := []error{errors.New("error 1"), errors.New("error 2")}
	output := captureStdout(func() {
		PrintErrors(errs)
	})
	if !strings.Contains(output, "error 1") || !strings.Contains(output, "error 2") {
		t.Errorf("Output missing expected errors: %s", output)
	}

	// Test many errors
	manyErrs := make([]error, 15)
	for i := 0; i < 15; i++ {
		manyErrs[i] = errors.New("error")
	}
	output = captureStdout(func() {
		PrintErrors(manyErrs)
	})
	if !strings.Contains(output, "and 5 more errors") {
		t.Errorf("Output missing 'more errors' message: %s", output)
	}
}

func TestFormatNumber(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{0, "0"},
		{100, "100"},
		{1000, "1,000"},
		{1000000, "1,000,000"},
		{1234567, "1,234,567"},
	}
	for _, tt := range tests {
		if got := FormatNumber(tt.n); got != tt.want {
			t.Errorf("FormatNumber(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

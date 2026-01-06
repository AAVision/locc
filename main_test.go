package main

import (
	"flag"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestSplitAndTrim(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sep      string
		expected []string
	}{
		{
			name:     "Empty string",
			input:    "",
			sep:      ",",
			expected: nil,
		},
		{
			name:     "Single value",
			input:    "val1",
			sep:      ",",
			expected: []string{"val1"},
		},
		{
			name:     "Multiple values",
			input:    "val1,val2,val3",
			sep:      ",",
			expected: []string{"val1", "val2", "val3"},
		},
		{
			name:     "With spaces",
			input:    " val1 ,  val2 ,val3  ",
			sep:      ",",
			expected: []string{"val1", "val2", "val3"},
		},
		{
			name:     "Empty parts",
			input:    "val1,,val2",
			sep:      ",",
			expected: []string{"val1", "val2"},
		},
		{
			name:     "Only separator",
			input:    ",,,",
			sep:      ",",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitAndTrim(tt.input, tt.sep)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("splitAndTrim(%q, %q) = %v, want %v", tt.input, tt.sep, result, tt.expected)
			}
		})
	}
}

func TestParseFlags(t *testing.T) {
	// Save original args and flag set
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	tests := []struct {
		name         string
		args         []string
		wantPath     string
		wantWorkers  int
		wantHidden   bool
		wantFormat   string
		wantExcludes []string
	}{
		{
			name:       "Default values",
			args:       []string{"cmd"},
			wantPath:   ".",
			wantFormat: "default",
		},
		{
			name:        "Specific path and workers",
			args:        []string{"cmd", "-p", "/tmp", "-w", "4"},
			wantPath:    "/tmp",
			wantWorkers: 4,
		},
		{
			name:       "Hidden and format",
			args:       []string{"cmd", "-H", "-f", "json"},
			wantPath:   ".",
			wantHidden: true,
			wantFormat: "json",
		},
		{
			name:         "Exclude dirs",
			args:         []string{"cmd", "-x", "dir1,dir2"},
			wantPath:     ".",
			wantExcludes: []string{"dir1", "dir2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			config := parseFlags()

			if config.Path != tt.wantPath {
				t.Errorf("Path = %q, want %q", config.Path, tt.wantPath)
			}
			if tt.wantWorkers != 0 && config.Workers != tt.wantWorkers {
				t.Errorf("Workers = %d, want %d", config.Workers, tt.wantWorkers)
			}
			if config.IncludeHidden != tt.wantHidden {
				t.Errorf("IncludeHidden = %v, want %v", config.IncludeHidden, tt.wantHidden)
			}
			if tt.wantFormat != "" && config.OutputFormat != tt.wantFormat {
				t.Errorf("OutputFormat = %q, want %q", config.OutputFormat, tt.wantFormat)
			}
			if !reflect.DeepEqual(config.ExcludeDirs, tt.wantExcludes) {
				t.Errorf("ExcludeDirs = %v, want %v", config.ExcludeDirs, tt.wantExcludes)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "main-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	os.WriteFile(filepath.Join(tmpDir, "test.go"), []byte("package main"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "data.unknown"), []byte("unknown"), 0644)

	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "Success directory",
			config: &Config{
				Path:         tmpDir,
				OutputFormat: "default",
				Quiet:        true,
			},
			wantErr: false,
		},
		{
			name: "Success file",
			config: &Config{
				Path:         filepath.Join(tmpDir, "test.go"),
				OutputFormat: "compact",
				Quiet:        true,
			},
			wantErr: false,
		},
		{
			name: "Non-existent path",
			config: &Config{
				Path: "/non/existent/path",
			},
			wantErr: true,
		},
		{
			name: "Unsupported file",
			config: &Config{
				Path: filepath.Join(tmpDir, "data.unknown"),
			},
			wantErr: false,
		},
		{
			name: "Show errors",
			config: &Config{
				Path:       tmpDir,
				ShowErrors: true,
				Quiet:      true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrintUsage(t *testing.T) {
	output := captureStdout(func() {
		printUsage()
	})
	if !strings.Contains(output, "Usage:") {
		t.Errorf("printUsage output missing 'Usage:'")
	}
}

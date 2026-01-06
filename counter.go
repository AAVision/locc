package main

import (
	"bufio"
	"os"
	"strings"
)

// FileStats holds the line count statistics for a single file
type FileStats struct {
	FilePath     string
	Language     string
	Extension    string
	BlankLines   int
	CommentLines int
	CodeLines    int
	TotalLines   int
}

// LanguageStats holds aggregated statistics for a language
type LanguageStats struct {
	Language     string
	FileCount    int
	BlankLines   int
	CommentLines int
	CodeLines    int
	TotalLines   int
}

// CountResult represents the result of counting a file
type CountResult struct {
	Stats *FileStats
	Error error
}

// CountLines counts the lines in a file and categorizes them
func CountLines(filePath string, lang *Language) (*FileStats, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := &FileStats{
		FilePath:  filePath,
		Language:  lang.Name,
		Extension: "",
	}

	scanner := bufio.NewScanner(file)
	// Increase buffer size for files with very long lines
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	inMultiLineComment := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		stats.TotalLines++

		// Check for blank lines
		if trimmedLine == "" {
			stats.BlankLines++
			continue
		}

		// Handle multi-line comments
		if lang.MultiLineStart != "" && lang.MultiLineEnd != "" {
			if inMultiLineComment {
				stats.CommentLines++
				if strings.Contains(trimmedLine, lang.MultiLineEnd) {
					inMultiLineComment = false
				}
				continue
			}

			// Check if line starts a multi-line comment
			if strings.HasPrefix(trimmedLine, lang.MultiLineStart) {
				stats.CommentLines++
				// Check if it also ends on the same line
				if !strings.HasSuffix(trimmedLine, lang.MultiLineEnd) ||
					(lang.MultiLineStart == lang.MultiLineEnd && strings.Count(trimmedLine, lang.MultiLineStart) == 1) {
					// For languages like Python where start and end are the same (""")
					if lang.MultiLineStart == lang.MultiLineEnd {
						count := strings.Count(trimmedLine, lang.MultiLineStart)
						if count == 1 || count%2 == 1 {
							inMultiLineComment = true
						}
					} else if !strings.Contains(trimmedLine[len(lang.MultiLineStart):], lang.MultiLineEnd) {
						inMultiLineComment = true
					}
				}
				continue
			}

			// Check if line contains multi-line comment start somewhere in the middle
			if strings.Contains(trimmedLine, lang.MultiLineStart) {
				// Line has code before the comment start
				if strings.Contains(trimmedLine, lang.MultiLineEnd) {
					// Comment is contained within the line, count as code
					stats.CodeLines++
				} else {
					// Multi-line comment starts but doesn't end
					stats.CodeLines++
					inMultiLineComment = true
				}
				continue
			}
		}

		// Handle single-line comments
		if lang.SingleLineComment != "" && strings.HasPrefix(trimmedLine, lang.SingleLineComment) {
			stats.CommentLines++
			continue
		}

		// Everything else is code
		stats.CodeLines++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

// CountLinesGeneric counts lines for files without specific language support
func CountLinesGeneric(filePath string) (*FileStats, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := &FileStats{
		FilePath: filePath,
		Language: "Unknown",
	}

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		stats.TotalLines++

		if trimmedLine == "" {
			stats.BlankLines++
		} else {
			stats.CodeLines++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

// AggregateStats aggregates file statistics by language
func AggregateStats(fileStats []*FileStats) map[string]*LanguageStats {
	langStats := make(map[string]*LanguageStats)

	for _, fs := range fileStats {
		if fs == nil {
			continue
		}

		lang := fs.Language
		if _, exists := langStats[lang]; !exists {
			langStats[lang] = &LanguageStats{
				Language: lang,
			}
		}

		langStats[lang].FileCount++
		langStats[lang].BlankLines += fs.BlankLines
		langStats[lang].CommentLines += fs.CommentLines
		langStats[lang].CodeLines += fs.CodeLines
		langStats[lang].TotalLines += fs.TotalLines
	}

	return langStats
}

// TotalStats calculates the total statistics across all languages
func TotalStats(langStats map[string]*LanguageStats) *LanguageStats {
	total := &LanguageStats{
		Language: "Total",
	}

	for _, ls := range langStats {
		total.FileCount += ls.FileCount
		total.BlankLines += ls.BlankLines
		total.CommentLines += ls.CommentLines
		total.CodeLines += ls.CodeLines
		total.TotalLines += ls.TotalLines
	}

	return total
}

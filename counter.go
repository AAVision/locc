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
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	inMultiLine := false
	multiLineLevel := 0
	inString := false
	stringEnd := ""

	for scanner.Scan() {
		line := scanner.Text()
		stats.TotalLines++

		lineHasCode := false
		lineHasComment := false

		for i := 0; i < len(line); {
			if inString {
				lineHasCode = true
				if strings.HasPrefix(line[i:], stringEnd) {
					// Check if escaped
					escaped := false
					if i > 0 && line[i-1] == '\\' {
						bsCount := 0
						for j := i - 1; j >= 0 && line[j] == '\\'; j-- {
							bsCount++
						}
						if bsCount%2 == 1 {
							escaped = true
						}
					}
					if !escaped {
						inString = false
						i += len(stringEnd)
					} else {
						i++
					}
				} else {
					i++
				}
				continue
			}

			if inMultiLine {
				lineHasComment = true

				// Check for nested multi-line start
				if lang.NestedComments && lang.MultiLineStart != "" && strings.HasPrefix(line[i:], lang.MultiLineStart) {
					multiLineLevel++
					i += len(lang.MultiLineStart)
					continue
				}

				// Check for multi-line end
				if lang.MultiLineEnd != "" && strings.HasPrefix(line[i:], lang.MultiLineEnd) {
					if multiLineLevel > 0 {
						multiLineLevel--
						i += len(lang.MultiLineEnd)
					} else {
						inMultiLine = false
						i += len(lang.MultiLineEnd)
					}
				} else {
					i++
				}
				continue
			}

			// Not in string or multi-line comment

			// Check for single line comment
			if lang.SingleLineComment != "" && strings.HasPrefix(line[i:], lang.SingleLineComment) {
				lineHasComment = true
				break // Rest of line is comment
			}

			// Check for multi-line comment start
			if lang.MultiLineStart != "" && strings.HasPrefix(line[i:], lang.MultiLineStart) {
				inMultiLine = true
				lineHasComment = true
				i += len(lang.MultiLineStart)
				continue
			}

			// Check for string start
			foundString := false
			for _, delim := range lang.StringDelimiters {
				if strings.HasPrefix(line[i:], delim) {
					inString = true
					stringEnd = delim
					lineHasCode = true
					i += len(delim)
					foundString = true
					break
				}
			}
			if foundString {
				continue
			}

			// Check for code
			if !isWhitespace(line[i]) {
				lineHasCode = true
			}
			i++
		}

		if lineHasCode {
			stats.CodeLines++
		} else if lineHasComment {
			stats.CommentLines++
		} else if stats.TotalLines > 0 { // Should always be true here
			// Check if it's truly blank or just whitespace
			stats.BlankLines++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
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

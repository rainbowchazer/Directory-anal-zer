package report

import (
	"diranalyzer/internal/scanner"
	"fmt"
	"os"
)

func GenReport(files []scanner.FileInfo, reportFile string) error {
	if len(files) == 0 {
		return fmt.Errorf("no files to process")
	}

	var extensions = make(map[string]int)

	for _, file := range files {
		extensions[file.Extension]++
	}

	var report string
	report += "=== File Report ===\n\n"
	report += fmt.Sprintf("Total files: %d\n\n", len(files))
	report += "File types: \n"

	for ext, count := range extensions {
		report += fmt.Sprintf("Count of files with %s extensions: %d\n", ext, count)
	}

	report += "\nFile details:\n"
	report += fmt.Sprintf("%-50s %-10s %-20s %-10s\n", "Name", "Size", "Last Modified", "Extension")

	for _, file := range files {
		report += fmt.Sprintf("%-50s %-10d %-20s %-10s\n", file.FileName, file.ByteSize, file.DateOfChange.Format("2006-01-02 15:04:05"), file.Extension)
	}

	err := os.WriteFile(reportFile, []byte(report), 0644)
	if err != nil {
		return fmt.Errorf("failed to write report to file: %v", err)
	}

	return nil
}

func GenReportExt(files []string, reportFile string) error {
	if len(files) == 0 {
		return fmt.Errorf("no files to process")
	}

	var report string
	report += "=== File Report ===\n\n"
	report += fmt.Sprintf("Total files with this extension: %d\n\n", len(files))
	report += "File names: \n"

	for _, file := range files {
		report += fmt.Sprintf("%s\n", file)
	}

	err := os.WriteFile(reportFile, []byte(report), 0644)
	if err != nil {
		return fmt.Errorf("failed to write report to file: %v", err)
	}

	return nil
}

func GenReportWord(files []string, reportFile string) error {
	if len(files) == 0 {
		return fmt.Errorf("no files to process")
	}

	var report string
	report += "=== File Report ===\n\n"
	report += fmt.Sprintf("Total files with this keyword: %d\n\n", len(files))
	report += "File names: \n"

	for _, file := range files {
		report += fmt.Sprintf("%s\n", file)
	}

	err := os.WriteFile(reportFile, []byte(report), 0644)
	if err != nil {
		return fmt.Errorf("failed to write report to file: %v", err)
	}

	return nil
}

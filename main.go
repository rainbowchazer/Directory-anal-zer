package main

import (
	"bufio"
	"diranalyzer/internal/report"
	"diranalyzer/internal/scanner"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	directory := "C:/Users/NASTY/Documents"
	_, err := os.Stat(directory)
	if err != nil {
		log.Fatal("System error while checking directory existence")
	}

	extension := ".txt"
	keyword := "lab"

	fmt.Println("Do you want to analyse directory(1), find files with extension(2) or files with keywords(3)?")
	reader := bufio.NewReader(os.Stdin)
	opt, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error while entering:", err)
		return
	}
	_, _, _ = reader.ReadLine()

	switch opt {
	case 1:
		startTime := time.Now()
		files, err := scanner.ScanDir(directory)
		if err != nil {
			log.Fatalf("Error while scanning directory: %v", err)
		}

		reportFile := "report.txt"
		err = report.GenReport(files, reportFile)
		if err != nil {
			log.Fatalf("Error while genereting reports: %v", err)
		}

		fmt.Printf("Report was succesfully generated: %s\n", reportFile)
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		fmt.Printf("Время выполнения: %.3f секунд\n", elapsed.Seconds())
	case 2:
		files, err := scanner.FindFilesWithExtension(directory, extension)
		if err != nil {
			log.Fatalf("Error while scanning directory: %v", err)
		}
		reportFile := "report.txt"
		err = report.GenReportExt(files, reportFile)
		if err != nil {
			log.Fatalf("Error while genereting reports: %v", err)
		}

		fmt.Printf("Report was succesfully generated: %s\n", reportFile)
	case 3:
		files, err := scanner.FindFilesWithKeyword(directory, keyword)
		if err != nil {
			log.Fatalf("Error while scanning directory: %v", err)
		}
		reportFile := "report.txt"
		err = report.GenReportExt(files, reportFile)
		if err != nil {
			log.Fatalf("Error while genereting reports: %v", err)
		}

		fmt.Printf("Report was succesfully generated: %s\n", reportFile)
	default:
		fmt.Println("Wrong enter")
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"Vong_Soukorng_G2_Labw3/utils/crack"
)

func main() {
	wordlistPath := flag.String("wordlist", "nord_vpn.txt", "path to wordlist (default: nord_vpn.txt)")
	targetHash := flag.String("hash", "6a85dfd77d9cb35770c9dc6728d73d3f", "MD5 hash to crack (default: the lab hash)")
	verboseFile := flag.String("verbose", "verbose_lab1.txt", "path to verbose output file (default: verbose_lab1.txt)")
	showEvery := flag.Int("show-every", 100000, "print progress every N lines (0 disables)")
	flag.Parse()

	// Prepare logger (stdout + optional file)
	var f *os.File
	var err error
	if *verboseFile != "" {
		dir := filepath.Dir(*verboseFile)
		if dir != "." && dir != "" {
			_ = os.MkdirAll(dir, 0755)
		}
		f, err = os.Create(*verboseFile)
		if err != nil {
			log.Fatalf("Failed to create verbose file: %v", err)
		}
		defer f.Close()
	}

	// log function that writes to stdout and verbose file if provided
	logLine := func(format string, a ...interface{}) {
		t := time.Now().Format("2006-01-02 15:04:05")
		line := fmt.Sprintf("%s: %s\n", t, fmt.Sprintf(format, a...))
		fmt.Print(line)
		if f != nil {
			_, _ = f.WriteString(line)
		}
	}

	// open wordlist
	file, err := os.Open(*wordlistPath)
	if err != nil {
		log.Fatalf("Could not open wordlist: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// make scanner buffer larger for long lines
	const maxCapacity = 10 * 1024 * 1024 // 10MB
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, maxCapacity)

	lineNum := 0
	start := time.Now()
	found := false
	var foundWord string

	for scanner.Scan() {
		lineNum++
		word := scanner.Text()
		// compute md5
		h := crack.HashMD5(word)
		// verbose log every line (comment out if too verbose)
		if f != nil {
			_, _ = f.WriteString(fmt.Sprintf("[%d] word='%s' md5=%s\n", lineNum, word, h))
		}
		if *showEvery > 0 && lineNum%*showEvery == 0 {
			logLine("Processed %d lines (last word: %s). Elapsed: %s", lineNum, word, time.Since(start))
		}
		if h == *targetHash {
			found = true
			foundWord = word
			logLine("MATCH FOUND at line %d -> '%s' md5=%s", lineNum, word, h)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	elapsed := time.Since(start)
	if found {
		logLine("DONE. Cracked: '%s' (hash: %s). Lines checked: %d. Time: %s",
			foundWord, *targetHash, lineNum, elapsed)
	} else {
		logLine("NOT FOUND. Checked %d lines. Time: %s", lineNum, elapsed)
	}
}

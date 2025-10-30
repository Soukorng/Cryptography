package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"Vong_Soukorng_G2_Labw3/utils/crack"
)

func main() {
	// defaults use SHA-1 (40 hex chars)
	wordlistPath := flag.String("wordlist", "nord_vpn.txt", "path to wordlist (default: nord_vpn.txt)")
	targetHash := flag.String("hash", "aa1c7d931cf140bb35a5a16adeb83a551649c3b9", "SHA-1 hash to crack (default: lab hash)")
	verboseFile := flag.String("verbose", "verbose_lab2.txt", "path to verbose output file (default: verbose_lab2.txt)")
	showEvery := flag.Int("show-every", 100000, "print progress every N lines (0 disables)")
	flag.Parse()

	// Normalize the target hash: trim spaces + lowercase
	*targetHash = strings.ToLower(strings.TrimSpace(*targetHash))

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
		// trim whitespace from candidate
		wordTrimmed := strings.TrimSpace(word)
		if wordTrimmed == "" {
			continue
		}

		// compute sha1
		h := strings.ToLower(crack.HashSHA1(wordTrimmed))

		// verbose log every line (comment out if too verbose)
		if f != nil {
			_, _ = f.WriteString(fmt.Sprintf("[%d] word='%s' sha1=%s\n", lineNum, wordTrimmed, h))
		}
		if *showEvery > 0 && lineNum%*showEvery == 0 {
			logLine("Processed %d lines (last word: %s). Elapsed: %s", lineNum, wordTrimmed, time.Since(start))
		}
		if h == *targetHash {
			found = true
			foundWord = wordTrimmed
			logLine("MATCH FOUND at line %d -> '%s' sha1=%s", lineNum, wordTrimmed, h)
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

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
	// defaults use SHA-512 (128 hex chars)
	wordlistPath := flag.String("wordlist", "nord_vpn.txt", "path to wordlist (default: nord_vpn.txt)")
	// your provided SHA-512 target (default)
	targetHash := flag.String("hash", "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38", "SHA-512 hash to crack (default: provided target)")
	verboseFile := flag.String("verbose", "verbose_lab3.txt", "path to verbose output file (default: verbose_lab3.txt)")
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

		// compute sha512
		h := strings.ToLower(crack.HashSHA512(wordTrimmed))

		// verbose log every line (comment out if too verbose)
		if f != nil {
			_, _ = f.WriteString(fmt.Sprintf("[%d] word='%s' sha512=%s\n", lineNum, wordTrimmed, h))
		}
		if *showEvery > 0 && lineNum%*showEvery == 0 {
			logLine("Processed %d lines (last word: %s). Elapsed: %s", lineNum, wordTrimmed, time.Since(start))
		}
		if h == *targetHash {
			found = true
			foundWord = wordTrimmed
			logLine("MATCH FOUND at line %d -> '%s' sha512=%s", lineNum, wordTrimmed, h)
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

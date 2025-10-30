package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
	"golang.org/x/crypto/sha3"
)

//Lab #0: Proof the Hash Program

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("======== Proof the Hash Program ========")
	fmt.Print("\nPlease input value 1: ")
	input1, _ := reader.ReadString('\n')
	fmt.Print("Please input value 2: ")
	input2, _ := reader.ReadString('\n')

	// Remove newline characters
	input1 = strings.TrimSpace(input1)
	input2 = strings.TrimSpace(input2)

	proofMe(input1, input2)
}

func proofMe(txt1, txt2 string) {
	fmt.Println("\n========== Hash Comparison ==========")

	hashCompare("MD5", fmt.Sprintf("%x", md5.Sum([]byte(txt1))), fmt.Sprintf("%x", md5.Sum([]byte(txt2))))
	hashCompare("SHA1", fmt.Sprintf("%x", sha1.Sum([]byte(txt1))), fmt.Sprintf("%x", sha1.Sum([]byte(txt2))))
	hashCompare("SHA256", fmt.Sprintf("%x", sha256.Sum256([]byte(txt1))), fmt.Sprintf("%x", sha256.Sum256([]byte(txt2))))
	hashCompare("SHA512", fmt.Sprintf("%x", sha512.Sum512([]byte(txt1))), fmt.Sprintf("%x", sha512.Sum512([]byte(txt2))))
	hashCompare("SHA3-256", fmt.Sprintf("%x", sha3.Sum256([]byte(txt1))), fmt.Sprintf("%x", sha3.Sum256([]byte(txt2))))
}

func hashCompare(name, hashA, hashB string) {
	fmt.Printf("\nHash (%s):\n", name)
	fmt.Printf("  Output A = %s\n", hashA)
	fmt.Printf("  Output B = %s\n", hashB)

	if hashA == hashB {
		fmt.Println("  => Match!")
	} else {
		fmt.Println("  => No Match!")
	}
}

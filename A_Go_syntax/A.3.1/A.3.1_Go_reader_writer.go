package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "Now is the winter of our discontent," +
		"\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input)) // 1. 스캐너 생성
	scanner.Split(bufio.ScanWords)                        // 2. 단어 단위로 검색

	count := 0

	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}

package wordcounter

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		return len(p), fmt.Errorf("reading input: %v", err)
	}
	return len(p), nil
}

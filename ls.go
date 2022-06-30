package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func HasPrefix(slice []byte, prefix string) bool {
	if len(slice) >= len(prefix) && string(slice[:len(prefix)]) == prefix {
		return true
	}
	return false
}

func main() {
	file, err := os.OpenFile("./log/test.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend|os.FileMode(0644))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

	reader := bufio.NewReader(os.Stdin)
	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	prefix := "Content-Length: "
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		chunks = append(chunks, buf[:n]...)
		if n == cap(buf) {
			continue
		}
		//log.Print(string(chunks))

		i, len_start, len_end := 0, -1, -1
		for i < len(chunks) {
			if len_start < 0 && HasPrefix(chunks[i:], prefix) {
				len_start = i + len(prefix)
				i = len_start + 1
				continue
			}
			if chunks[i] == '\r' {
				len_end = i
				break
			}
			i++
		}
		if len_start >= 0 && len_end >= 0 {
			length, _ := strconv.ParseInt(string(chunks[len_start:len_end]), 10, 64)
			start := len_end + 4
			str := string(chunks[start:start + int(length) + 1])
			log.Print(str)
		}

		chunks = chunks[:0]
  }
}


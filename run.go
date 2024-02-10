package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/bits-and-blooms/bloom/v3"
)

type SafeFileWriter struct {
	file  *os.File
	mutex sync.Mutex
}

func NewSafeFileWriter(file *os.File) *SafeFileWriter {
	return &SafeFileWriter{
		file: file,
	}
}

func (sfw *SafeFileWriter) WriteString(s string) (n int, err error) {
	sfw.mutex.Lock()
	defer sfw.mutex.Unlock()
	return io.WriteString(sfw.file, s)
}

var file1 string
var file2 string
var wordcount uint
var fpRate float64

func main() {
	flag.StringVar(&file1, "receiver", "receiver.txt", "wordlist which will get appended with new entries")
	flag.StringVar(&file2, "sender", "sender.txt", "wordlist which will be included in the receiver after run")
	flag.UintVar(&wordcount, "receiversize", 10000000, "line count of the receiver wordlist")
	flag.Float64Var(&fpRate, "fprate", 0.01, "Rate of false positive with a rejection to append (Default 1%: 0.01) ")

	flag.Parse()
	var wg sync.WaitGroup

	filter := bloom.NewWithEstimates(wordcount, 0.01)

	reciverF, err := os.OpenFile(file1, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer reciverF.Close()
	safeWriter := NewSafeFileWriter(reciverF)

	senderF, err := os.OpenFile(file2, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer senderF.Close()

	reciverSC := bufio.NewScanner(reciverF)
	senderSC := bufio.NewScanner(senderF)
	fmt.Println("[+] Building Filter")
	for reciverSC.Scan() {
		wg.Add(1)
		txt := reciverSC.Text()
		go func(toADD string) {
			defer wg.Done()
			filter.Add([]byte(toADD))

		}(txt)
	}
	if err := reciverSC.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
	wg.Wait()

	fmt.Println("[+] Appending Stuff")
	for senderSC.Scan() {
		text := senderSC.Text()
		wg.Add(1)
		go func(toCheck string) {
			defer wg.Done()
			if !filter.Test([]byte(toCheck)) {
				safeWriter.WriteString("\n" + toCheck)
			}
		}(text)
	}
	if err := senderSC.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
	wg.Wait()

}

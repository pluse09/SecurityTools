package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type ReaderImpl struct{}

type WriterImpl struct{}

func (readerImpl *ReaderImpl) Read(b []byte) (int, error) {
	fmt.Println("in > ")
	return os.Stdin.Read(b)
}

func (writerImpl *WriterImpl) Write(b []byte) (int, error) {
	fmt.Println("out > ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader ReaderImpl
		writer WriterImpl
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}

}

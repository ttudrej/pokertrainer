// package main

import (
	"bytes"
	"fmt"
	"io"
)

// #####################################################

// #####################################################
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// #####################################################
func (bwcPtr *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwcPtr.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwcPtr.buffer.Len() > 8 {

		_, err := bwcPtr.buffer.Read(v)

		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// #####################################################
func (bwcPtr *BufferedWriterCloser) Close() error {
	for bwcPtr.buffer.Len() > 0 {

		data := bwcPtr.buffer.Next(8)
		_, err := fmt.Println(string(data))

		if err != nil {
			return err
		}
	}
	return nil
}

// #####################################################
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

// #####################################################
func main() {

	var myObj interface{} = NewBufferedWriterCloser()

	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello YouTube viewers, this is a test"))
		wc.Close()
	}

	r, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Converson failed")
	}

}

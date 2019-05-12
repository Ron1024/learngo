package main

import "fmt"

type Reader interface {
	Read()
}
type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	Name string
}

func (f *File) Read() {
	fmt.Println("Read file")
}

func (f *File) Write() {
	fmt.Println("Write File")
}

func Test(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	var f File
	// Test(&f)

	var b interface{}
	b = f
	v, ok := b.(ReadWriter)
	fmt.Println(v, ok)
}

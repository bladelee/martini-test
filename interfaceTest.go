package main

import (
	. "bufio"
	"fmt"
	"github.com/codegangsta/inject"
)

type MyReadWriter struct {
	inject.Injector
	*Reader // *bufio.Reader
	*Writer // *bufio.Writer
}

func (rw MyReadWriter) Read(p []byte) (n int, err error) {
	fmt.Print("Read")
	return 0, nil
}

func (rw MyReadWriter) Write(p []byte) (n int, err error) {
	fmt.Print("write")
	return 0, nil
}

func testmain2() {
	rw := new(MyReadWriter)
	rw.MapTo('a', 'b')
	p := []byte{'p', 'k'}
	rw.Read(p)
	rw.Write(p)
	//    rw.reader.Read(p)
	//	rw.write.Read(p)
}

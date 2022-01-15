package rtmp

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type logger int

func (l logger) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	switch {
	case strings.HasPrefix(str, "server") && l >= 1,
		strings.HasPrefix(str, "stream") && l >= 1,
		strings.HasPrefix(str, "event") && l >= 1,
		strings.HasPrefix(str, "data") && l >= 1,
		strings.HasPrefix(str, "msg") && l >= 2:
		l2.Println(str)
	default:
		if l >= 1 {
			l2.Println(str)
		}
	}
}

var (
	l  = logger(0)
	l2 *log.Logger
)

func init() {
	l2 = log.New(os.Stderr, "", 0)
	l2.SetFlags(log.Lmicroseconds)
}

func LogLevel(i int) {
	l = logger(i)
}

type stream struct {
	r io.ReadWriteCloser
}

func (s stream) Read(p []byte) (n int, err error) {
	n, err = s.r.Read(p)
	if err != nil {
		panic(err)
	}
	return
}

func (s stream) Write(p []byte) (n int, err error) {
	n, err = s.r.Write(p)
	if err != nil {
		panic(err)
	}
	return
}

func (s stream) Close() {
	s.r.Close()
}

func ReadBuf(r io.Reader, n int) (b []byte) {
	b = make([]byte, n)
	r.Read(b)
	return
}

func ReadInt(r io.Reader, n int) (ret int) {
	b := ReadBuf(r, n)
	for i := 0; i < n; i++ {
		ret <<= 8
		ret += int(b[i])
	}
	return
}

func ReadIntLE(r io.Reader, n int) (ret int) {
	b := ReadBuf(r, n)
	for i := 0; i < n; i++ {
		ret <<= 8
		ret += int(b[n-i-1])
	}
	return
}

func WriteBuf(w io.Writer, buf []byte) {
	w.Write(buf)
}

func WriteInt(w io.Writer, v int, n int) {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[n-i-1] = byte(v & 0xff)
		v >>= 8
	}
	WriteBuf(w, b)
}

func WriteIntLE(w io.Writer, v int, n int) {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(v & 0xff)
		v >>= 8
	}
	WriteBuf(w, b)
}

package main

import (
	"bytes"
	"io"
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/string-compression

// https://github.com/thefrol/leetcode-go

// нужно скомпрессировать строку по определененным правилам. и заменить исходную строку,
// а вернуть сколько сколько теперь байт в строке

// пишем просто строка входная будет заодно буфером
// compressor воплощает io.Writer

func compress(chars []byte) int {

	comp := compressor{w: bytes.NewBuffer(chars[:0])}
	n, err := comp.Write(chars)
	if err != nil {
		log.Fatal(err)
	}

	return n
}

type compressor struct {
	c     byte
	count int
	w     io.Writer
}

func (comp *compressor) WriteByte(c byte) int {
	n := 0
	if comp.c != c {
		n = comp.Flush()
		comp.c = c
	}
	comp.count++
	return n
}

func (comp *compressor) Reset() {
	comp.c = 0
	comp.count = 0
}

func (comp *compressor) Flush() int {
	if comp.count == 0 {
		return 0
	}

	res := []byte{comp.c}
	if comp.count > 1 {
		res = append(res, []byte(strconv.Itoa(comp.count))...)
	}

	n, err := comp.w.Write(res)
	if err != nil {
		log.Fatal(err)
	}

	comp.Reset()
	return n
}

func (comp *compressor) Write(bb []byte) (int, error) {
	counter := 0
	for _, b := range bb {
		counter += comp.WriteByte(b)
	}
	counter += comp.Flush()
	return counter, nil
}

func Test_compress(t *testing.T) {

	tests := []struct {
		s, want string
	}{
		//{"a", "a"},
		//{"aa", "a2"},
		//{"aab", "a2b"},
		{"aaaaaaaaaaaaaaaaa", "a17"},
		{"aaaaaaaaaaaaaaaaabccc", "a17bc3"},
		{"daaaaaaaaaaaaaaaaabccc", "da17bc3"},
		{"daaaaaaaaaaaaaaaaabcccz", "da17bc3z"},
	}
	for _, tt := range tests {

		t.Run(tt.s, func(t *testing.T) {
			buf := []byte(tt.s)
			n := compress([]byte(buf))
			assert.Equal(t, tt.want, string(buf[0:n]))
		})
	}
}

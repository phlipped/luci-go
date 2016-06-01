// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ar

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

var (
	TestFile1 = ("" +
		// ar file header
		"!<arch>\n" +
		// filename len	- 16 bytes
		"#1/9            " +
		// modtime		- 12 bytes
		"1447140471  " +
		// owner id		- 6 bytes
		"1000  " +
		// group id		- 6 bytes
		"1000  " +
		// file mode	- 8 bytes
		"100640  " +
		// Data size	- 10 bytes
		"15        " +
		// File magic	- 2 bytes
		"\x60\n" +
		// File name	- 9 bytes
		"filename1" +
		// File data	- 6 bytes
		"abc123" +
		// Padding		- 1 byte
		"\n" +
		"")
)

func TestWriterCreatesTestFile1(t *testing.T) {
	b := bytes.NewBufferString("")
	data := []byte("abc123")

	ar := NewWriter(b)
	err := ar.Add("filename1", data)
	if err != nil {
		t.Fatalf("Add: %v", err)
	}
	err = ar.Close()
	if err != nil {
		t.Fatalf("Close: %v", err)
	}

	if g, e := b.Bytes(), []byte(TestFile1); !bytes.Equal(g, e) {
		t.Errorf("\ndata = \n%v\n%v", g, e)
	}
}

func TestReaderOnTestFile1(t *testing.T) {

	r := strings.NewReader(TestFile1)

	ar, err := NewReader(r)
	if err != nil {
		t.Fatalf("NewReader: %v", err)
	}

	h, err := ar.Next()
	if err != nil {
		t.Fatalf("Header: %v", err)
	}
	if g, e := h.Name(), "filename1"; g != e {
		t.Errorf("Name() = %q; want %q", g, e)
	}
	if g, e := h.Size(), int64(6); g != e {
		t.Errorf("Size() = %d; want %d", g, e)
	}
	if g, e := h.ModTime(), time.Unix(1447140471, 0); !g.Equal(e) {
		t.Errorf("ModTime() = %v; want %v", g, e)
	}

	data := make([]byte, 6)
	n, err := ar.Read(data)
	if err != nil {
		t.Fatalf("Data: %v", err)
	}
	if g, e := n, 6; g != e {
		t.Errorf("n = %v; want %v", g, e)
	}
	if g, e := data, []byte("abc123"); !bytes.Equal(g, e) {
		t.Errorf("data = %v; want %v", g, e)
	}

	err = ar.Close()
	if err != nil {
		t.Fatalf("Close: %v", err)
	}
}

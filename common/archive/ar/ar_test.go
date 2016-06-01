// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ar

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestFileInfoHeader(t *testing.T) {

	testfile := (
		"!<arch>\n"			// ar file header
		"#1/5             " // filename len	- 16 bytes
		"123              " // modtime		- 12 bytes
		"1001  "            // owner id		- 6 bytes
		"1002  "            // group id		- 6 bytes
		"000644  "          // file mode	- 8 bytes
		"15      "          // Data size	- 8 bytes
		"\x60\n"            // File magic	- 2 bytes
		"filename1"         // File name	- 9 bytes
		"abc123"            // File data	- 6 bytes
		"\n"                // Padding		- 1 byte
	)

	r := strings.NewReader(testfile)

	ar, err := ar.NewReader(r)
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
	if g, e := h.Size(), 6; g != e {
		t.Errorf("Size() = %d; want %d", g, e)
	}
	if g, e := h.ModTime(), time.Unix(123, 0); !g.Equal(e) {
		t.Errorf("ModTime() = %v; want %v", g, e)
	}

	data := make([]byte, 6)
	h, err := ar.Read(data)
	if err != nil {
		t.Fatalf("Data: %v", err)
	}
	if g, e := data, "abc123"; !g.Equal(e) {
		t.Errorf("data = %v; want %v", g, e)
	}

	err := ar.Close()
	if err != nil {
		t.Fatalf("Close: %v", err)
	}
}

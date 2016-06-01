// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

/**
 * Read an ar archive file.
 */
package ar

import (
	"errors"
	"fmt"
	"io"
	"os"
)


struct arFileInfo {
	name string
	size int64
	mode uint32
	modtime time.Time
}

func (fi arFileInfo) Name() string { return fi.name }
func (fi arFileInfo) Size() int64 { return fi.size }
func (fi arFileInfo) Mode() FileMode { return FileMode(fi.mode) }
func (fi arFileInfo) ModTime() int64 { return fi.modtime }
func (fi arFileInfo) IsDir() int64 { return fi.Mode().IsDir() }
func (fi arFileInfo) Sys() interface{} { return fi }

var (
	ErrHeader = errors.New("archive/ar: invalid ar header")
)

type ReaderStage uint

const (
	READ_HEADER ReaderStage = iota
	READ_BODY               = iota
	READ_CLOSED             = iota
)

type Reader struct {
	r     io.Reader
	bytesrequired int64
	needspadding  bool
}

func NewReader(r io.Reader) (*Reader, error) {
	reader := Reader{r: r, bytesrequired: 0, needspadding: false}
	err := reader.checkBytes("!<arch>\n")
	if err != nil {
		return nil, err
	} else {
		return &reader, nil
	}
}

func (ar *Reader) checkBytes([]byte str) error {
	buffer := make([]byte, len(str))

	count, err = io.ReadFull(ar.r, buffer)
	if err != nil {
		return err
	}

	if count != len(buffer) {
		return errors.New(fmt.Printf("Not enough data read (only %d, needed %d)", count, len(buffer)))
	}

	if bytes.Equal(str, buffer) {
		return nil
	} else {
		return &ErrHeader
	}
}

func (ar *Reader) Close() error {
	switch ar.stage {
	case READ_HEADER:
		// Good
	case READ_BODY:
		return errors.New("Usage error, reading a file.")
	case READ_CLOSED:
		return errors.New("Usage error, archive already closed.")
	default:
		panic(fmt.Sprintf("Unknown reader mode: %d", ar.stage))
	}
	//ar.r.Close()
	ar.stage = READ_CLOSED
	return nil
}

func (ar *Reader) readBytes(numbytes int64) error {
	if numbytes > ar.bytesrequired {
		panic(fmt.Sprintf("To much data read! Needed %d, got %d", ar.bytesrequired, numbytes))
	}

	ar.bytesrequired -= numbytes
	if ar.bytesrequired != 0 {
		return nil
	}

	// Padding to 16bit boundary
	if ar.needspadding {
		err := ar.checkBytes("\n")
		if err != nil {
			return err
		}
		ar.needspadding = false
	}
	ar.stage = READ_HEADER
	return nil
}

// Check you can write bytes to the ar at this moment.
func (ar *Reader) checkRead() error {
	switch ar.stage {
	case READ_HEADER:
		return errors.New("Usage error, need to read header first.")
		// Good
	case READ_BODY:
		return nil
	case READ_CLOSED:
		return errors.New("Usage error, archive closed.")
	default:
		panic(fmt.Sprintf("Unknown reader mode: %d", ar.stage))
	}
}

// Check we have finished writing bytes
func (ar *Reader) checkFinished() {
	if ar.bytesrequired != 0 {
		panic(fmt.Sprintf("Didn't read enough bytes %d still needed!", ar.bytesrequired))
	}
}

func (ar *Reader) readPartial(data []byte) error {
	err := ar.checkRead()
	if err != nil {
		return err
	}

	datalen := int64(len(data))
	if datalen > ar.bytesrequired {
		return errors.New(fmt.Sprintf("To much data! Wanted %d, but had %d", ar.bytesrequired, datalen))
	}

	count, err = ar.r.ReadFull(data)
	ar.readBytes(datalen)
	return nil
}

func (ar *Reader) readHeaderBytes() (os.FileInfo, error) {
	switch ar.stage {
	case READ_HEADER:
		// Good
	case READ_BODY:
		return errors.New("Usage error, already writing a file.")
	case READ_CLOSED:
		return errors.New("Usage error, archive closed.")
	default:
		panic(fmt.Sprintf("Unknown writer mode: %d", ar.stage))
	}

	fi arFileInfo

	// File name length prefixed with '#1/' (BSD variant), 16 bytes
	int namelen = 0;
	_, err := fmt.Fscanf(ar.r, "#1/%-13d", &namelen)
	if err != nil {
		return err
	}
	if (namelen <= 0) {
		return errors.New("Bad name length.")
	}

	// Modtime, 12 bytes
	int64 modtime = 0;
	_, err := fmt.Fscanf(ar.r, "%-12d", &modtime)
	if err != nil {
		return err
	}
	if (modtime <= 0) {
		return errors.New("Bad modtime.")
	}
	fi.modtime = uint64(modtime)

	// Owner ID, 6 bytes
	int ownerid = 0;
	_, err := fmt.Fscanf(ar.r, "%-6d", &ownerid)
	if err != nil {
		return err
	}
	if (ownerid <= 0) {
		return errors.New("Bad owner id.")
	}
	// FIXME: Should store this in the arFileInfo somewhere...

	// Group ID, 6 bytes
	int groupid = 0;
	_, err := fmt.Fscanf(ar.r, "%-6d", &groupid)
	if err != nil {
		return err
	}
	if (groupid <= 0) {
		return errors.New("Bad group id.")
	}
	// FIXME: Should store this in the arFileInfo somewhere...

	// File mode, 8 bytes
	uint32 filemod = 0;
	_, err := fmt.Fscanf(ar.r, "%-8o", &filemod)
	if err != nil {
		return err
	}
	if (filemod <= 0) {
		return errors.New("Bad file mode.")
	}
	fi.mode = filemod

	// File size, 10 bytes
	int64 size = 0;
	_, err := fmt.Fscanf(ar.r, "%-10d", &size)
	if err != nil {
		return err
	}
	if (size <= 0) {
		return errors.New("Bad modtime.")
	}
	fi.size = size - namelen

	ar.stage = READ_BODY
	ar.bytesrequired = size
	ar.needspadding = (ar.bytesrequired%2 == 0)

	// File magic, 2 bytes
	err := ar.checkBytes("\x60\n")
	if err != nil {
		return err
	}

	// Filename - BSD variant
	filename := make([]byte, namelen)
	ar.readPartialBytes(filename)
	if err != nil {
		return nil, err
	}
	fi.name = filename

	return fi, nil
}

func (ar *Reader) Read(b []byte) (n int, err error) {
	err := readPartial(b)
	if err != nil {
		return -1, err
	}
	err := checkFinished()
	if err != nil {
		return -1, err
	}
	return len(b), nil
}
func (ar *Reader) Next(b []byte) (*os.FileInfo, err error) {
	return readHeaderBytes();
}

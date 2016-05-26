/**
 * Write an ar archive file.
 */
package ar

import (
    "errors"
	"fmt"
	"io"
	"os"
)

const DEFAULT_MODTIME = 1447140471
const DEFAULT_USER = 1000
const DEFAULT_GROUP = 1000
const DEFAULT_MODE = 0100640  // 100640 -- Octal

type ArWriterStage uint
const (
	WRITE_HEADER ArWriterStage = iota
	WRITE_BODY = iota
	WRITE_CLOSED = iota
)

type ArWriter struct {
	w io.WriteCloser
	stage ArWriterStage

	bytesrequired int64
	needspadding bool
}

func NewWriter(w io.Writer) *ArWriter {
	io.WriteString(w, "!<arch>\n")
	return &ArWriter{w: w, stage: WRITE_HEADER, bytesrequired: 0, needspadding: 0}
}

func (aw *ArWriter) Close() {
	switch aw.stage {
	case WRITE_HEADER:
		// Good
	case WRITE_BODY:
		return &errorString{"Usage error, writing a file."}
	case WRITE_CLOSED:
		return &errorString{"Usage error, archive already closed."}
	default:
		panic(fmt.Sprintf("Unknown writer mode: %d", aw.stage))
	}
	aw.w.Close()
	aw.stage = WRITE_CLOSED
}

func (aw *ArWriter) wroteBytes(bytes uint64) error {
	if (len(data) > aw.bytesrequired) {
		panic(fmt.Sprintf("To much data written! Needed %d, got %d", aw.bytesrequired, len(data)))
	}

	aw.bytesrequired -= bytes
	if aw.bytesrequired != 0 {
		return nil
	}

	// Padding to 16bit boundary
	if aw.needspadding {
		io.WriteString(aw.w, "\n")
		aw.needspadding = false
    }
	aw.stage = WRITE_HEADER
}

// Check you can write bytes to the ar at this moment.
func (aw *ArWriter) checkWrite() error {
	switch aw.stage {
	case WRITE_HEADER:
		return &errorString{"Usage error, need to write header first."}
		// Good
	case WRITE_BODY:
		return nil
	case WRITE_CLOSED:
		return &errorString{"Usage error, archive closed."}
	default:
		panic(fmt.Sprintf("Unknown writer mode: %d", aw.stage))
	}
}

// Check we have finished writing bytes
func (aw *ArWriter) checkFinished() {
	if (aw.bytesrequired != 0) {
		panic(fmt.Sprintf("Didn't write enough bytes %d still needed, archive corrupted!", aw.bytesrequired))
	}
}

func (aw *ArWriter) writePartial(data []byte) error {
	err := aw.checkWrite()
	if err != nil {
		return err
	}

	if (len(data) > aw.bytesrequired) {
		return &errorString{fmt.Sprintf("To much data! Needed %d, got %d", aw.bytesrequired, len(data))}
	}

	aw.w.Write(data)
	aw.wroteBytes(len(data))
}

func (aw *ArWriter) WriteReader(data io.Reader) error {
	err := aw.checkWrite()
	if err != nil {
		return err
	}

	count, err = io.Copy(aw.w, data)
	if err != nil {
		panic(fmt.Sprintf("err while copying (%s), archive is probably corrupted!", err))
	}
	aw.wroteBytes(count)
	aw.checkFinished()

	return nil
}

func (aw *ArWriter) WriteBytes(data []byte) error {
	err := aw.checkWrite()
	if err != nil {
		return err
	}

	if (len(data) != aw.bytesrequired) {
		return &errorString{fmt.Sprintf("Wrong amount of data! Needed %d, got %d", aw.bytesrequired, len(data))}
	}

	aw.writePartial(data)
	aw.checkFinished()
	return nil
}


func (aw *ArWriter) writeHeaderBytes(name string, size int64, modtime uint64, ownerid uint, groupid uint, filemod uint) error {
	switch aw.stage {
	case WRITE_HEADER:
		// Good
	case WRITE_BODY:
		return &errorString{"Usage error, already writing a file."}
	case WRITE_CLOSED:
		return &errorString{"Usage error, archive closed."}
	default:
		panic(fmt.Sprintf("Unknown writer mode: %d", aw.stage))
	}

    // File name length prefixed with '#1/' (BSD variant), 16 bytes
	fmt.Fprintf(aw.w, "#1/%-13d", len(name))

	// Modtime, 12 bytes
	fmt.Fprintf(aw.w, "%-12d", modtime)

    // Owner ID, 6 bytes
	//fmt.Fprintf(aw.w, "1000  ")
	fmt.Fprintf(aw.w, "%-6d", ownerid

    // Group ID, 6 bytes
    //fmt.Fprintf(aw.w, "1000  ")
    fmt.Fprintf(aw.w, "%-6d", groupid)

	// File mode, 8 bytes
	//fmt.Fprintf(aw.w, "100640  ")
	fmt.Fprintf(aw.w, "%-8o", filemod)

    aw.bytesrequired := len(name)+size

    // File size, 10 bytes
	fmt.Fprintf(aw.w, "%-10d", aw.bytesrequired)

    // File magic, 2 bytes
	io.WriteString(aw.w, "\x60\n")

	aw.stage = WRITE_BODY
	aw.needspadding = (aw.bytesrequired % 2 == 0)

	// Filename - BSD variant
	return aw.writePartial([]byte(name))
}

func (aw *ArWriter) WriteHeaderDefault(name string, size int64) error {
	return aw.writeHeaderBytes(name, size, 1447140471, DEFAULT_USER, DEFAULT_GROUP, DEFAULT_MODE)
}

func (aw *ArWriter) WriteHeader(stat *os.FileInfo) error {
	if (stat.IsDir()) {
		return &errorString{"Only work with files, not directories."}
	}

	mode := stat.Mode()
	if (stat.Mode().ModeSymlink) {
		return &errorString{"Only work with files, not symlinks."}
	}

	/* FIXME: Should we also exclude other "special" files?
	if (stat.Mode().ModeType != 0) {
		return &argError{stat, "Only work with plain files."}
	}
	*/

	// FIXME: Where do we get user/group from - they don't appear to be in Go's Mode() object?
	return aw.writeHeaderBytes(stat.Name(), stat.Size(), stat.ModTime().Unix(), DEFAULT_USER, DEFAULT_GROUP, mode.Mode())
}

func (aw *ArWriter) Add(name string, data []byte) {
	err := aw.WriteHeaderDefault(name, len(data))
	if err != nil {
		return err
	}

	return aw.WriteBytes(data)
}

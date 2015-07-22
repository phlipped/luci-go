// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package iotools

import (
	"net"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

// TestDeadlineReader tests the DeadlineReader struct.
func TestDeadlineReader(t *testing.T) {
	Convey(`A DeadlineReader with a nil connection`, t, func() {
		dr := new(DeadlineReader)

		Convey(`Should panic when Read is called.`, func() {
			dr := dr
			buf := make([]byte, 1)
			So(func() { dr.Read(buf) }, ShouldPanic)
		})

		Convey(`Should panic when Close is called.`, func() {
			dr := dr
			So(func() { dr.Close() }, ShouldPanic)
		})
	})

	// Open a local listening connection.
	Convey(`With a local listening connection`, t, func() {
		proto := "tcp"
		ln, err := net.Listen(proto, "127.0.0.1:0")
		if err != nil {
			proto = "tcp6"
			ln, err = net.Listen(proto, "[::0]:0")
		}
		So(err, ShouldBeNil)

		// Accept connections and don't write any data.
		dataC := make(chan []byte)
		go func() {
			c, err := ln.Accept()
			if err != nil {
				panic("Error while accepting a client connection.")
			}
			defer c.Close()

			data := <-dataC
			if data != nil {
				c.Write(data)
			}
		}()

		// Dial into the local listener.
		c, err := net.Dial(proto, ln.Addr().String())
		So(err, ShouldBeNil)

		// Create a deadline reader.
		dr := &DeadlineReader{Conn: c}

		// Wrap it in a deadline reader.
		Convey(`Given a deadline reader with no deadline, should block on read.`, func() {
			dr := dr
			dataC := dataC
			dr.Deadline = 0 * time.Second

			// Have the server send data (goroutine).
			go func() {
				dataC <- []byte{0xAA}
			}()

			// Connect and read bytes.
			buf := make([]byte, 1)
			amount, err := dr.Read(buf)
			So(err, ShouldBeNil)
			So(amount, ShouldEqual, 1)
		})

		// Wrap it in a deadline reader.
		Convey(`Given a deadline reader with a deadline, should timeout.`, func() {
			dr := dr
			dataC := dataC
			dr.Deadline = 1 * time.Millisecond

			// Have the server send data (goroutine).
			go func() {
				time.Sleep(10 * time.Millisecond)
				dataC <- []byte{0xAA}
			}()

			// Connect and read bytes.
			buf := make([]byte, 1)
			_, err := dr.Read(buf)
			So(err, ShouldNotBeNil)
			switch e := err.(type) {
			case net.Error:
				So(e.Timeout(), ShouldBeTrue)

			default:
				t.Error("Invalid error type.")
			}
		})

		Reset(func() {
			dr.Close()
			ln.Close()
		})
	})
}
// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

// Quick tool for generating directories to walk.

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync/atomic"
	"os"

	"github.com/luci/luci-go/common/dirtools"
	"github.com/luci/luci-go/common/isolated"
)

var method = flag.String("method", "simple", "Method used to walk the tree")
var dir = flag.String("dir", "", "Directory to walk")

//var do = flags.Choice("do", "null", ["null", "print", "read"])
var do = flag.String("do", "nothing", "Action to perform on the files")
var smallfilesize = flag.Int64("smallfilesize", 64*1024, "Size to consider a small file")
var repeat = flag.Int("repeat", 1, "Repeat the walk x times")

// Walker which does nothing but count the files of each type
type NullWalker struct {
	smallfiles uint64
	largefiles uint64
}

func (n *NullWalker) SmallFile(filename string, alldata []byte) {
	atomic.AddUint64(&n.smallfiles, 1)
}
func (n *NullWalker) LargeFile(filename string) {
	atomic.AddUint64(&n.largefiles, 1)
}
func (n *NullWalker) Error(pathname string, err error) {
	log.Fatal("%s:%s", pathname, err)
}

// Walker which just prints the filenames of everything
type PrintWalker struct {
	NullWalker
	obuf io.Writer
}

func (p *PrintWalker) PrintFile(filename string) {
	fmt.Fprintln(p.obuf, filename)
}
func (p *PrintWalker) SmallFile(filename string, alldata []byte) {
	p.NullWalker.SmallFile(filename, alldata)
	p.PrintFile(filename)
}
func (p *PrintWalker) LargeFile(filename string) {
	p.NullWalker.LargeFile(filename)
	p.PrintFile(filename)
}

// Walker which prints the size of everything
type SizeWalker struct {
	NullWalker
	obuf io.Writer
}

func (s *SizeWalker) SizeFile(filename string, size int64) {
	fmt.Fprintf(s.obuf, "%s: %d\n", filename, size)
}
func (s *SizeWalker) SmallFile(filename string, alldata []byte) {
	s.NullWalker.SmallFile(filename, alldata)
	s.SizeFile(filename, int64(len(alldata)))
}
func (s *SizeWalker) LargeFile(filename string) {
	s.NullWalker.LargeFile(filename)
	stat, err := os.Stat(filename)
	if err != nil {
		s.Error(filename, err)
	} else {
		s.SizeFile(filename, stat.Size())
	}
}

// Walker which reads the whole file
type ReadWalker struct {
	NullWalker
}

func (r *ReadWalker) SmallFile(filename string, alldata []byte) {
	r.NullWalker.SmallFile(filename, alldata)
}
func (r *ReadWalker) LargeFile(filename string) {
	r.NullWalker.LargeFile(filename)
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		r.Error(filename, err)
	}
}

// Walker which hashes all the files
type HashWalker struct {
	NullWalker
	obuf io.Writer
}

func (h *HashWalker) HashedFile(filename string, digest isolated.HexDigest) {
	fmt.Fprintf(h.obuf, "%s: %v\n", filename, digest)
}
func (h *HashWalker) SmallFile(filename string, alldata []byte) {
	h.NullWalker.SmallFile(filename, alldata)
	h.HashedFile(filename, isolated.HashBytes(alldata))
}
func (h *HashWalker) LargeFile(filename string) {
	h.NullWalker.LargeFile(filename)
	d, _ := isolated.HashFile(filename)
	h.HashedFile(filename, isolated.HexDigest(d.Digest))
}

// FIXME: Walker which hashes all the files in parallel

func main() {
	flag.Parse()

	if _, err := os.Stat(*dir); err != nil {
		log.Fatal("Directory not found: %s", err)
	}

	var stats *NullWalker
	var obs dirtools.WalkObserver
	switch *do {
	case "nothing":
		o := &NullWalker{}
		stats = o
		obs = o
	case "print":
		o := &PrintWalker{obuf: os.Stderr}
		stats = &o.NullWalker
		obs = o
	case "size":
		o := &SizeWalker{obuf: os.Stderr}
		stats = &o.NullWalker
		obs = o
	case "read":
		o := &ReadWalker{}
		stats = &o.NullWalker
		obs = o
	case "hash":
		o := &HashWalker{obuf: os.Stderr}
		stats = &o.NullWalker
		obs = o
	default:
		log.Fatal("Invalid action '%s'", *do)
	}

	for i := 0; i < *repeat; i++ {
		stats.smallfiles = 0
		stats.largefiles = 0

		switch *method {
		case "simple":
			dirtools.WalkBasic(*dir, *smallfilesize, obs)
		case "nostat":
			dirtools.WalkNoStat(*dir, *smallfilesize, obs)
		case "parallel":
			dirtools.WalkParallel(*dir, *smallfilesize, obs)
		default:
			log.Fatal("Invalid walk method '%s'", *method)
		}
		fmt.Printf("Found %d small files and %d large files\n", stats.smallfiles, stats.largefiles)
	}
	fmt.Fprintf(os.Stderr, "Found %d small files and %d large files\n", stats.smallfiles, stats.largefiles)
}

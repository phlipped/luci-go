// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
//	"path/filepath"
	"time"

	"github.com/luci/luci-go/client/internal/common"
	//"github.com/luci/luci-go/client/isolate"
	"github.com/luci/luci-go/client/isolatedclient"
    "github.com/luci/luci-go/common/api/isolate/isolateservice/v1"
    "github.com/luci/luci-go/common/isolated"

	//"github.com/luci/luci-go/common/units"
	"github.com/maruel/subcommands"

	"github.com/mithro/genar"
)

type ReadSeekerCloser interface {
	io.Reader
	io.Seeker
//	io.Closer
}
type ToHash struct {
	path string
}
type ToCheck struct {
	digest isolateservice.HandlersEndpointsV1Digest
	name string
	source ReadSeekerCloser
}
type ToPush struct {
	state *isolatedclient.PushState
	name string
	source ReadSeekerCloser
}

func HashFile(is isolatedclient.IsolateServer, src <-chan *ToHash, dst chan<- *ToCheck) {
	for tohash := range src {
		fmt.Printf("hashing %s\n", tohash.path)
		d, _ := isolated.HashFile(tohash.path)
		f, _ := os.Open(tohash.path)
		dst <- &ToCheck{
			digest: d,
			source: f,
			name: tohash.path,
		}
	}
	close(dst)
}

const CHECK_BATCH_SIZE = 20;
func ChckFile(is isolatedclient.IsolateServer, canceler common.Canceler, src <-chan *ToCheck, dst chan<- *ToPush) {
	check_count := 0

	pool := common.NewGoroutinePool(5, canceler)
	defer func() {
		_ = pool.Wait()
	}()

	done := false
	for !done {
		var digests [CHECK_BATCH_SIZE]*isolateservice.HandlersEndpointsV1Digest
		var topush [CHECK_BATCH_SIZE]ToPush

		index := 0
Loop:
		for index < CHECK_BATCH_SIZE && !done {
			select {
			case tocheck, more := <-src:
				if (!more) {
					done = true
					break Loop
				}
				digests[index] = &tocheck.digest
				topush[index] = ToPush{state:nil, source:tocheck.source, name:tocheck.name}
				index += 1
			case <-time.After(time.Millisecond * 10):
				break Loop
			}
		}

		if index > 0 {
			inner_count := check_count
			inner_index := index
			pool.Schedule(func() {
				fmt.Printf("checking(%d) %d files\n", inner_count, inner_index)
				pushstates, err := is.Contains(digests[:inner_index])
				if err != nil {
					fmt.Println("checking(%d) error:", inner_count, err)
					return
				}
				for j, state := range pushstates {
					topush[j].state = state
					if state != nil {
						fmt.Printf("need to push(%d): %s\n", inner_count, topush[j].name)
						dst <- &topush[j]
					} else {
						fmt.Printf("skipping(%d): %s\n", inner_count, topush[j].name)
//						sources[j].Close()
					}
				}
			}, func() {})
			check_count += 1
		}
	}
	_ = pool.Wait()
	close(dst)
}


func PushFile(is isolatedclient.IsolateServer, canceler common.Canceler, src <-chan *ToPush, dst chan<- bool) {
	pool := common.NewGoroutinePool(100, canceler)
	defer func() {
		_ = pool.Wait()
	}()

	for topush := range src {
		pool.Schedule(func() {
			fmt.Printf("pushing: %s\n", topush.name)
			err := is.Push(topush.state, topush.source)
			if err != nil {
				fmt.Println("pushing err:", err)
			} else {
				fmt.Println("pushed:", topush.state)
			}
//			topush.source.Close()
		}, func() {})
	}
	_ = pool.Wait()
	close(dst)
}

// ---
type SmallFilesCollection struct {
	index int
	buffer *bytes.Buffer
	hash hash.Hash
	ar *genar.ArWriter
}

func NewSmallFilesCollection(index int) *SmallFilesCollection {
	var o SmallFilesCollection
	o.index = index
	o.buffer = new(bytes.Buffer)
	o.hash = isolated.GetHash()

	var w io.Writer = o.buffer
	w = io.MultiWriter(w, o.hash)
	o.ar = genar.NewWriter(w)
	return &o
}

func (b SmallFilesCollection) RequestCheck(dst chan<- *ToCheck) {
	fmt.Printf("rotating smallfilescollection-%d\n", b.index)
	dst <- &ToCheck{
		digest: isolateservice.HandlersEndpointsV1Digest{
			Digest: string(isolated.Sum(b.hash)),
			IsIsolated: false,
			Size: int64(b.buffer.Len()),
		},
		source: bytes.NewReader(b.buffer.Bytes()),
		name: fmt.Sprintf("smallfilescollection-%d", b.index),
	}
}
//

const SMALLFILES_MAXSIZE = 1024*1024*100 // 100MBytes

func upload(is isolatedclient.IsolateServer, path string) {
	hash_chan := make(chan *ToHash, 10)
	chck_chan := make(chan *ToCheck, 1)
	push_chan := make(chan *ToPush, 10)
	done_chan := make(chan bool)

	canceler := common.NewCanceler()

	go HashFile(is, hash_chan, chck_chan)
	go ChckFile(is, canceler, chck_chan, push_chan)
	go PushFile(is, canceler, push_chan, done_chan)

	paths := []string{ path }
	smallfiles_buffer := NewSmallFilesCollection(0)
	var largefiles_queue []string
	errors := genar.FastWalk(
		"", paths,
		func(name string, data []byte) {
			//fmt.Println("smallfile", name)
			smallfiles_buffer.ar.Write(name[len(path)+1:], data)
			if (smallfiles_buffer.buffer.Len() > SMALLFILES_MAXSIZE) {
				smallfiles_buffer.RequestCheck(chck_chan)
				smallfiles_buffer = NewSmallFilesCollection(smallfiles_buffer.index+1)
			}
		},
		func(name string) {
			largefiles_queue = append(largefiles_queue, name)
		},
	)
	smallfiles_buffer.RequestCheck(chck_chan)

	for _, err := range errors {
		fmt.Println(err.Name, err.Err)
	}

	for _, name := range largefiles_queue {
		hash_chan <- &ToHash{name}
	}

	close(hash_chan)
	<-done_chan
}

var cmdFastArchive = &subcommands.Command{
	UsageLine: "fastarchive <options>",
	ShortDesc: "creates a .isolated file and uploads the tree to an isolate server.",
	LongDesc:  "All the files listed in the .isolated file are put in the isolate server cache via isolateserver.py.",
	CommandRun: func() subcommands.CommandRun {
		c := fastArchiveRun{}
		c.commonServerFlags.Init()
		c.isolateFlags.Init(&c.Flags)
		return &c
	},
}

type fastArchiveRun struct {
	commonServerFlags
	isolateFlags
}

func (c *fastArchiveRun) Parse(a subcommands.Application, args []string) error {
	if err := c.commonServerFlags.Parse(); err != nil {
		return err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := c.isolateFlags.Parse(cwd, RequireIsolatedFile); err != nil {
		return err
	}
	if len(args) != 0 {
		return errors.New("position arguments not expected")
	}
	return nil
}

func (c *fastArchiveRun) main(a subcommands.Application, args []string) error {
/*
	out := os.Stdout
	prefix := "\n"
	if c.defaultFlags.Quiet {
		out = nil
		prefix = ""
	}
	start := time.Now()
*/
	client, err := c.createAuthClient()
	if err != nil {
		return err
	}

	is := isolatedclient.New(client, c.isolatedFlags.ServerURL, c.isolatedFlags.Namespace)
	fmt.Println(c.Isolate)
	upload(is, c.Isolate)

	return nil
}

func (c *fastArchiveRun) Run(a subcommands.Application, args []string) int {
	if err := c.Parse(a, args); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	cl, err := c.defaultFlags.StartTracing()
	if err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	defer cl.Close()
	if err := c.main(a, args); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	return 0
}

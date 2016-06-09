// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

/**

WARNING: THIS FUNCTION IS SLOWER THAN THE NON-PARALLEL VERSION!

**/
package dirtools

import (
	"github.com/eapache/channels"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync/atomic"
)

type fileQueue struct {
	queued   uint64
	finished uint64
	items    channels.Channel
	waiton   chan bool
}

func (q *fileQueue) add(s string) {
	atomic.AddUint64(&q.queued, 1)
	q.items.In() <- s
}

func (q *fileQueue) done() {
	atomic.AddUint64(&q.finished, 1)

	if q.queued == q.finished {
		q.items.Close()
		q.waiton <- true
	}
}

func (q *fileQueue) wait() {
	<-q.waiton
}

func examinePath(queue *fileQueue, smallfile_limit int64, obs WalkObserver) {
	for ipath := range queue.items.Out() {
		path := ipath.(string)

		fi, err := os.Stat(path)
		if err != nil {
			obs.Error(path, err)
			return
		}

		if fi.IsDir() {
			f, err := os.Open(path)
			if err != nil {
				obs.Error(path, err)
			}

			dircontents, err := f.Readdirnames(-1)
			if err != nil {
				obs.Error(path, err)
			}
			sort.Strings(dircontents)
			for _, name := range dircontents {
				fname := filepath.Join(path, name)
				queue.add(fname)
			}
		} else {
			if fi.Size() < smallfile_limit {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					obs.Error(path, err)
					return
				}
				if int64(len(data)) != fi.Size() {
					panic("file size was wrong!")
				}
				obs.SmallFile(path, data)
			} else {
				obs.LargeFile(path)
			}
		}
		queue.done()
	}
}

func WalkParallel(root string, smallfile_limit int64, obs WalkObserver) {
	queue := fileQueue{queued: 0, finished: 0, items: channels.NewInfiniteChannel(), waiton: make(chan bool)}

	for w := 0; w <= 10; w++ {
		go examinePath(&queue, smallfile_limit, obs)
	}

	queue.add(root)
	queue.wait()
}

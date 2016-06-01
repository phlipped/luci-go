// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

/**

This function works strangely for performance reasons, I'll try and explain below.

File systems have been heavily optimised for doing a directory walk in inode
order. It can be an order of magnitude faster to walk the directory this way.
*However*, we want out output to be in sorted order so it is deterministic.

Calling `stat` a file is one of the most expensive things you can do. It is
equivalent to reading 64/128k of data. Hence, if you have a lot of small files
then just reading their contents directly is more efficient.

**/
package dirtools

import (
	"io"
	"os"
	"path/filepath"
	"sort"
)

type SmallFile struct {
	name string
	data []byte
}
type SmallFileByName []SmallFile

func (a SmallFileByName) Len() int           { return len(a) }
func (a SmallFileByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SmallFileByName) Less(i, j int) bool { return a[i].name < a[j].name }

/*
type LargeFile struct {
	name string
}
type LargeFileByName []LargeFile
func (a LargeFileByName) Len() int { return len(a) }
func (a LargeFileByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a LargeFileByName) Less(i, j int) bool { return a[i].name < a[j].name }
*/

type EntryError struct {
	name string
	err  error
}
type EntryErrorByName []EntryError

func (a EntryErrorByName) Len() int           { return len(a) }
func (a EntryErrorByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a EntryErrorByName) Less(i, j int) bool { return a[i].name < a[j].name }

func walkNoStatInternal(base string, files []string, smallfile_limit int64, obs WalkObserver) {
	var errors []EntryError
	var smallfiles []SmallFile
	var largefiles []string
	var dirs []EntryError

	for _, name := range files {
		fname := filepath.Join(base, name)
		file, err := os.Open(fname)

		if err != nil {
			errors = append(errors, EntryError{fname, err})
			continue
		}

		block := make([]byte, smallfile_limit)
		count, err := file.Read(block)
		if err != io.EOF && err != nil {
			// Its probably a directory
			dirs = append(dirs, EntryError{fname, err})
			continue
		}

		// This file was bigger than the block size, stat it
		if int64(count) == smallfile_limit {
			/*
				stat, err := file.Stat()
				if err != nil {
					errors = append(errors, EntryError{fname, err})
					continue
				}
			*/
			largefiles = append(largefiles, fname) //LargeFile{name: fname, stat: &stat})

			// This file was smaller than the block size
		} else {
			smallfiles = append(smallfiles, SmallFile{name: fname, data: block[:count]})
		}
		file.Close()
	}

	sort.Sort(SmallFileByName(smallfiles))
	for _, f := range smallfiles {
		obs.SmallFile(f.name, f.data)
	}

	sort.Strings(largefiles)
	for _, fname := range largefiles {
		obs.LargeFile(fname)
	}

	sort.Sort(EntryErrorByName(dirs))
	for _, d := range dirs {
		file, err := os.Open(d.name)
		if err != nil {
			obs.Error(d.name, d.err)
			continue
		}

		names, err := file.Readdirnames(0)
		if err != nil {
			obs.Error(d.name, d.err)
			continue
		}
		walkNoStatInternal(d.name, names, smallfile_limit, obs)
	}
}

func WalkNoStat(root string, smallfile_limit int64, obs WalkObserver) {
	paths := []string{root}
	walkNoStatInternal("", paths, smallfile_limit, obs)
}

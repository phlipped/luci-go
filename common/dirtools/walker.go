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
	"io/ioutil"
	"os"
	"path/filepath"
)

/**
SmallFile and LargeFile must be called in sorted order.
*/
type WalkObserver interface {
	SmallFile(filename string, alldata []byte)
	LargeFile(filename string)

	//StartDir(dirname string) error
	//FinishDir(dirname string)

	Error(pathname string, err error)
}

func WalkBasic(root string, smallfile_limit int64, obs WalkObserver) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			obs.Error(path, err)
			return nil
		}

		if info.Size() < smallfile_limit {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				obs.Error(path, err)
				return nil
			}
			if int64(len(data)) != info.Size() {
				panic("file size was wrong!")
			}
			obs.SmallFile(path, data)
		} else {
			obs.LargeFile(path)
		}
		return nil
	})
}

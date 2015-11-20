// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package bootstrap

// Environment variable names
const (
	// EnvStreamServerPath is the path to the Butler's stream server endpoint.
	//
	// This can be used by applications to initiate a new Butler stream with an
	// existing Butler stream server process. If a subprocess is launched with a
	// stream server configuration, it will propagate this path to its child
	// processes.
	EnvStreamServerPath = "LOGDOG_STREAM_SERVER_PATH"

	// EnvStreamPrefix is the environment variable set to the configured
	// stream name prefix. command-line stream name prefix.
	//
	// Since this name is prepended to streams by the Butler, it is not possible
	// for subprocesses to know the actual stream name that Butler will use for
	// their streams without this information.
	EnvStreamPrefix = "LOGDOG_STREAM_PREFIX"
)

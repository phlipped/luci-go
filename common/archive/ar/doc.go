// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package ar implements access to ar archives. ar is probably the simplest
// format that standard tools under Linux support.
//
// The base ar format only supports files which are 16 characters long. There
// are multiple methods for supporting longer file names.
//
// This package only supports the "BSD variant" as the filename is stored allow
// "streaming" support (rather than the Sys V / GNU variant which requires
// building a symbol table).
//
// References:
// * https://en.wikipedia.org/wiki/Ar_(Unix)
package archive

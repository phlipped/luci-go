// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package test

import "github.com/luci/luci-go/tools/internal/svctool/testdata"

// CompoundServer embeds an interface.
type CompoundServer interface {
	test.S1Server
}

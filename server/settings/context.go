// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package settings

import (
	"golang.org/x/net/context"
)

type contextKey int

// Use injects Settings into the context to be used by Get and Set.
func Use(c context.Context, s *Settings) context.Context {
	return context.WithValue(c, contextKey(0), s)
}

// GetSettings grabs Settings from the context if it's there.
func GetSettings(c context.Context) *Settings {
	if s, ok := c.Value(contextKey(0)).(*Settings); ok && s != nil {
		return s
	}
	return nil
}

// Get returns setting value for the given key. It will be deserialized into
// the supplied value. Caller is responsible to pass correct type here. If
// the setting is not set or the context doesn't have settings implementation
// in it, returns ErrNoSettings.
func Get(c context.Context, key string, value interface{}) error {
	if s := GetSettings(c); s != nil {
		return s.Get(c, key, value)
	}
	return ErrNoSettings
}

// Set changes a setting value for the given key. New settings will apply only
// when existing in-memory cache expires. In particular, Get() right after Set()
// may still return old value. Returns ErrNoSettings if context doesn't have
// Settings implementation.
func Set(c context.Context, key string, value interface{}, who, why string) error {
	if s := GetSettings(c); s != nil {
		return s.Set(c, key, value, who, why)
	}
	return ErrNoSettings
}
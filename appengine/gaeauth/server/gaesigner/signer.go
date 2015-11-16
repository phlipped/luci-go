// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package gaesigner implements signing.Signer interface using GAE App Identity
// API.
package gaesigner

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"

	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/server/auth/signing"
	"github.com/luci/luci-go/server/proccache"
)

// Use installs signing.Signer into the context configured to use GAE API.
func Use(c context.Context) context.Context {
	return signing.SetSigner(c, signer{})
}

////

// signer implements signing.Signer using GAE App Identity API.
type signer struct{}

func (signer) SignBytes(c context.Context, blob []byte) (keyName string, signature []byte, err error) {
	return appengine.SignBytes(c, blob)
}

func (signer) Certificates(c context.Context) (*signing.PublicCertificates, error) {
	certs, err := cachedCerts(c)
	if err != nil {
		return nil, err
	}
	return certs.(*signing.PublicCertificates), nil
}

////

type certsCacheKey int

// cachedCerts caches this app certs in local memory for 1 hour.
var cachedCerts = proccache.Cached(certsCacheKey(0), func(c context.Context, key interface{}) (interface{}, time.Duration, error) {
	aeCerts, err := appengine.PublicCertificates(c)
	if err != nil {
		return nil, 0, err
	}
	certs := make([]signing.Certificate, len(aeCerts))
	for i, ac := range aeCerts {
		certs[i] = signing.Certificate{
			KeyName:            ac.KeyName,
			X509CertificatePEM: string(ac.Data),
		}
	}
	return &signing.PublicCertificates{
		Certificates: certs,
		Timestamp:    signing.JSONTime(clock.Now(c)),
	}, time.Hour, nil
})
// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package collector

import (
	"bytes"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/luci/luci-go/common/clock/testclock"
	"github.com/luci/luci-go/common/config"
	"github.com/luci/luci-go/common/errors"
	"github.com/luci/luci-go/common/logdog/butlerproto"
	"github.com/luci/luci-go/common/logdog/types"
	"github.com/luci/luci-go/common/proto/logdog/logpb"
	"github.com/luci/luci-go/server/internal/logdog/collector/coordinator"
	"github.com/luci/luci-go/server/logdog/storage/memory"
	"golang.org/x/net/context"

	. "github.com/luci/luci-go/common/testing/assertions"
	. "github.com/smartystreets/goconvey/convey"
)

// TestCollector runs through a series of end-to-end Collector workflows and
// ensures that the Collector behaves appropriately.
func testCollectorImpl(t *testing.T, caching bool) {
	Convey(fmt.Sprintf(`Using a test configuration with caching == %v`, caching), t, func() {
		c, _ := testclock.UseTime(context.Background(), testclock.TestTimeLocal)

		tcc := &testCoordinator{}
		st := &testStorage{Storage: &memory.Storage{}}

		coll := &Collector{
			Coordinator: tcc,
			Storage:     st,
		}
		defer coll.Close()

		bb := bundleBuilder{
			Context: c,
		}

		if caching {
			coll.Coordinator = coordinator.NewCache(coll.Coordinator, 0, 0)
		}

		Convey(`Can process multiple single full streams from a Butler bundle.`, func() {
			bb.addFullStream("foo/+/bar", 128)
			bb.addFullStream("foo/+/baz", 256)

			So(coll.Process(c, bb.bundle()), ShouldBeNil)

			So(tcc, shouldHaveRegisteredStream, "test-project", "foo/+/bar", 127)
			So(st, shouldHaveStoredStream, "test-project", "foo/+/bar", indexRange{0, 127})

			So(tcc, shouldHaveRegisteredStream, "test-project", "foo/+/baz", 255)
			So(st, shouldHaveStoredStream, "test-project", "foo/+/baz", indexRange{0, 255})
		})

		Convey(`Will return a transient error if a transient error happened while registering.`, func() {
			tcc.errC = make(chan error, 1)
			tcc.errC <- errors.WrapTransient(errors.New("test error"))

			bb.addFullStream("foo/+/bar", 128)
			err := coll.Process(c, bb.bundle())
			So(err, ShouldNotBeNil)
			So(errors.IsTransient(err), ShouldBeTrue)
		})

		Convey(`Will return an error if a non-transient error happened while registering.`, func() {
			tcc.errC = make(chan error, 1)
			tcc.errC <- errors.New("test error")

			bb.addFullStream("foo/+/bar", 128)
			err := coll.Process(c, bb.bundle())
			So(err, ShouldNotBeNil)
			So(errors.IsTransient(err), ShouldBeFalse)
		})

		Convey(`Will return a transient error if a transient error happened while terminating.`, func() {
			tcc.errC = make(chan error, 2)
			tcc.errC <- nil                                            // Register
			tcc.errC <- errors.WrapTransient(errors.New("test error")) // Terminate

			bb.addFullStream("foo/+/bar", 128)
			err := coll.Process(c, bb.bundle())
			So(err, ShouldNotBeNil)
			So(errors.IsTransient(err), ShouldBeTrue)
		})

		Convey(`Will return an error if a non-transient error happened while terminating.`, func() {
			tcc.errC = make(chan error, 2)
			tcc.errC <- nil                      // Register
			tcc.errC <- errors.New("test error") // Terminate

			bb.addFullStream("foo/+/bar", 128)
			err := coll.Process(c, bb.bundle())
			So(err, ShouldNotBeNil)
			So(errors.IsTransient(err), ShouldBeFalse)
		})

		Convey(`Will return a transient error if a transient error happened on storage.`, func() {
			// Single transient error.
			count := int32(0)
			st.err = func() error {
				if atomic.AddInt32(&count, 1) == 1 {
					return errors.WrapTransient(errors.New("test error"))
				}
				return nil
			}

			bb.addFullStream("foo/+/bar", 128)
			err := coll.Process(c, bb.bundle())
			So(err, ShouldNotBeNil)
			So(errors.IsTransient(err), ShouldBeTrue)
		})

		Convey(`Will drop invalid LogStreamDescriptor bundle entries and process the valid ones.`, func() {
			be := bb.genBundleEntry("foo/+/trash", 1337, 4, 5, 6, 7, 8)
			bb.addBundleEntry(be)

			bb.addStreamEntries("foo/+/trash", 0, 1, 3) // Invalid: non-contiguous
			bb.addFullStream("foo/+/bar", 32)

			err := coll.Process(c, bb.bundle())
			So(err, ShouldNotBeNil)
			So(errors.IsTransient(err), ShouldBeFalse)

			So(tcc, shouldHaveRegisteredStream, "test-project", "foo/+/bar", 32)
			So(st, shouldHaveStoredStream, "test-project", "foo/+/bar", indexRange{0, 31})

			So(tcc, shouldHaveRegisteredStream, "test-project", "foo/+/trash", 1337)
			So(st, shouldHaveStoredStream, "test-project", "foo/+/trash", 4, 5, 6, 7, 8)
		})

		Convey(`Will drop streams with missing (invalid) secrets.`, func() {
			b := bb.genBase()
			b.Secret = nil
			bb.addFullStream("foo/+/bar", 4)

			err := coll.Process(c, bb.bundle())
			So(err, ShouldErrLike, "invalid prefix secret")
			So(errors.IsTransient(err), ShouldBeFalse)
		})

		// TODO(dnj): Remove this test once this is no longer supported.
		Convey(`Will use legacy bundle entry secret if bundle secret is missing.`, func() {
			b := bb.genBase()

			be := bb.genBundleEntry("foo/+/bar", 4, 0, 1, 2, 3)
			be.DeprecatedEntrySecret = b.Secret
			b.Secret = nil
			bb.addBundleEntry(be)

			err := coll.Process(c, bb.bundle())
			So(err, ShouldBeNil)
		})

		Convey(`Will drop messages with mismatching secrets.`, func() {
			bb.addStreamEntries("foo/+/bar", -1, 0, 1, 2)
			So(coll.Process(c, bb.bundle()), ShouldBeNil)

			// Push another bundle with a different secret.
			b := bb.genBase()
			b.Secret = bytes.Repeat([]byte{0xAA}, types.PrefixSecretLength)
			be := bb.genBundleEntry("foo/+/bar", 4, 3, 4)
			be.TerminalIndex = 1337
			bb.addBundleEntry(be)
			bb.addFullStream("foo/+/baz", 3)
			So(coll.Process(c, bb.bundle()), ShouldBeNil)

			So(tcc, shouldHaveRegisteredStream, "test-project", "foo/+/bar", -1)
			So(st, shouldHaveStoredStream, "test-project", "foo/+/bar", indexRange{0, 2})

			So(tcc, shouldHaveRegisteredStream, "test-project", "foo/+/baz", 2)
			So(st, shouldHaveStoredStream, "test-project", "foo/+/baz", indexRange{0, 2})
		})

		Convey(`With an empty project name`, func() {
			b := bb.genBase()
			b.Project = ""
			bb.addFullStream("foo/+/baz", 3)

			// TODO(dnj): Enable this when project name is required.
			SkipConvey(`Will drop the stream.`, func() {

				err := coll.Process(c, bb.bundle())
				So(err, ShouldErrLike, "invalid project name")
				So(errors.IsTransient(err), ShouldBeFalse)
			})

			Convey(`Will register the stream.`, func() {
				So(coll.Process(c, bb.bundle()), ShouldBeNil)

				So(tcc, shouldHaveRegisteredStream, "", "foo/+/baz", 2)
				So(st, shouldHaveStoredStream, "", "foo/+/baz", indexRange{0, 2})
			})
		})

		Convey(`Will drop streams with invalid project names.`, func() {
			b := bb.genBase()
			b.Project = "!!!invalid name!!!"
			So(config.ProjectName(b.Project).Validate(), ShouldNotBeNil)

			err := coll.Process(c, bb.bundle())
			So(err, ShouldErrLike, "invalid bundle project name")
			So(errors.IsTransient(err), ShouldBeFalse)
		})

		Convey(`Will drop streams with empty bundle prefixes.`, func() {
			b := bb.genBase()
			b.Prefix = ""

			err := coll.Process(c, bb.bundle())
			So(err, ShouldErrLike, "invalid bundle prefix")
			So(errors.IsTransient(err), ShouldBeFalse)
		})

		Convey(`Will drop streams with invalid bundle prefixes.`, func() {
			b := bb.genBase()
			b.Prefix = "!!!invalid prefix!!!"
			So(types.StreamName(b.Prefix).Validate(), ShouldNotBeNil)

			err := coll.Process(c, bb.bundle())
			So(err, ShouldErrLike, "invalid bundle prefix")
			So(errors.IsTransient(err), ShouldBeFalse)
		})

		Convey(`Will drop streams whose descriptor prefix doesn't match its bundle's prefix.`, func() {
			bb.addStreamEntries("baz/+/bar", 3, 0, 1, 2, 3, 4)

			err := coll.Process(c, bb.bundle())
			So(err, ShouldErrLike, "mismatched bundle and entry prefixes")
			So(errors.IsTransient(err), ShouldBeFalse)
		})

		Convey(`Will return no error if the data has a corrupt bundle header.`, func() {
			So(coll.Process(c, []byte{0x00}), ShouldBeNil)
			So(tcc, shouldNotHaveRegisteredStream, "test-project", "foo/+/bar")
		})

		Convey(`Will drop bundles with unknown ProtoVersion string.`, func() {
			buf := bytes.Buffer{}
			w := butlerproto.Writer{ProtoVersion: "!!!invalid!!!"}
			w.Write(&buf, &logpb.ButlerLogBundle{})

			So(coll.Process(c, buf.Bytes()), ShouldBeNil)

			So(tcc, shouldNotHaveRegisteredStream, "test-project", "foo/+/bar")
		})

		Convey(`Will not ingest records if the stream is archived.`, func() {
			tcc.register(coordinator.LogStreamState{
				Project:       "test-project",
				Path:          "foo/+/bar",
				Secret:        testSecret,
				TerminalIndex: -1,
				Archived:      true,
			})

			bb.addStreamEntries("foo/+/bar", 3, 0, 1, 2, 3, 4)
			So(coll.Process(c, bb.bundle()), ShouldBeNil)

			So(tcc, shouldHaveRegisteredStream, "test-project", "foo/+/bar", -1)
			So(st, shouldHaveStoredStream, "test-project", "foo/+/bar")
		})

		Convey(`Will not ingest records if the stream is purged.`, func() {
			tcc.register(coordinator.LogStreamState{
				Project:       "test-project",
				Path:          "foo/+/bar",
				Secret:        testSecret,
				TerminalIndex: -1,
				Purged:        true,
			})

			So(coll.Process(c, bb.bundle()), ShouldBeNil)

			So(tcc, shouldHaveRegisteredStream, "test-project", "foo/+/bar", -1)
			So(st, shouldHaveStoredStream, "test-project", "foo/+/bar")
		})

		Convey(`Will not ingest a bundle with no bundle entries.`, func() {
			So(coll.Process(c, bb.bundle()), ShouldBeNil)
		})

		Convey(`Will not ingest a bundle whose log entries don't match their descriptor.`, func() {
			be := bb.genBundleEntry("foo/+/bar", 4, 0, 1, 2, 3, 4)

			// Add a binary log entry. This does NOT match the text descriptor, and
			// should fail validation.
			be.Logs = append(be.Logs, &logpb.LogEntry{
				StreamIndex: 2,
				Sequence:    2,
				Content: &logpb.LogEntry_Binary{
					&logpb.Binary{
						Data: []byte{0xd0, 0x6f, 0x00, 0xd5},
					},
				},
			})
			bb.addBundleEntry(be)
			So(coll.Process(c, bb.bundle()), ShouldErrLike, "invalid log entry")

			So(tcc, shouldNotHaveRegisteredStream, "test-project", "foo/+/bar")
		})
	})
}

// TestCollector runs through a series of end-to-end Collector workflows and
// ensures that the Collector behaves appropriately.
func TestCollector(t *testing.T) {
	t.Parallel()

	testCollectorImpl(t, false)
}

// TestCollectorWithCaching runs through a series of end-to-end Collector
// workflows and ensures that the Collector behaves appropriately.
func TestCollectorWithCaching(t *testing.T) {
	t.Parallel()

	testCollectorImpl(t, true)
}

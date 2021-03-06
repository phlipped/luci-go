// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package coordinator

import (
	"fmt"

	"github.com/luci/luci-go/common/api/logdog_coordinator/services/v1"
	"github.com/luci/luci-go/common/config"
	"github.com/luci/luci-go/common/errors"
	"github.com/luci/luci-go/common/logdog/types"
	"golang.org/x/net/context"
)

// Coordinator is an interface to a remote LogDog Coordinator service. This is
// a simplified version of the Coordinator's Service API tailored specifically
// to the Collector's usage.
//
// All Coordiantor methods will return transient-wrapped errors if appropriate.
type Coordinator interface {
	// RegisterStream registers a log stream state.
	RegisterStream(c context.Context, s *LogStreamState, desc []byte) (*LogStreamState, error)
	// TerminateStream registers the terminal index of a log stream state.
	TerminateStream(c context.Context, s *LogStreamState) error
}

// LogStreamState is a local representation of a remote stream's state. It is a
// subset of the remote state with the necessary elements for the Collector to
// operate and update.
type LogStreamState struct {
	ID string // The stream's Coordinator ID, populated by the Coordinator.

	Project       config.ProjectName // Project name.
	Path          types.StreamPath   // Stream path.
	ProtoVersion  string             // Stream protocol version string.
	Secret        types.PrefixSecret // Secret.
	TerminalIndex types.MessageIndex // Terminal index, <0 for unterminated.
	Archived      bool               // Is the stream archived?
	Purged        bool               // Is the stream purged?
}

type coordinatorImpl struct {
	c logdog.ServicesClient
}

// NewCoordinator returns a Coordinator implementation that uses a
// logdog.ServicesClient.
func NewCoordinator(s logdog.ServicesClient) Coordinator {
	return &coordinatorImpl{s}
}

func (c *coordinatorImpl) RegisterStream(ctx context.Context, s *LogStreamState, desc []byte) (*LogStreamState, error) {
	// Client-side validate our parameters.
	// TODO(dnj): Force this validation when empty project is not accepted.
	if s.Project != "" {
		if err := s.Project.Validate(); err != nil {
			return nil, fmt.Errorf("failed to validate project: %s", err)
		}
	}
	if err := s.Path.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate path: %s", err)
	}
	if err := s.Secret.Validate(); err != nil {
		return nil, fmt.Errorf("invalid secret: %s", err)
	}

	req := logdog.RegisterStreamRequest{
		Project:      string(s.Project),
		Secret:       []byte(s.Secret),
		ProtoVersion: s.ProtoVersion,
		Desc:         desc,
	}

	resp, err := c.c.RegisterStream(ctx, &req)
	switch {
	case err != nil:
		return nil, err
	case resp.State == nil:
		return nil, errors.New("missing stream state")
	}

	return &LogStreamState{
		ID:            resp.Id,
		Project:       s.Project,
		Path:          s.Path,
		ProtoVersion:  resp.State.ProtoVersion,
		Secret:        types.PrefixSecret(resp.State.Secret),
		TerminalIndex: types.MessageIndex(resp.State.TerminalIndex),
		Archived:      resp.State.Archived,
		Purged:        resp.State.Purged,
	}, nil
}

func (c *coordinatorImpl) TerminateStream(ctx context.Context, s *LogStreamState) error {
	// Client-side validate our parameters.
	// TODO(dnj): Force this validation when empty project is not accepted.
	if s.Project != "" {
		if err := s.Project.Validate(); err != nil {
			return fmt.Errorf("failed to validate project: %s", err)
		}
	}
	if s.ID == "" {
		return errors.New("missing stream ID")
	}
	if s.TerminalIndex < 0 {
		return errors.New("refusing to terminate with non-terminal state")
	}

	req := logdog.TerminateStreamRequest{
		Project:       string(s.Project),
		Id:            s.ID,
		Secret:        []byte(s.Secret),
		TerminalIndex: int64(s.TerminalIndex),
	}

	if _, err := c.c.TerminateStream(ctx, &req); err != nil {
		return err
	}
	return nil
}

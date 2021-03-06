// Code generated by svcdec; DO NOT EDIT

package logdog

import (
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"

	google_protobuf1 "github.com/luci/luci-go/common/proto/google"
)

type DecoratedServices struct {
	// Service is the service to decorate.
	Service ServicesServer
	// Prelude is called in each method before forwarding the call to Service.
	// If Prelude returns an error, it is returned without forwarding the call.
	Prelude func(c context.Context, methodName string, req proto.Message) (context.Context, error)
}

func (s *DecoratedServices) GetConfig(c context.Context, req *google_protobuf1.Empty) (*GetConfigResponse, error) {
	c, err := s.Prelude(c, "GetConfig", req)
	if err != nil {
		return nil, err
	}
	return s.Service.GetConfig(c, req)
}

func (s *DecoratedServices) RegisterStream(c context.Context, req *RegisterStreamRequest) (*RegisterStreamResponse, error) {
	c, err := s.Prelude(c, "RegisterStream", req)
	if err != nil {
		return nil, err
	}
	return s.Service.RegisterStream(c, req)
}

func (s *DecoratedServices) LoadStream(c context.Context, req *LoadStreamRequest) (*LoadStreamResponse, error) {
	c, err := s.Prelude(c, "LoadStream", req)
	if err != nil {
		return nil, err
	}
	return s.Service.LoadStream(c, req)
}

func (s *DecoratedServices) TerminateStream(c context.Context, req *TerminateStreamRequest) (*google_protobuf1.Empty, error) {
	c, err := s.Prelude(c, "TerminateStream", req)
	if err != nil {
		return nil, err
	}
	return s.Service.TerminateStream(c, req)
}

func (s *DecoratedServices) ArchiveStream(c context.Context, req *ArchiveStreamRequest) (*google_protobuf1.Empty, error) {
	c, err := s.Prelude(c, "ArchiveStream", req)
	if err != nil {
		return nil, err
	}
	return s.Service.ArchiveStream(c, req)
}

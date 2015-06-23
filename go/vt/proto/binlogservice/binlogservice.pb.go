// Code generated by protoc-gen-go.
// source: binlogservice.proto
// DO NOT EDIT!

/*
Package binlogservice is a generated protocol buffer package.

It is generated from these files:
	binlogservice.proto

It has these top-level messages:
*/
package binlogservice

import proto "github.com/golang/protobuf/proto"
import binlogdata "github.com/youtube/vitess/go/vt/proto/binlogdata"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

func init() {
}

// Client API for UpdateStream service

type UpdateStreamClient interface {
	// StreamUpdate streams the binlog events, to know which objects have changed.
	StreamUpdate(ctx context.Context, in *binlogdata.StreamUpdateRequest, opts ...grpc.CallOption) (*binlogdata.StreamUpdateResponse, error)
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(ctx context.Context, in *binlogdata.StreamKeyRangeRequest, opts ...grpc.CallOption) (*binlogdata.StreamKeyRangeResponse, error)
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(ctx context.Context, in *binlogdata.StreamTablesRequest, opts ...grpc.CallOption) (*binlogdata.StreamTablesResponse, error)
}

type updateStreamClient struct {
	cc *grpc.ClientConn
}

func NewUpdateStreamClient(cc *grpc.ClientConn) UpdateStreamClient {
	return &updateStreamClient{cc}
}

func (c *updateStreamClient) StreamUpdate(ctx context.Context, in *binlogdata.StreamUpdateRequest, opts ...grpc.CallOption) (*binlogdata.StreamUpdateResponse, error) {
	out := new(binlogdata.StreamUpdateResponse)
	err := grpc.Invoke(ctx, "/binlogservice.UpdateStream/StreamUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateStreamClient) StreamKeyRange(ctx context.Context, in *binlogdata.StreamKeyRangeRequest, opts ...grpc.CallOption) (*binlogdata.StreamKeyRangeResponse, error) {
	out := new(binlogdata.StreamKeyRangeResponse)
	err := grpc.Invoke(ctx, "/binlogservice.UpdateStream/StreamKeyRange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateStreamClient) StreamTables(ctx context.Context, in *binlogdata.StreamTablesRequest, opts ...grpc.CallOption) (*binlogdata.StreamTablesResponse, error) {
	out := new(binlogdata.StreamTablesResponse)
	err := grpc.Invoke(ctx, "/binlogservice.UpdateStream/StreamTables", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UpdateStream service

type UpdateStreamServer interface {
	// StreamUpdate streams the binlog events, to know which objects have changed.
	StreamUpdate(context.Context, *binlogdata.StreamUpdateRequest) (*binlogdata.StreamUpdateResponse, error)
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(context.Context, *binlogdata.StreamKeyRangeRequest) (*binlogdata.StreamKeyRangeResponse, error)
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(context.Context, *binlogdata.StreamTablesRequest) (*binlogdata.StreamTablesResponse, error)
}

func RegisterUpdateStreamServer(s *grpc.Server, srv UpdateStreamServer) {
	s.RegisterService(&_UpdateStream_serviceDesc, srv)
}

func _UpdateStream_StreamUpdate_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(binlogdata.StreamUpdateRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(UpdateStreamServer).StreamUpdate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _UpdateStream_StreamKeyRange_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(binlogdata.StreamKeyRangeRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(UpdateStreamServer).StreamKeyRange(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _UpdateStream_StreamTables_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(binlogdata.StreamTablesRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(UpdateStreamServer).StreamTables(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _UpdateStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binlogservice.UpdateStream",
	HandlerType: (*UpdateStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StreamUpdate",
			Handler:    _UpdateStream_StreamUpdate_Handler,
		},
		{
			MethodName: "StreamKeyRange",
			Handler:    _UpdateStream_StreamKeyRange_Handler,
		},
		{
			MethodName: "StreamTables",
			Handler:    _UpdateStream_StreamTables_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
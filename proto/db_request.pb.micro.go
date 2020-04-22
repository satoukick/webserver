// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: db_request.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DBToDoQuery service

type DBToDoQueryService interface {
	ToDoQuery(ctx context.Context, in *ToDoQueryRequest, opts ...client.CallOption) (*ToDoQueryResponse, error)
}

type dBToDoQueryService struct {
	c    client.Client
	name string
}

func NewDBToDoQueryService(name string, c client.Client) DBToDoQueryService {
	return &dBToDoQueryService{
		c:    c,
		name: name,
	}
}

func (c *dBToDoQueryService) ToDoQuery(ctx context.Context, in *ToDoQueryRequest, opts ...client.CallOption) (*ToDoQueryResponse, error) {
	req := c.c.NewRequest(c.name, "DBToDoQuery.ToDoQuery", in)
	out := new(ToDoQueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DBToDoQuery service

type DBToDoQueryHandler interface {
	ToDoQuery(context.Context, *ToDoQueryRequest, *ToDoQueryResponse) error
}

func RegisterDBToDoQueryHandler(s server.Server, hdlr DBToDoQueryHandler, opts ...server.HandlerOption) error {
	type dBToDoQuery interface {
		ToDoQuery(ctx context.Context, in *ToDoQueryRequest, out *ToDoQueryResponse) error
	}
	type DBToDoQuery struct {
		dBToDoQuery
	}
	h := &dBToDoQueryHandler{hdlr}
	return s.Handle(s.NewHandler(&DBToDoQuery{h}, opts...))
}

type dBToDoQueryHandler struct {
	DBToDoQueryHandler
}

func (h *dBToDoQueryHandler) ToDoQuery(ctx context.Context, in *ToDoQueryRequest, out *ToDoQueryResponse) error {
	return h.DBToDoQueryHandler.ToDoQuery(ctx, in, out)
}
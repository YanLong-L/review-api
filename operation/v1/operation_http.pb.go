// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v4.25.3
// source: api/operation/v1/operation.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationOperationAuditAppeal = "/api.operation.v1.Operation/AuditAppeal"
const OperationOperationAuditReview = "/api.operation.v1.Operation/AuditReview"

type OperationHTTPServer interface {
	AuditAppeal(context.Context, *AuditAppealRequest) (*AuditAppealReply, error)
	AuditReview(context.Context, *AuditReviewRequest) (*AuditReviewReply, error)
}

func RegisterOperationHTTPServer(s *http.Server, srv OperationHTTPServer) {
	r := s.Route("/")
	r.POST("o/v1/review/audit", _Operation_AuditReview0_HTTP_Handler(srv))
	r.POST("o/v1/appeal/audit", _Operation_AuditAppeal0_HTTP_Handler(srv))
}

func _Operation_AuditReview0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuditReviewRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationAuditReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuditReview(ctx, req.(*AuditReviewRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuditReviewReply)
		return ctx.Result(200, reply)
	}
}

func _Operation_AuditAppeal0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuditAppealRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationAuditAppeal)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuditAppeal(ctx, req.(*AuditAppealRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuditAppealReply)
		return ctx.Result(200, reply)
	}
}

type OperationHTTPClient interface {
	AuditAppeal(ctx context.Context, req *AuditAppealRequest, opts ...http.CallOption) (rsp *AuditAppealReply, err error)
	AuditReview(ctx context.Context, req *AuditReviewRequest, opts ...http.CallOption) (rsp *AuditReviewReply, err error)
}

type OperationHTTPClientImpl struct {
	cc *http.Client
}

func NewOperationHTTPClient(client *http.Client) OperationHTTPClient {
	return &OperationHTTPClientImpl{client}
}

func (c *OperationHTTPClientImpl) AuditAppeal(ctx context.Context, in *AuditAppealRequest, opts ...http.CallOption) (*AuditAppealReply, error) {
	var out AuditAppealReply
	pattern := "o/v1/appeal/audit"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperationAuditAppeal))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperationHTTPClientImpl) AuditReview(ctx context.Context, in *AuditReviewRequest, opts ...http.CallOption) (*AuditReviewReply, error) {
	var out AuditReviewReply
	pattern := "o/v1/review/audit"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperationAuditReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

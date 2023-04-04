// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package video

import (
	"context"

	"giligili/app/video/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BarrageInfo         = pb.BarrageInfo
	DeleteVideoReq      = pb.DeleteVideoReq
	DeleteVideoResp     = pb.DeleteVideoResp
	DislikeVideoReq     = pb.DislikeVideoReq
	DislikeVideoResp    = pb.DislikeVideoResp
	GetBarrageListReq   = pb.GetBarrageListReq
	GetBarrageListResp  = pb.GetBarrageListResp
	GetHotVideoListReq  = pb.GetHotVideoListReq
	GetHotVideoListResp = pb.GetHotVideoListResp
	GetVideoDetailReq   = pb.GetVideoDetailReq
	GetVideoDetailResp  = pb.GetVideoDetailResp
	GetVideoListReq     = pb.GetVideoListReq
	GetVideoListResp    = pb.GetVideoListResp
	HotVideoInfo        = pb.HotVideoInfo
	IsExistVideoReq     = pb.IsExistVideoReq
	IsExistVideoResp    = pb.IsExistVideoResp
	LikeVideoReq        = pb.LikeVideoReq
	LikeVideoResp       = pb.LikeVideoResp
	UndislikeVideoReq   = pb.UndislikeVideoReq
	UndislikeVideoResp  = pb.UndislikeVideoResp
	UnlikeVideoReq      = pb.UnlikeVideoReq
	UnlikeVideoResp     = pb.UnlikeVideoResp
	UploadVideoReq      = pb.UploadVideoReq
	UploadVideoResp     = pb.UploadVideoResp
	VideoInfo           = pb.VideoInfo

	Video interface {
		GetVideoList(ctx context.Context, in *GetVideoListReq, opts ...grpc.CallOption) (*GetVideoListResp, error)
		GetVideoDetail(ctx context.Context, in *GetVideoDetailReq, opts ...grpc.CallOption) (*GetVideoDetailResp, error)
		UploadVideo(ctx context.Context, in *UploadVideoReq, opts ...grpc.CallOption) (*UploadVideoResp, error)
		DeleteVideo(ctx context.Context, in *DeleteVideoReq, opts ...grpc.CallOption) (*DeleteVideoResp, error)
		LikeVideo(ctx context.Context, in *LikeVideoReq, opts ...grpc.CallOption) (*LikeVideoResp, error)
		UnlikeVideo(ctx context.Context, in *UnlikeVideoReq, opts ...grpc.CallOption) (*UnlikeVideoResp, error)
		DislikeVideo(ctx context.Context, in *DislikeVideoReq, opts ...grpc.CallOption) (*DislikeVideoResp, error)
		UndislikeVideo(ctx context.Context, in *UndislikeVideoReq, opts ...grpc.CallOption) (*UndislikeVideoResp, error)
		GetBarrageList(ctx context.Context, in *GetBarrageListReq, opts ...grpc.CallOption) (*GetBarrageListResp, error)
		GetHotVideoList(ctx context.Context, in *GetHotVideoListReq, opts ...grpc.CallOption) (*GetHotVideoListResp, error)
		IsExistVideo(ctx context.Context, in *IsExistVideoReq, opts ...grpc.CallOption) (*IsExistVideoResp, error)
	}

	defaultVideo struct {
		cli zrpc.Client
	}
)

func NewVideo(cli zrpc.Client) Video {
	return &defaultVideo{
		cli: cli,
	}
}

func (m *defaultVideo) GetVideoList(ctx context.Context, in *GetVideoListReq, opts ...grpc.CallOption) (*GetVideoListResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.GetVideoList(ctx, in, opts...)
}

func (m *defaultVideo) GetVideoDetail(ctx context.Context, in *GetVideoDetailReq, opts ...grpc.CallOption) (*GetVideoDetailResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.GetVideoDetail(ctx, in, opts...)
}

func (m *defaultVideo) UploadVideo(ctx context.Context, in *UploadVideoReq, opts ...grpc.CallOption) (*UploadVideoResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.UploadVideo(ctx, in, opts...)
}

func (m *defaultVideo) DeleteVideo(ctx context.Context, in *DeleteVideoReq, opts ...grpc.CallOption) (*DeleteVideoResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.DeleteVideo(ctx, in, opts...)
}

func (m *defaultVideo) LikeVideo(ctx context.Context, in *LikeVideoReq, opts ...grpc.CallOption) (*LikeVideoResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.LikeVideo(ctx, in, opts...)
}

func (m *defaultVideo) UnlikeVideo(ctx context.Context, in *UnlikeVideoReq, opts ...grpc.CallOption) (*UnlikeVideoResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.UnlikeVideo(ctx, in, opts...)
}

func (m *defaultVideo) DislikeVideo(ctx context.Context, in *DislikeVideoReq, opts ...grpc.CallOption) (*DislikeVideoResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.DislikeVideo(ctx, in, opts...)
}

func (m *defaultVideo) UndislikeVideo(ctx context.Context, in *UndislikeVideoReq, opts ...grpc.CallOption) (*UndislikeVideoResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.UndislikeVideo(ctx, in, opts...)
}

func (m *defaultVideo) GetBarrageList(ctx context.Context, in *GetBarrageListReq, opts ...grpc.CallOption) (*GetBarrageListResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.GetBarrageList(ctx, in, opts...)
}

func (m *defaultVideo) GetHotVideoList(ctx context.Context, in *GetHotVideoListReq, opts ...grpc.CallOption) (*GetHotVideoListResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.GetHotVideoList(ctx, in, opts...)
}

func (m *defaultVideo) IsExistVideo(ctx context.Context, in *IsExistVideoReq, opts ...grpc.CallOption) (*IsExistVideoResp, error) {
	client := pb.NewVideoClient(m.cli.Conn())
	return client.IsExistVideo(ctx, in, opts...)
}

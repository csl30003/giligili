package video

import (
	"context"
	"errors"
	"giligili/app/video/api/internal/svc"
	"giligili/app/video/api/internal/types"
	"giligili/app/video/rpc/video"
	"giligili/common/ctxData"
	"io"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/rand"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	uploadPath = "../file"
)

type UploadVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
	File   map[string][]*multipart.FileHeader
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadVideoLogic {
	return &UploadVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

// UploadVideo 上传视频，可以考虑用消息队列给消费者干，这样文件上传失败怎么办？
func (l *UploadVideoLogic) UploadVideo(req *types.UploadVideoReq) (resp *types.UploadVideoResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)

	localFileName := uploadPath
	file, fileHeader, err := l.r.FormFile("file")
	if err != nil {
		logx.Error("文件打开失败", err)
		return nil, err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			logx.Error("文件关闭失败", err)
		}
	}(file)

	// 判断文件类型
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logx.Error("以字节读文件失败", err)
		return nil, err
	}
	filetype := http.DetectContentType(fileBytes)
	if filetype != "video/avi" && filetype != "video/mp4" && filetype != "audio/mpeg" {
		return nil, errors.New("文件格式错误")
	}

	// 创建本地文件
	localFileName += "/" + rand.String(7) + fileHeader.Filename
	out, err := os.Create(localFileName)
	if err != nil {
		logx.Error("本地文件创建失败", err)
		return nil, err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			logx.Error("本地文件关闭失败", err)
		}
	}(out)

	// 拷贝文件
	_, err = io.Copy(out, file)
	if err != nil {
		logx.Error("文件拷贝失败", err)
		return nil, err
	}

	logx.Info(fileHeader.Filename, "文件上传成功")

	uploadVideoResp, err := l.svcCtx.VideoRpcClient.UploadVideo(l.ctx, &video.UploadVideoReq{
		UserId:      userId,
		Title:       req.Title,
		Description: req.Description,
		Url:         localFileName,
	})
	if err != nil {
		return nil, err
	}

	return &types.UploadVideoResp{Success: uploadVideoResp.Success}, nil
}

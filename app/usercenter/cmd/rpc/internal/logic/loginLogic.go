package logic

import (
	"context"

	"XavierLookLook/app/usercenter/cmd/rpc/internal/svc"
	"XavierLookLook/app/usercenter/cmd/rpc/pb"
	"XavierLookLook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginResp{}, nil
}

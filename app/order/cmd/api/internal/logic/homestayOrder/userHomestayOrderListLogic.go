package homestayOrder

import (
	"context"

	"XavierLookLook/app/order/cmd/api/internal/svc"
	"XavierLookLook/app/order/cmd/api/internal/types"
	"XavierLookLook/app/order/cmd/rpc/order"
	"XavierLookLook/common/ctxdata"
	"XavierLookLook/common/tool"
	"XavierLookLook/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户订单列表
func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderListLogic) UserHomestayOrderList(req *types.UserHomestayOrderListReq) (resp *types.UserHomestayOrderListResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx) //get login user id
	OrderRpc, err := l.svcCtx.OrderRpc.UserHomestayOrderList(l.ctx, &order.UserHomestayOrderListReq{
		UserId:      userId,
		TraderState: req.TradeState,
		PageSize:    req.PageSize,
		LastId:      req.LastId,
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get user homestay order list"), "Failed to get user homestay order list err : %v ,req:%+v", err, req)
	}

	var typesUserHomestayOrderList []types.UserHomestayOrderListView

	if len(OrderRpc.List) > 0 {

		for _, homestayOrder := range OrderRpc.List {

			var typeHomestayOrder types.UserHomestayOrderListView
			_ = copier.Copy(&typeHomestayOrder, homestayOrder)

			typeHomestayOrder.OrderTotalPrice = tool.Fen2Yuan(int64(homestayOrder.OrderTotalPrice))

			typesUserHomestayOrderList = append(typesUserHomestayOrderList, typeHomestayOrder)
		}
	}

	return &types.UserHomestayOrderListResp{
		List: typesUserHomestayOrderList,
	}, nil
}

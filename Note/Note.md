# model 小tricks

```go
// 民宿详情
func (l *HomestayDetailLogic) HomestayDetail(in *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	// todo: add your logic here and delete this line
	homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), " HomestayDetail db err , id : %d ", in.Id)
	}

	var pbHomestay pb.Homestay
	if homestay != nil {
		_ = copier.Copy(&pbHomestay, homestay)
	}

	return &pb.HomestayDetailResp{
		Homestay: &pbHomestay,
	}, nil
}


//在这里为什么我们不去findlist，是因为我们在findone方法中有缓存，我们一个个根据id查询数据时候，
//只有第一次会命中db，其他时间基本都是命中的redis cache，这样不仅速度快，
//就算流量激增的时候，也不会全部打到db上，而是都在redis上，这样会大大提高我们系统的访问速度以及db支撑能力。

func (m *defaultHomestayModel) FindOne(ctx context.Context, id int64) (*Homestay, error) {
	looklookTravelHomestayIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayIdPrefix, id)
	var resp Homestay
	err := m.QueryRowCtx(ctx, &resp, looklookTravelHomestayIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

```
**go-zero使用了配套内置工具goctl生成的model，自带sqlc+sqlx实现的代码，实现了自动缓存管理，我们根本不需要去管理缓存，只需要用sqlx写 sql数据，sqlc会自动帮我们管理缓存，并且是通过singleflight ,也就是说即使缓存在某个时间失效，在失效那一刻同时有大量并发请求进来时，go-zero在查询db时候也只会放行一个线程进来，其他线程是在等待，当这个线程从数据库拿数据回来之后将该数据缓存到redis同时所有之前等待线程共享此数据返回，后续在进来的线程查相同数据时，就只会进入到redis中而不会进入到db。**
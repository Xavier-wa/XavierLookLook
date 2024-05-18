package logic

import (
	"XavierLookLook/app/mqueue/cmd/job/internal/svc"
	"XavierLookLook/app/mqueue/cmd/job/jobtype"
	"context"

	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register job
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	//scheduler job
	mux.Handle(jobtype.ScheduleSettleRecord, NewSettleRecordHandler(l.svcCtx))
	//defer job
	mux.Handle(jobtype.DeferCloseHomestayOrder, NewCloseHomestayOrderHandler(l.svcCtx))

	//queue job , asynq support queue job //调用小程序通知用户
	mux.Handle(jobtype.MsgPaySuccessNotifyUser, NewPaySuccessNotifyUserHandler(l.svcCtx))
	// wait you fill..

	return mux
}

package logic

import (
	"context"

	"dl/api/internal/svc"
	"dl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDLLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDLLogic {
	return &CreateDLLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDLLogic) CreateDL(req *types.DLCreateRequest) (resp *types.DLCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

package logic

import (
	"context"
	"iot-platform/helper"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceListLogic {
	return &DeviceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceListLogic) DeviceList(req *types.DeviceListRequest) (resp *types.DeviceListResponse, err error) {
	// int() 是标准类型转换
	// .(int) 是类型断言的转换
	req.Size = helper.If(req.Size == 0, 20, req.Size).(int)
	req.Page = helper.If(req.Page == 0, 0, (req.Page-1)*req.Size).(int)
	var count int64
	resp = new(types.DeviceListResponse)

	data := make([]*types.DeviceListBaisc, 0)
	err = models.GetDeviceList(req.Name).Count(&count).Limit(req.Size).Offset(req.Page).Find(&data).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}
	resp.Count = count
	resp.List = data

	return
}

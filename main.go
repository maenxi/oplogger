package main

import (
	"context"
	"fmt"
	oplog "team.wphr.vip/technology-group/sdk/sdk-biz-oplog"
	sdk "team.wphr.vip/technology-group/sdk/sdk-go-grpc-erpoploggo"
	"time"
)

func main() {
	// init 使用默认配置可不传
	oplog.Init(oplog.LoggerOption{}, oplog.GrpcTargetOption{
		Target: "localhost:8990",
	})

	ctx := context.Background()
	opl := oplog.New()

	req := sdk.LoggerAddReq{
		AppId:         1,
		BizName:       "biz_name",
		BizKey:        "1",
		BizBatchKey:   "1",
		BizType:       10010,
		Operator:      "操作人",
		OpType:        1,
		OldData:       "{\"j\":\"x\", \"x\":\"xx\", \"yy\":\"yy\"}",
		NewData:       "{\"j\":\"x\", \"x\":\"xx\", \"yy\":\"yy\"}",
		ExtraInfo:     "{\"j\":\"x\", \"x\":\"xx\", \"yy\":\"yy\"}",
		Remark:        "备注",
		RequestParams: "{\"j\":\"x\", \"x\":\"xx\", \"yy\":\"yy\"}",
	}

	for true {
		aRsp, err := opl.AddLogger(ctx, &req)
		fmt.Printf("单条插入 准同步=%+v||err=%v\n", aRsp, err)

		req1 := req
		req2 := req
		req1.BizType = 10011
		req2.BizType = 10012

		batchReq := &sdk.LoggerBatchAddReq{
			LoggerAddData: []*sdk.LoggerAddReq{
				&req,
				&req1,
				&req2,
			},
		}
		baRsp, err := opl.BatchAddLogger(ctx, batchReq)
		fmt.Printf("批量插入 准同步=%+v||err=%v\n", baRsp, err)

		time.Sleep(time.Second * 3)
	}
}

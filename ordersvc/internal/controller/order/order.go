package order

import (
	"base_util/common"
	"context"
	"ordersvc/internal/entity/vo"
	"ordersvc/internal/grcpclient"
	"protos"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/metadata"
)

func CreateOrder(r *ghttp.Request) {

	span := opentracing.GlobalTracer().StartSpan("CreateOrder")
	defer span.Finish()

	req := &vo.OrderCreateReq{}
	err := common.ParseJsonReq(r, req)
	if err != nil {
		r.SetError(err)
		return
	}

	traceId := uuid.New().String()
	span.SetBaggageItem("trace_id", traceId)
	span.SetBaggageItem("user_id", strconv.FormatInt(req.UserID, 10))

	ctx := opentracing.ContextWithSpan(r.GetCtx(), span)
	//将用户ID和traceID放入metadata传递到下游grpc服务
	userId := strconv.FormatInt(req.UserID, 10)
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", userId)
	ctx = metadata.AppendToOutgoingContext(ctx, "trace_id", traceId)

	//将tracer的span传递到下游grpc服务

	resp, err := grcpclient.GetUserClient().GetUsersInfo(ctx, &protos.User{
		Id: 2,
	})
	if err != nil {
		r.SetError(err)
		return
	}

	someLocalFunction(ctx, span)

	g.Log().Infof(r.GetCtx(), "username: %s", resp.Name)

	common.Success(r, &vo.OrderCreateResp{})
}

func someLocalFunction(ctx context.Context, upperSpan opentracing.Span) {
	// do something
	span := opentracing.GlobalTracer().StartSpan("someLocalFunction", opentracing.ChildOf(upperSpan.Context()))
	defer span.Finish()

	time.Sleep(time.Millisecond * 100)
}

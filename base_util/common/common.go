package common

import (
	"base_util/uerrors"
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

type DefaultHttpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Fail2(r *ghttp.Request, err *uerrors.VideoTranslateErr) {
	r.Response.Status = http.StatusUnprocessableEntity
	r.Response.WriteJson(DefaultHttpResponse{
		Code: err.RetCode,
		Msg:  err.Msg,
		Data: struct {
		}{},
	})
	r.Response.WriteExit()
}

func Success(r *ghttp.Request, v any) {
	r.Response.WriteJson(DefaultHttpResponse{
		Code: 0,
		Msg:  "ok",
		Data: v,
	})
	r.Response.WriteExit()
}

func ParseJsonReq(r *ghttp.Request, v any) error {

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Request.Body = io.NopCloser(bytes.NewBuffer(data)) // 恢复请求体
	err = json.Unmarshal(data, v)

	if err != nil {
		return err
	}

	return nil
}

package util

import (
	"github.com/astaxie/beego/context"
)

type ApiError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code int, msg... string) ApiError {
	err := ApiError{}
	err.Code = code
	if len(msg) > 0 {
		err.Msg = msg[0]
	} else {
		err.Msg = GetErrorMsg(code)
	}
	return err
}

type ErrorHandlerFunc func(ctx *context.Context)

var ErrHandle ErrorHandlerFunc

func init() {
	ErrHandle = json_err_handle_func
}

func json_err_handle_func(ctx *context.Context) {
	err := recover()

	if err != nil {
		ctx.Output.Status = 210
		if api_err, ok := err.(ApiError); ok {
			ctx.Output.JSON(api_err, false, false)
		} else if unknown_err, ok := err.(error); ok {
			json_err := ApiError{
				Code: 500,
				Msg:  unknown_err.Error(),
			}
			ctx.Output.JSON(json_err, false, false)
		} else {
			json_err := ApiError{
				Code: 500,
				Msg:  unknown_err.Error(),
			}
			ctx.Output.JSON(json_err, false, false)
		}
	}
}

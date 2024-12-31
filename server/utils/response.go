package utils

import (
	"github.com/savsgio/atreugo/v11"
)

func ErrorResponse(ctx *atreugo.RequestCtx, message string, err error) error {
	return ctx.JSONResponse(map[string]interface{}{
		"successful" : false,
		"message" : message + ": " + err.Error(),
	}, 500)
}

func BadRespone(ctx *atreugo.RequestCtx, message string) error {
	return ctx.JSONResponse(map[string]interface{}{
		"successful" : false,
		"message" : message,
	}, 400)
}

func UnauthorizedResponse(ctx *atreugo.RequestCtx) error {
	return ctx.JSONResponse(map[string]interface{}{
		"successful" : false,
		"message" : "Unauthorized",
	}, 401)
}

func OkResponse(ctx *atreugo.RequestCtx, message interface{}) error {
	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message" : message,
	}, 200)
}

func OkPlusResponse(ctx *atreugo.RequestCtx, message string, data interface{}, key string) error {
	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message" : message,
		key : data,
	}, 200)
}
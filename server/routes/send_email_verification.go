package routes

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func SendEmailVerification(ctx * atreugo.RequestCtx) error {
	client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Unauthorized",
		}, 401)
	}

	account := appwrite.NewAccount(client)
	
	response, err := account.CreateVerification(
		"http://localhost:2050/verify_email",
	)
	
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to send verification email: " + err.Error(),
		}, 500)
	}

	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message": response,
	}, 200)
}
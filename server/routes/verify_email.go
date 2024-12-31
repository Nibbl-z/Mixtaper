package routes

import ( 
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func VerifyEmail(ctx *atreugo.RequestCtx) error {
	userId := string(ctx.QueryArgs().Peek("userId"))
	secret := string(ctx.QueryArgs().Peek("secret"))

	client, success := utils.CreateClientWithHeaders(ctx)

	if !success {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Unauthorized",
		}, 401)
	}

	account := appwrite.NewAccount(client)
	
	_, err := account.UpdateVerification(userId, secret)

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to verify user: " + err.Error(),
		}, 500)
	}
	
	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message": "Verified successfully!",
	}, 200)
}
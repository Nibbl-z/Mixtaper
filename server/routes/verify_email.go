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
		return utils.UnauthorizedResponse(ctx)
	}

	account := appwrite.NewAccount(client)
	
	_, err := account.UpdateVerification(userId, secret)

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to verify user", err)
	}
	
	return utils.OkResponse(ctx, "Verified successfully!")
}
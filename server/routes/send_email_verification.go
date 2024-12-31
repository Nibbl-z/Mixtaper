package routes

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func SendEmailVerification(ctx * atreugo.RequestCtx) error {
	client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return utils.UnauthorizedResponse(ctx)
	}

	account := appwrite.NewAccount(client)
	
	response, err := account.CreateVerification(
		"http://localhost:2050/verify_email",
	)
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to send verification email", err)
	}
	
	return utils.OkResponse(ctx, response)
}
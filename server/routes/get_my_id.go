package routes

import (
	"server/utils"
	
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/savsgio/atreugo/v11"
)

func GetMyId(ctx *atreugo.RequestCtx) error {
	client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return utils.UnauthorizedResponse(ctx)
	}

	account := appwrite.NewAccount(client)
	user, err := account.Get()

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to get user", err)
	}

	return utils.OkResponse(ctx, user.Id)
}
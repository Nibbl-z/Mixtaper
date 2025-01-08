package routes

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func ChangeBio(ctx *atreugo.RequestCtx) error {
	bio := ctx.Request.Body()

	if len(bio) > 1000 {
		return utils.BadRespone(ctx, "Username must be less than 1000 characters!")
	}

	client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return utils.UnauthorizedResponse(ctx)
	}

	account := appwrite.NewAccount(client)
	
	_, err := account.UpdatePrefs(
		map[string]interface{}{
			"bio" : string(bio),
		},
	)
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to update bio", err)
	}

	return utils.OkResponse(ctx, "Updated bio successfully!")
}
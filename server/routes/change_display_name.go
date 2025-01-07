package routes

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func ChangeDisplayName(ctx *atreugo.RequestCtx) error {
	displayName := ctx.Request.Body()

	if len(displayName) > 20 {
		return utils.BadRespone(ctx, "Display name must be less than 20 characters!")
	}
	
	if len(displayName) < 3 {
		return utils.BadRespone(ctx, "Display name must be at least 3 characters long!")
	}

	client := utils.CreateClient()
	user_client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return utils.UnauthorizedResponse(ctx)
	}
	
	user, err := appwrite.NewAccount(user_client).Get()
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to get user", err)
	}

	users := appwrite.NewUsers(client)

	_, err = users.UpdateName(user.Id, string(displayName))

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to update display name", err)
	}

	return utils.OkResponse(ctx, "Updated display name successfully!")
}
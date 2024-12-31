package routes

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func ChangeUsername(ctx *atreugo.RequestCtx) error {
	username := ctx.Request.Body()

	if len(username) > 20 {
		return utils.BadRespone(ctx, "Username must be less than 20 characters!")
	}

	if len(username) < 3 {
		return utils.BadRespone(ctx, "Username must be at least 3 characters long!")
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
	
	database := appwrite.NewDatabases(client)
	
	_, err = database.UpdateDocument(
		"mixtaper",
		"usernames",
		user.Id,
		database.WithUpdateDocumentData(map[string]string{
			"username" : string(username),
		}),
	)
	
	if err != nil {
		if err.Error() == "Document with the requested ID already exists. Try again with a different ID or use ID.unique() to generate a unique ID." {
			return utils.BadRespone(ctx, "Username is already taken!")
		} else {
			return utils.ErrorResponse(ctx, "Failed to update username", err)
		}
	}

	return utils.OkResponse(ctx, "Username updated successfully!")
}
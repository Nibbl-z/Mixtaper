package routes

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func ChangeUsername(ctx *atreugo.RequestCtx) error {
	username := ctx.Request.Body()

	if len(username) > 20 {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Username must be less than 20 characters!",
		}, 400)
	}
	
	if len(username) < 3 {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Username must be at least 3 characters long!",
		}, 400)
	}
	
	client := utils.CreateClient()
	user_client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Unauthorized",
		}, 401)
	}
	
	user, err := appwrite.NewAccount(user_client).Get()
	
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to get user",
		}, 500)
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
			return ctx.JSONResponse(map[string]interface{}{
				"successful" : false,
				"message": "Username is already taken!",
			}, 400)
		} else {
			return ctx.JSONResponse(map[string]interface{}{
				"successful" : false,
				"message": "Failed to update username: " + err.Error(),
			}, 500)
		}

		
	}

	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message": "Username updated successfully!",
	}, 200)
}
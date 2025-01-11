package routes

import (
	"server/utils"
	
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/savsgio/atreugo/v11"
)

func GetUser(ctx *atreugo.RequestCtx) error {
	id := string(ctx.QueryArgs().Peek("id"))

	if string(id) == "" {
		return utils.BadRespone(ctx, "Level ID missing")
	}
	
	client := utils.CreateClient()
	
	users := appwrite.NewUsers(client)
	
	user, err := users.Get(id)
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to get user", err)
	}
	
	username, err := utils.GetUsernameByID(&client, id)
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to get username", err)
	}

	prefs, err := users.GetPrefs(id)

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to get bio", err)
	}

	var preferences utils.Prefs
	prefs.Decode(&preferences)
	

	return utils.OkResponse(ctx, utils.User{
		Username: username,
		DisplayName: user.Name,
		ID: user.Id,
		Bio: preferences.Bio,
		Verified: user.EmailVerification,
	})
}
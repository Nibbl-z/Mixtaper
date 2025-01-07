package routes

import (
	"encoding/json"
	"server/utils"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/savsgio/atreugo/v11"
)

type ChangePasswordData struct {
	PreviousPassword string `json:"previous"`
	NewPassword string `json:"new"`
}

func ChangePassword(ctx *atreugo.RequestCtx) error {
	var changePasswordData ChangePasswordData
	
	if err := json.Unmarshal(ctx.Request.Body(), &changePasswordData); err != nil {
		return utils.BadRespone(ctx, "Signup data was not provided correctly")
	}
	
	if len(changePasswordData.NewPassword) < 8 {
		return utils.BadRespone(ctx, "Password must be at least 8 characters long!")
	}
	
	if len(changePasswordData.NewPassword) > 64 {
		return utils.BadRespone(ctx, "Password must be less than 64 characters!")
	}

	client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return utils.UnauthorizedResponse(ctx)
	}
	
	accounts := appwrite.NewAccount(client)

	_, err := accounts.UpdatePassword(changePasswordData.NewPassword, accounts.WithUpdatePasswordOldPassword(changePasswordData.PreviousPassword))

	if err != nil {
		return utils.BadRespone(ctx, "Incorrect previous password")
	}

	return utils.OkResponse(ctx, "Password changed successfully!")
}
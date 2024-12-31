package routes

import (
	"encoding/json"
	"regexp"

	"github.com/appwrite/sdk-for-go/account"
	"github.com/savsgio/atreugo/v11"

	"server/utils"
)

type LoginData struct {
	Identifier string `json:"identifier"`
	Password string `json:"password"`
}

func Login(ctx *atreugo.RequestCtx) error {
	var loginRequest LoginData
	
	if err := json.Unmarshal(ctx.Request.Body(), &loginRequest); err != nil {
		return utils.BadRespone(ctx, "Login data was not provided correctly")
	}
	
	client := utils.CreateClient()

	var email string
	
	isEmail, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, loginRequest.Identifier)
		
	if err != nil {
		return utils.ErrorResponse(ctx, "Email regex failed", err)
	}

	if isEmail {
		email = loginRequest.Identifier
	} else {
		user, code := utils.GetUserByUsername(&client, loginRequest.Identifier)

		if code == 500 {
			return utils.ErrorResponse(ctx, "Failed to fetch user from database", err)
		} else if code == 400 {
			return utils.BadRespone(ctx, "Username was not found in the database!")
		}

		email = user.Email
	}
	
	account_service := account.New(client)

	session, err := account_service.CreateEmailPasswordSession(
		email,
		loginRequest.Password,
	)

	if err != nil {
		return utils.BadRespone(ctx, "Invalid credentials")
	}

	return utils.OkPlusResponse(ctx, "Logged in successfully!", session.Secret, "secret")
}
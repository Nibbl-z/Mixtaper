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
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Login data was not provided correctly",
		}, 400)
	}
	
	client := utils.CreateClient()

	var email string
	
	isEmail, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, loginRequest.Identifier)
		
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Email regex failed",
		}, 500)
	}

	if isEmail {
		email = loginRequest.Identifier
	} else {
		user, code := utils.GetUserByUsername(&client, loginRequest.Identifier)

		if code == 500 {
			return ctx.JSONResponse(map[string]interface{}{
				"successful" : false,
				"message": "Failed to fetch user from database",
			}, 500)
		} else if code == 400 {
			return ctx.JSONResponse(map[string]interface{}{
				"successful" : false,
				"message": "Username was not found in the database!",
			}, 400)
		}

		email = user.Email
	}
	
	account_service := account.New(client)

	session, err := account_service.CreateEmailPasswordSession(
		email,
		loginRequest.Password,
	)

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Invalid credentials",
		}, 400)
	}

	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message": "Logged in successfully!",
		"secret": session.Secret,
	}, 400)
}
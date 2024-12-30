package routes

import (
	"encoding/json"
	"os"
	"errors"
	"regexp"

	"github.com/appwrite/sdk-for-go/appwrite"
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
		return ctx.ErrorResponse(errors.New("signup data not provided correctly"), 400)
	}

	client := appwrite.NewClient(
		appwrite.WithEndpoint(os.Getenv("APPWRITE_API_ENDPOINT")),
		appwrite.WithProject(os.Getenv("APPWRITE_PROJECT_ID")),
		appwrite.WithKey(os.Getenv("APPWRITE_API_KEY")),
	)

	var email string
	
	isEmail, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, loginRequest.Identifier)
		
	if err != nil {
		return ctx.ErrorResponse(err, 500)
	}

	if isEmail {
		email = loginRequest.Identifier
	} else {
		user, err := utils.GetUserByUsername(&client, loginRequest.Identifier)
		if err != nil {
			return ctx.ErrorResponse(err, 500)
		}

		email = user.Email
	}
	
	account_service := account.New(client)

	session, err := account_service.CreateEmailPasswordSession(
		email,
		loginRequest.Password,
	)

	if err != nil {
		return ctx.ErrorResponse(err, 500)
	}

	return ctx.TextResponse(session.Secret)
}
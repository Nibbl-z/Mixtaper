package routes

import (
	"encoding/json"
	"os"
	"errors"
	
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/users"
	"github.com/savsgio/atreugo/v11"
	
	"golang.org/x/crypto/bcrypt"
)

type SignupData struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(ctx *atreugo.RequestCtx) error {
	var signupRequest SignupData
	
	if err := json.Unmarshal(ctx.Request.Body(), &signupRequest); err != nil {
		return ctx.ErrorResponse(errors.New("signup data not provided correctly"), 400)
	}
	
	client := appwrite.NewClient(
		appwrite.WithEndpoint(os.Getenv("APPWRITE_API_ENDPOINT")),
		appwrite.WithProject(os.Getenv("APPWRITE_PROJECT_ID")),
		appwrite.WithKey(os.Getenv("APPWRITE_API_KEY")),
	)
	
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
	
	if err != nil {
		return ctx.ErrorResponse(errors.New("password failed to hash"), 500)
	}
	
	users_service := users.New(client)
	
	response, err := users_service.CreateBcryptUser(
		id.Unique(),
		signupRequest.Email,
		string(hashed_password),
		users_service.WithCreateBcryptUserName(signupRequest.Username),
	)
	
	if err != nil {
		return ctx.ErrorResponse(err, 500)
	}

	return ctx.TextResponse(response.Name + " has signed up successfully!")
}
package routes

import (
	"encoding/json"
	"server/utils"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/query"

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
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Signup data was not provided correctly",
		}, 400)
	}
	
	client := utils.CreateClient()

	database := appwrite.NewDatabases(client)

	list, err := database.ListDocuments("mixtaper", "usernames", database.WithListDocumentsQueries(
		[]string{query.Equal("username", signupRequest.Username)},
	))

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to fetch used usernames",
		}, 500)
	}

	if list.Total > 0 {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Username is taken",
		}, 400)
	}
	
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
	
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to hash password",
		}, 500)
	}
	
	users := appwrite.NewUsers(client)
	
	user, err := users.CreateBcryptUser(
		id.Unique(),
		signupRequest.Email,
		string(hashed_password),
		users.WithCreateBcryptUserName(signupRequest.Username),
	)
	
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to create user",
		}, 500)
	}

	_, err = database.CreateDocument(
		"mixtaper",
		"usernames",
		user.Id,
		map[string]string{
			"username" : signupRequest.Username,
		},
	)

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to add username to database",
		}, 500)
	}

	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message": user.Name + " has signed up successfully!",
	}, 200)
}
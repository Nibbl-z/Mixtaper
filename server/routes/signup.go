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
		return utils.BadRespone(ctx, "Signup data was not provided correctly")
	}
	
	if len(signupRequest.Username) > 20 {
		return utils.BadRespone(ctx, "Username must be less than 20 characters!")
	}
	
	if len(signupRequest.Username) < 3 {
		return utils.BadRespone(ctx, "Username must be at least 3 characters long!")
	}
	
	client := utils.CreateClient()

	database := appwrite.NewDatabases(client)

	list, err := database.ListDocuments("mixtaper", "usernames", database.WithListDocumentsQueries(
		[]string{query.Equal("username", signupRequest.Username)},
	))

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to fetch used usernames", err)
	}
	
	if list.Total > 0 {
		return utils.BadRespone(ctx, "Username is taken")
	}
	
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to hash password", err)
	}
	
	users := appwrite.NewUsers(client)
	
	user, err := users.CreateBcryptUser(
		id.Unique(),
		signupRequest.Email,
		string(hashed_password),
		users.WithCreateBcryptUserName(signupRequest.Username),
	)
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to create user", err)
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
		return utils.ErrorResponse(ctx, "Failed to add username to database", err)
	}
	
	return utils.OkResponse(ctx, user.Name + " has signed up successfully!")
}
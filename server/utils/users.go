package utils

import (
	"os"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/query"

	"github.com/savsgio/atreugo/v11"
)

type User struct {
	Username string
	DisplayName string
	ID string
	Bio string
}

type Prefs struct {
	Bio string `json:"bio"`
}

type Username struct {
	Username string `json:"username"`
}

type UsernameList struct {
	*models.DocumentList
	Documents []Username `json:"documents"`
}

func CreateClientWithHeaders(ctx *atreugo.RequestCtx) (client.Client, bool) {
	sessionToken := ctx.Request.Header.Peek("Authorization")
	
	if sessionToken == nil { return client.Client{}, false }
	
	return appwrite.NewClient(
		appwrite.WithEndpoint(os.Getenv("APPWRITE_API_ENDPOINT")),
		appwrite.WithProject(os.Getenv("APPWRITE_PROJECT_ID")),
		appwrite.WithSession(string(sessionToken)),
	), true
}

func CreateClient() client.Client {
	return appwrite.NewClient(
		appwrite.WithEndpoint(os.Getenv("APPWRITE_API_ENDPOINT")),
		appwrite.WithProject(os.Getenv("APPWRITE_PROJECT_ID")),
		appwrite.WithKey(os.Getenv("APPWRITE_API_KEY")),
	)
}

func CheckUserExists(id string) bool {
	client := CreateClient()

	users := appwrite.NewUsers(client)
	
	user, err := users.Get(id)
	
	if err != nil || user == nil {
		return false
	}
	
	return true
}

func GetUserByUsername(client *client.Client, username string) (models.User, int) {
	database := appwrite.NewDatabases(*client)
	users := appwrite.NewUsers(*client)

	list, err := database.ListDocuments("mixtaper", "usernames", database.WithListDocumentsQueries(
		[]string{query.Equal("username", username)},
	))
	
	if err != nil {
		return models.User{}, 500
	}

	if list.Total == 0 {
		return models.User{}, 400
	}

	user, err := users.Get(list.Documents[0].Id)

	if err != nil {
		return models.User{}, 500
	}
	
	return *user, 200
}

func GetUsernameByID(client *client.Client, id string) (string, error) {
	database := appwrite.NewDatabases(*client)
	
	document, err := database.GetDocument("mixtaper", "usernames", id)

	if err != nil {
		return "", err
	}
	
	var username Username
	err = document.Decode(&username)
	
	if err != nil {
		return "", err
	}

	return username.Username, nil
}
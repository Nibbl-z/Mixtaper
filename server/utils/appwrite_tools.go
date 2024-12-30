package utils

import (
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/query"
)

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
		return models.User{},400
	}

	user, err := users.Get(list.Documents[0].Id)

	if err != nil {
		return models.User{}, 500
	}
	
	return *user, 200
}
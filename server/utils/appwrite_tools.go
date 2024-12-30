package utils

import (
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/query"
	"github.com/appwrite/sdk-for-go/users"
)

func GetUserByUsername(client *client.Client, username string) (models.User, error) {
	usersService := users.New(*client)
	
	usersList, err := usersService.List(
		usersService.WithListQueries([]string{query.Equal("name", username)}),
	)
	
	if err != nil {
		return models.User{}, err
	}
	
	return usersList.Users[0], nil
}
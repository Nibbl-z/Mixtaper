package utils

import (
	"github.com/appwrite/sdk-for-go/appwrite"
)

func CheckLevelExists(id string) bool {
	client := CreateClient()

	database := appwrite.NewDatabases(client)
	
	document, err := database.GetDocument(
		"mixtaper",
		"levels",
		string(id),
	)
	
	if err != nil || document == nil {
		return false
	}
	
	return true
}
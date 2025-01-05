package routes

import (
	"server/utils"
	
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/query"
	"github.com/savsgio/atreugo/v11"
)

func RecentLevels(ctx *atreugo.RequestCtx) error {
	client := utils.CreateClient()
	database := appwrite.NewDatabases(client)

	levels, err := database.ListDocuments(
		"mixtaper",
		"levels",
		database.WithListDocumentsQueries(
			[]string{
				query.OrderDesc("$createdAt"),
			},
		),
	)

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to fetch recent levels", err)
	}

	var levelList utils.LevelList
	levels.Decode(&levelList)

	return utils.OkResponse(ctx, levelList.Documents)
}
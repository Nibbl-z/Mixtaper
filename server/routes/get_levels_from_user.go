package routes

import (
	"server/utils"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/query"
	"github.com/savsgio/atreugo/v11"
)

func GetLevelsFromUser(ctx *atreugo.RequestCtx) error {
	id := ctx.Request.Body()
	
	if string(id) == "" {
		return utils.BadRespone(ctx, "User ID missing")
	}
	
	if !utils.CheckUserExists(string(id)) {
		return utils.BadRespone(ctx, "User doesn't exist")
	}
	
	client := utils.CreateClient()
	
	database := appwrite.NewDatabases(client)
	
	levels, err := database.ListDocuments(
		"mixtaper",
		"levels",
		database.WithListDocumentsQueries(
			[]string{query.Equal("uploader", string(id))},
		),
	)
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to fetch levels from database", err)
	}

	var levelList utils.LevelList
	levels.Decode(&levelList)
	
	return utils.OkResponse(ctx, levelList.Documents)
}
package routes

import (
	"server/utils"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/query"
	"github.com/savsgio/atreugo/v11"
)

type LevelList struct {
	*models.DocumentList
	Documents []Level `json:"documents"`
}

func GetLevelsFromUser(ctx *atreugo.RequestCtx) error {
	id := ctx.Request.Body()
	
	if string(id) == "" {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "User ID missing",
		}, 400)
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
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to fetch levels from database",
		}, 500)
	}

	var levelList LevelList
	levels.Decode(&levelList)
	
	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message": levelList.Documents,
	})
}
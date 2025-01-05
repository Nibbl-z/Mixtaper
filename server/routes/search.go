package routes

import (
	"fmt"
	"server/utils"
	
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/query"
	"github.com/savsgio/atreugo/v11"
)

func Search(ctx *atreugo.RequestCtx) error {
	searchQuery := string(ctx.QueryArgs().Peek("query"))
	
	if string(searchQuery) == "" {
		return utils.BadRespone(ctx, "Search query missing")
	}
	
	client := utils.CreateClient()
	database := appwrite.NewDatabases(client)
	fmt.Println(searchQuery)
	levels, err := database.ListDocuments(
		"mixtaper",
		"levels",
		database.WithListDocumentsQueries(
			[]string{
				query.Or([]string{
					query.Search("songName", string(searchQuery)),
					query.Search("songArtist", string(searchQuery)),
					query.Search("description", string(searchQuery)),
					query.Search("chartName", string(searchQuery)),
				}),
			},
		),
	)
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to search for levels", err)
	}

	var levelList utils.LevelList
	levels.Decode(&levelList)

	fmt.Println(levelList)

	return utils.OkResponse(ctx, levelList.Documents)
}
package routes

import ( 
	"encoding/json"

	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func EditLevel(ctx *atreugo.RequestCtx) error {
	var editRequest utils.EditPostData
	
	if err := json.Unmarshal(ctx.Request.Body(), &editRequest); err != nil {
		return utils.BadRespone(ctx, "Post data was not provided correctly")
	}

	if editRequest.SongName == "" || editRequest.SongArtist == "" || editRequest.ID == "" {
		return utils.BadRespone(ctx, "Post data is missing required fields")
	}

	if !utils.CheckLevelExists(editRequest.ID) {
		return utils.BadRespone(ctx, "Level doesn't exist")
	}
	
	client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return utils.UnauthorizedResponse(ctx)
	}
	
	database := appwrite.NewDatabases(client)
	
	document := map[string]string{
		"songName" : editRequest.SongName,
		"songArtist" : editRequest.SongArtist,
	}
	
	if editRequest.Description != "" {
		document["description"] = editRequest.Description
	}
	
	if editRequest.ChartName != "" {
		document["chartName"] = editRequest.ChartName
	}
	
	_, err := database.UpdateDocument(
		"mixtaper",
		"levels",
		editRequest.ID,
		database.WithUpdateDocumentData(document),
	)

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to update level", err)
	}
	
	return utils.OkResponse(ctx, "Updated level successfully!")
}
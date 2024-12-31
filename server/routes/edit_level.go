package routes

import ( 
	"encoding/json"

	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

type EditPostData struct {
	SongName string `json:"songName"`
	SongArtist string `json:"songArtist"`
	ChartName string `json:"chartName,omitempty"`
	Description string `json:"description,omitempty"`
	ID string `json:"id"`
}

func EditLevel(ctx *atreugo.RequestCtx) error {
	var editRequest EditPostData
	
	if err := json.Unmarshal(ctx.Request.Body(), &editRequest); err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Post data was not provided correctly",
		}, 400)
	}

	if editRequest.SongName == "" || editRequest.SongArtist == "" || editRequest.ID == "" {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Post data is missing required fields",
		}, 400)
	}
	
	client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Unauthorized",
		}, 401)
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
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to update level: " + err.Error(),
		}, 500)
	}
	
	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message" : "Updated level successfully!",
	}, 200)
}
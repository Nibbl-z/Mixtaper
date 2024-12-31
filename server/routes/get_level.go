package routes

import (
	"server/utils"
	
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/savsgio/atreugo/v11"
)

type Level struct {
	SongName string `json:"songName"`
	SongArtist string `json:"songArtist"`
	ChartName string `json:"chartName,omitempty"`
	Description string `json:"description,omitempty"`
	Uploader string `json:"uploader"`
}

func GetLevel(ctx *atreugo.RequestCtx) error {
	id := ctx.Request.Body()
	
	if string(id) == "" {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Level ID missing",
		}, 400)
	}

	client := utils.CreateClient()

	database := appwrite.NewDatabases(client)

	document, err := database.GetDocument(
		"mixtaper",
		"levels",
		string(id),
	)

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to get level",
		}, 500)
	}
	
	var level Level 
	err = document.Decode(&level)
	
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Could not parse level data:" + err.Error(),
		}, 500)
	}
	
	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message" : level,
	})
}
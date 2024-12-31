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
		return utils.BadRespone(ctx, "Level ID missing")
	}
	
	client := utils.CreateClient()
	
	database := appwrite.NewDatabases(client)

	document, err := database.GetDocument(
		"mixtaper",
		"levels",
		string(id),
	)

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to get level", err)
	}
	
	var level Level 
	err = document.Decode(&level)
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Could not parse level data", err)
	}

	return utils.OkResponse(ctx, level)
}
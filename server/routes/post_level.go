package routes

import (
	"encoding/json"
	"server/utils"
	
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/role"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/savsgio/atreugo/v11"
)

type PostData struct {
	SongName string `json:"songName"`
	SongArtist string `json:"songArtist"`
	ChartName string `json:"chartName,omitempty"`
	Description string `json:"description,omitempty"`
}

func PostLevel(ctx *atreugo.RequestCtx) error {
	var postRequest PostData
	
	if err := json.Unmarshal(ctx.Request.Body(), &postRequest); err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Post data was not provided correctly",
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
	account := appwrite.NewAccount(client)

	user, err := account.Get()

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to get user",
		}, 500)
	}
	
	if !user.EmailVerification {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Please verify your email before posting!",
		}, 400)
	}

	document := map[string]string{
		"songName" : postRequest.SongName,
		"songArtist" : postRequest.SongArtist,
		"uploader" : user.Id,
	}

	if postRequest.Description != "" {
		document["description"] = postRequest.Description
	}
	
	if postRequest.ChartName != "" {
		document["chartName"] = postRequest.ChartName
	}
	
	level, err := database.CreateDocument(
		"mixtaper",
		"levels",
		id.Unique(),
		document,
		database.WithCreateDocumentPermissions(
			[]string{
				permission.Read(role.Any()),
				permission.Write(role.User(user.Id, "")),
			},
		),
	)

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to post level metadata to database: " + err.Error(),
		}, 500)
	}

	return ctx.JSONResponse(map[string]interface{}{
		"successful": true,
		"message": "Posted level successfully!",
		"id" : level.Id,
	}, 200)
}
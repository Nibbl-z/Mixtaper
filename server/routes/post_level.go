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
		return utils.BadRespone(ctx, "Post data was not provided correctly")
	}
	
	client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return utils.UnauthorizedResponse(ctx)
	}
	
	database := appwrite.NewDatabases(client)
	account := appwrite.NewAccount(client)

	user, err := account.Get()

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to get user", err)
	}
	
	if !user.EmailVerification {
		return utils.BadRespone(ctx, "Please verify your email before posting!")
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
		return utils.ErrorResponse(ctx, "Failed to post level metadata to database", err)
	}

	return utils.OkPlusResponse(ctx, "Posted level successfully!", level.Id, "id")
}
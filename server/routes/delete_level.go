package routes

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func DeleteLevel(ctx *atreugo.RequestCtx) error {
	id := ctx.Request.Body()
	
	if string(id) == "" {
		return utils.BadRespone(ctx, "Level ID missing")
	}

	client, success := utils.CreateClientWithHeaders(ctx)
	
	if !success {
		return utils.UnauthorizedResponse(ctx)
	}
	
	database := appwrite.NewDatabases(client)
	storage := appwrite.NewStorage(client)

	_, err := database.DeleteDocument(
		"mixtaper",
		"levels",
		string(id),
	)
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to delete level", err)
	}

	_, err = storage.DeleteFile(
		"riq_files",
		string(id),
	)

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to delete RIQ", err)
	}

	return utils.OkResponse(ctx, "Deleted level successfully!")
}
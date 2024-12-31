package routes

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func DeleteLevel(ctx *atreugo.RequestCtx) error {
	id := ctx.Request.Body()
	
	if string(id) == "" {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Level ID missing",
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
	storage := appwrite.NewStorage(client)

	_, err := database.DeleteDocument(
		"mixtaper",
		"levels",
		string(id),
	)
	
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to delete level: " + err.Error(),
		}, 500)
	}

	_, err = storage.DeleteFile(
		"riq_files",
		string(id),
	)

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to delete RIQ: " + err.Error(),
		}, 500)
	}
	
	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message": "Deleted level successfully!",
	}, 200)
}
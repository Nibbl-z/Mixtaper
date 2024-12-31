package routes

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func DownloadRiq(ctx *atreugo.RequestCtx) error {
	id := ctx.Request.Body()
	
	if string(id) == "" {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Level ID missing",
		}, 400)
	}

	client := utils.CreateClient()
	
	storage := appwrite.NewStorage(client)

	fileBytes, err := storage.GetFileDownload(
		"riq_files",
		string(id),
	)

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to download .riq",
		}, 500)
	}

	return ctx.RawResponseBytes(*fileBytes, 200)
}

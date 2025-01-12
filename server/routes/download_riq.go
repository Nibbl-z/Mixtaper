package routes

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/appwrite/sdk-for-go/appwrite"
	"server/utils"
)

func DownloadRiq(ctx *atreugo.RequestCtx) error {
	id := string(ctx.QueryArgs().Peek("id"))
	
	if string(id) == "" {
		return utils.BadRespone(ctx, "Level ID missing")
	}

	client := utils.CreateClient()
	
	storage := appwrite.NewStorage(client)

	fileBytes, err := storage.GetFileDownload(
		"riq_files",
		string(id),
	)
	
	if err != nil {	
		return utils.ErrorResponse(ctx, "Failed to download riq", err)
	}
	
	return ctx.RawResponseBytes(*fileBytes, 200)
}
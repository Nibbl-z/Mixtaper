package routes

import (
	"io/fs"
	"os"
	
	"server/utils"

	"github.com/appwrite/sdk-for-go/account"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/file"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/role"
	"github.com/savsgio/atreugo/v11"
)

func UploadRiq(ctx *atreugo.RequestCtx) error {
	client, success := utils.CreateClientWithHeaders(ctx)

	if !success {
		return utils.UnauthorizedResponse(ctx)
	}

	id := ctx.Request.Header.Peek("ID")

	if id == nil {
		return utils.BadRespone(ctx, "Level ID is missing")
	}
	
	if !utils.CheckLevelExists(string(id)) {
		return utils.BadRespone(ctx, "Level doesn't exist")
	}

	if ctx.Request.Header.ContentLength() > 20 * 1000000 {
		return utils.BadRespone(ctx, "File must be less than 20MB.")
	}
	
	account_service := account.New(client)
	
	user, err := account_service.Get()
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to get user", err)
	}

	if !user.EmailVerification {
		return utils.BadRespone(ctx, "Please verify your email before posting!")
	}
	
	os.Mkdir("uploads", fs.FileMode(0644))
	
	body := ctx.Request.Body()
	
	path := "uploads/" + string(id) + ".riq"

	err = os.WriteFile(
		path, 
		body,
		fs.FileMode(0644),
	)

	

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to create temp upload file", err)
	}

	// Check RIQ file to be valid

	if err = utils.CheckRiq(path); err != nil {
		return utils.BadRespone(ctx, ".riq file is invalid!")
	}
	
	remixData, err := utils.GetRemixData(path)
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to decode .riq", err)
	}

	storage := appwrite.NewStorage(client)
	
	_, err = storage.GetFile("riq_files", string(id))
	
	if err == nil {
		return utils.BadRespone(ctx, ".riq file already exists")
	}
	
	database := appwrite.NewDatabases(client)
	
	document := map[string]interface{}{
		"bpm" : utils.GetBPM(remixData),
		"gamesUsed" : utils.GetGames(remixData),
	}

	_, err = database.UpdateDocument(
		"mixtaper",
		"levels",
		string(id),
		database.WithUpdateDocumentData(document),
	)

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to add .riq data to database", err)
	}
	
	permissions := []string{
		permission.Read(role.Any()),
		permission.Update(role.User(user.Id, "")),
		permission.Delete(role.User(user.Id, "")),
	}

	_, err = storage.CreateFile(
		"riq_files",
		string(id),
		file.NewInputFile(path, string(id) + ".riq"),
		storage.WithCreateFilePermissions(permissions),
	)
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to upload .riq to server", err)
	}

	fileRemoveErr := os.Remove(path)

	if fileRemoveErr != nil {
		return utils.ErrorResponse(ctx, "Failed to remove temp file from server", err)
	}
	
	return utils.OkResponse(ctx, "Uploaded .riq successfully!")
}
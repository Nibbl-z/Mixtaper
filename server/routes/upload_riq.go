package routes

import (
	"fmt"
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

	// TODO: check if level id is real
	
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
	fmt.Println(string(body))

	path := "uploads/" + user.Id + ".riq"

	err = os.WriteFile(
		path, 
		body,
		fs.FileMode(0644),
	)

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to create temp upload file", err)
	}
	
	permissions := []string{
		permission.Read(role.Any()),
		permission.Update(role.User(user.Id, "")),
		permission.Delete(role.User(user.Id, "")),
	}
	
	storage := appwrite.NewStorage(client)
	
	_, err = storage.CreateFile(
		"riq_files",
		string(id),
		file.NewInputFile(path, user.Id + ".riq"),
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
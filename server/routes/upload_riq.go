package routes

import (
	"fmt"
	"io/fs"
	"os"

	"server/utils"

	"github.com/appwrite/sdk-for-go/account"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/file"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/role"
	"github.com/savsgio/atreugo/v11"
)

func UploadRiq(ctx *atreugo.RequestCtx) error {
	
	client, success := utils.CreateClientWithHeaders(ctx)

	if !success {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Unauthorized",
		}, 401)
	}
	
	account_service := account.New(client)
	
	user, err := account_service.Get()
	
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to get user",
		}, 500)
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
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to create temp upload file",
		}, 500)
	}
	
	permissions := []string{
		permission.Read(role.Any()),
		permission.Update(role.User(user.Id, "")),
		permission.Delete(role.User(user.Id, "")),
	}
	
	storage := appwrite.NewStorage(client)
	
	_, err = storage.CreateFile(
		"riq_files",
		id.Unique(),
		file.NewInputFile(path, user.Id + ".riq"),
		storage.WithCreateFilePermissions(permissions),
	)

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to upload .riq to server",
		}, 500)
	}

	fileRemoveErr := os.Remove(path)

	if fileRemoveErr != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"successful" : false,
			"message": "Failed to remove temp file",
		}, 500)
	}

	return ctx.JSONResponse(map[string]interface{}{
		"successful" : true,
		"message": "Uploaded .riq successfully!",
	}, 200)
}
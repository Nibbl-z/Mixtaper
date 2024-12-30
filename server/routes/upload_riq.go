package routes

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/appwrite/sdk-for-go/account"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/file"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/role"
	"github.com/savsgio/atreugo/v11"
)

func UploadRiq(ctx *atreugo.RequestCtx) error {
	sessionToken := ctx.Request.Header.Peek("Authorization")
	
	if sessionToken == nil {
		return ctx.ErrorResponse(errors.New("unauthorized"), 401)
	}
	
	client := appwrite.NewClient(
		appwrite.WithEndpoint(os.Getenv("APPWRITE_API_ENDPOINT")),
		appwrite.WithProject(os.Getenv("APPWRITE_PROJECT_ID")),
		appwrite.WithSession(string(sessionToken)),
	)

	account_service := account.New(client)
	
	user, err := account_service.Get()
	
	if err != nil {
		return ctx.ErrorResponse(err, 500)
	}
	
	os.Mkdir("uploads", fs.FileMode(0644))
	
	body := ctx.Request.Body()
	fmt.Println(string(body))

	err = os.WriteFile(
		"uploads/" + user.Id + ".riq", 
		body,
		fs.FileMode(0644),
	)

	if err != nil {
		return ctx.ErrorResponse(err, 500)
	}
	
	permissions := []string{
		permission.Read(role.Any()),
		permission.Update(role.User(user.Id, "")),
		permission.Delete(role.User(user.Id, "")),
	}
	
	storage := appwrite.NewStorage(client)
	
	file, err := storage.CreateFile(
		"riq_files",
		id.Unique(),
		file.NewInputFile("uploads/" + user.Id + ".riq",  user.Id + ".riq"),
		storage.WithCreateFilePermissions(permissions),
	)
	
	if err != nil {
		return ctx.ErrorResponse(err, 500)
	}

	return ctx.TextResponse(file.Id)
}
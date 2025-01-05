package routes

import (
	"bytes"
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

func UploadCoverArt(ctx *atreugo.RequestCtx) error {
	client, success := utils.CreateClientWithHeaders(ctx)

	if (!success) {
		return utils.UnauthorizedResponse(ctx)
	}

	id := ctx.Request.Header.Peek("ID")

	if id == nil {
		return utils.BadRespone(ctx, "Level ID is missing")
	}
	
	if !utils.CheckLevelExists(string(id)) {
		return utils.BadRespone(ctx, "Level doesn't exist")
	}
	
	if ctx.Request.Header.ContentLength() > 20 * 500000 {
		return utils.BadRespone(ctx, "File must be less than 5MB.")
	}

	accountService := account.New(client)
	user, err := accountService.Get()
	
	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to get user", err)
	}

	if !user.EmailVerification {
		return utils.BadRespone(ctx, "Please verify your email before posting!")
	}
	
	os.Mkdir("covers", fs.FileMode(0644))

	body := ctx.Request.Body()
	
	
	
	isPng := bytes.Equal(body[:4], []byte{0x89, 0x50, 0x4E, 0x47})
	isJpeg := bytes.Equal(body[:3], []byte{0xFF, 0xD8, 0xFF})
	
	extension := ""
	
	if isPng {
		extension = ".png"
	} else if isJpeg {
		extension = ".jpg"
	}
	
	path := "covers/" + string(id) + extension

	// Check image file to be valid
	
	if !isPng && !isJpeg {
		return utils.BadRespone(ctx, "Invalid image. Only .png and .jpeg are supported.")
	}

	storage := appwrite.NewStorage(client)
	
	_, err = storage.GetFile("cover_art", string(id))
	
	if err == nil {
		return utils.BadRespone(ctx, "Cover art already exists")
	}
	
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

	_, err = storage.CreateFile(
		"cover_art",
		string(id),
		file.NewInputFile(path, string(id) + extension),
		storage.WithCreateFilePermissions(permissions),
	)

	if err != nil {
		fileRemoveErr := os.Remove(path)
	
		if fileRemoveErr != nil {
			return utils.ErrorResponse(ctx, "Failed to remove temp file from server", err)
		}
		return utils.ErrorResponse(ctx, "Failed to upload .riq to server", err)
	}
	
	fileRemoveErr := os.Remove(path)
	
	if fileRemoveErr != nil {
		return utils.ErrorResponse(ctx, "Failed to remove temp file from server", err)
	}

	return utils.OkResponse(ctx, "Uploaded cover art successfully!")
}
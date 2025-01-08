package routes

import (
	"bytes"
	"io/fs"
	"os"
	"server/utils"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/file"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/role"
	"github.com/savsgio/atreugo/v11"
)

func UploadPfp(ctx *atreugo.RequestCtx) error {
	client, success := utils.CreateClientWithHeaders(ctx)

	if !success {
		return utils.UnauthorizedResponse(ctx)
	}

	if ctx.Request.Header.ContentLength() > 5 * 1000000 {
		return utils.BadRespone(ctx, "File must be less than 5MB.")
	}
	
	accounts := appwrite.NewAccount(client)
	user, err := accounts.Get()

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to get user", err)
	}
	
	storage := appwrite.NewStorage(client)

	os.Mkdir("pfps", fs.FileMode(0644))

	body := ctx.Request.Body()
	
	isPng := bytes.Equal(body[:4], []byte{0x89, 0x50, 0x4E, 0x47})
	isJpeg := bytes.Equal(body[:3], []byte{0xFF, 0xD8, 0xFF})
	
	extension := ""
	
	if isPng {
		extension = ".png"
	} else if isJpeg {
		extension = ".jpg"
	}
	
	path := "pfps/" + string(user.Id) + extension

	if !isPng && !isJpeg {
		return utils.BadRespone(ctx, "Invalid image. Only .png and .jpeg are supported.")
	}

	err = os.WriteFile(path, body, fs.FileMode(0644))

	if err != nil {
		return utils.ErrorResponse(ctx, "Failed to create temp upload file", err)
	}

	f, err := storage.GetFile("profile_pictures", user.Id)
	
	if f != nil {
		_, err = storage.DeleteFile("profile_pictures", user.Id)

		if err != nil {
			fileRemoveErr := os.Remove(path)
		
			if fileRemoveErr != nil {
				return utils.ErrorResponse(ctx, "Failed to remove temp file from server", err)
			}

			return utils.ErrorResponse(ctx, "Failed to delete previous profile picture from server", err)
		}
	}

	permissions := []string{
		permission.Read(role.Any()),
		permission.Update(role.User(user.Id, "")),
		permission.Delete(role.User(user.Id, "")),
	}
	
	_, err = storage.CreateFile(
		"profile_pictures",
		user.Id,
		file.NewInputFile(path, user.Id + extension),
		storage.WithCreateFilePermissions(permissions),
	)
	
	if err != nil {
		fileRemoveErr := os.Remove(path)
	
		if fileRemoveErr != nil {
			return utils.ErrorResponse(ctx, "Failed to remove temp file from server", err)
		}
		return utils.ErrorResponse(ctx, "Failed to upload profile picture to server", err)
	}
	
	fileRemoveErr := os.Remove(path)
	
	if fileRemoveErr != nil {
		return utils.ErrorResponse(ctx, "Failed to remove temp file from server", err)
	}

	return utils.OkResponse(ctx, "Uploaded profile picture successfully!")
}
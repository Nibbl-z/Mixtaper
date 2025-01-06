package main

import (
	"server/routes"
	
	"github.com/joho/godotenv"
	"github.com/savsgio/atreugo/v11"
)

func main() {
    err := godotenv.Load()
	
	if err != nil {
		panic(err)
	}
	
	config := atreugo.Config{Addr: "localhost:2050", NoDefaultServerHeader: true, MaxRequestBodySize: 20 * 1024 * 1024}
    server := atreugo.New(config)
    
    server.OPTIONS("/*", func(ctx *atreugo.RequestCtx) error {
        ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:5173")
        ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, ID")
        ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
        return ctx.TextResponse("", 204)
    })
    
    server.UseBefore(func(ctx *atreugo.RequestCtx) error {
        ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:5173")
        ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, ID")
        ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
        
        return ctx.Next()
    })
    
    server.GET("/", func(ctx *atreugo.RequestCtx) error {
        return ctx.JSONResponse(atreugo.JSON{"Method": "GET"})
    })
    
    server.POST("/signup", routes.Signup)
    server.POST("/login", routes.Login)
    server.POST("/change_username", routes.ChangeUsername)
    
    server.POST("/send_email_verification", routes.SendEmailVerification)
    server.PUT("/verify_email", routes.VerifyEmail)

    server.GET("/get_level", routes.GetLevel)
    server.GET("/get_levels_from_user", routes.GetLevelsFromUser)
    server.GET("/download_riq", routes.DownloadRiq)
    server.GET("/search", routes.Search)
    server.GET("/get_user", routes.GetUser)
    
    server.POST("/upload_riq", routes.UploadRiq)
    server.POST("/upload_cover_art", routes.UploadCoverArt)
    server.POST("/post_level", routes.PostLevel)
    server.POST("/edit_level", routes.EditLevel)
    server.POST("/delete_level", routes.DeleteLevel)
    
    server.GET("/recent_levels", routes.RecentLevels)
    
    if err := server.ListenAndServe(); err != nil {
        panic(err)
    }
}
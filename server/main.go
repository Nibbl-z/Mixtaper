package main

import (
    "github.com/joho/godotenv"
    "github.com/savsgio/atreugo/v11"
    "server/routes"
)

func main() {
    err := godotenv.Load()
	
	if err != nil {
		panic(err)
	}
	
	config := atreugo.Config{Addr: "localhost:2050"}
    server := atreugo.New(config)
    
    server.GET("/", func(ctx *atreugo.RequestCtx) error {
        return ctx.TextResponse("Hai!")
    })
    
    server.POST("/signup", routes.Signup)
    server.POST("/login", routes.Login)
    server.POST("/upload_riq", routes.UploadRiq)
    server.POST("/post_level", routes.PostLevel)
    server.POST("/change_username", routes.ChangeUsername)

    server.GET("/get_level", routes.GetLevel)
    server.GET("/download_riq", routes.DownloadRiq)

    if err := server.ListenAndServe(); err != nil {
        panic(err)
    }
}
package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func successResp(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}

func ForbiddenResp(ctx *gin.Context, data error) {
	ctx.JSON(http.StatusForbidden, gin.H{
		"code": http.StatusForbidden,
		"msg":  data.Error(),
		"data": nil,
	})
}

func RunServer(port int, mode string) {
	if IS_PRO {
		gin.SetMode(gin.ReleaseMode)
		logFile, _ := os.OpenFile("go_everything.log", os.O_RDWR|os.O_CREATE, 0666)
		defer logFile.Close()
		gin.DefaultWriter = logFile
		if mode == "dev" {
			gin.SetMode(gin.DebugMode)
			logFile.Close()
			gin.DefaultWriter = os.Stdout
		}
	}
	r := gin.Default()
	if IS_PRO {
		r.Any("/", func(context *gin.Context) {
			context.Writer.Write(IndexHtml)
		})
		// get static dir children files
		staticFS, _ := fs.Sub(StaticAssets, "dist/assets")
		r.StaticFS("/assets", http.FS(staticFS))
	} else {
		r.Use(CorsMiddleware())
		r.Any("/", func(context *gin.Context) {
			context.JSON(200, "welcome to use go Searcher!")
		})
	}
	r.GET("/search", func(ctx *gin.Context) {
		key := ctx.Query("key")
		if key == "" {
			ForbiddenResp(ctx, errors.New("key is none"))
			return
		}
		key = DecodeB64String(key)
		store := ItemStore{}
		items, err := store.Find(key)
		if err != nil {
			ForbiddenResp(ctx, err)
			return
		}
		successResp(ctx, items)
	})
	r.GET("/indexDisk", func(ctx *gin.Context) {
		directory := ctx.Query("name")
		if directory == "" {
			ForbiddenResp(ctx, errors.New("directory is none"))
			return
		}
		directory = DecodeB64String(directory)
		//if err := IndexFile(directory); err != nil {
		//	ForbiddenResp(ctx, err)
		//	return
		//}
		go func() {
			IndexFile(directory)
		}()
		successResp(ctx, nil)
	})
	r.GET("/disk", func(ctx *gin.Context) {
		store := DiskStore{}
		disks, err := store.Find()
		if err != nil {
			ForbiddenResp(ctx, err)
			return
		}
		successResp(ctx, disks)
	})
	r.GET("/open", func(ctx *gin.Context) {
		openPath := ctx.Query("path")
		openPath = DecodeB64String(openPath)
		if openPath == "" {
			ForbiddenResp(ctx, errors.New("path is none"))
			return
		}
		_, err := os.Stat(openPath)
		if err != nil {
			ForbiddenResp(ctx, err)
			return
		}
		exec.Command("cmd.exe", "/c", "explorer", openPath).Start()
		successResp(ctx, nil)
	})
	log.Fatal(r.Run(fmt.Sprintf(":%d", port)))
}

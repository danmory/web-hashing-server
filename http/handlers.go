package http

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"gitlab.com/danmory/web-hashing-server/storages"
)

func HandleSave(storage *storages.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := ctx.Request.Body
		buf, err := io.ReadAll(body)
		if err != nil {
			ctx.JSON(404, "Error while reading body")
			fmt.Println(err)
			return
		}
		res, err := (*storage).Store(string(buf))
		if err != nil {
			ctx.JSON(500, err.Error())
			fmt.Println(err)
			return
		}
		ctx.JSON(200, res)
	}
}

func HandleRetrieve(storage *storages.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		target := ctx.Param("searched")
		res, err := (*storage).Find(target)
		if err != nil {
			ctx.JSON(500, err.Error())
			fmt.Println(err)
			return
		}
		ctx.JSON(200, res)
	}
}
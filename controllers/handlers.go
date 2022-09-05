package controllers

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/danmory/web-hashing-server/storages"
)

func HandleSave(storage *storages.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := ctx.Request.Body
		buf, err := io.ReadAll(body)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Error while reading body")
			log.Println(err)
			return
		}
		res, err := (*storage).Store(string(buf))
		if err != nil {
			ctx.String(http.StatusConflict, err.Error())
			log.Println(err)
			return
		}
		ctx.String(http.StatusOK, ctx.Request.Host+ctx.Request.URL.Path + res)
	}
}

func HandleRetrieve(storage *storages.Storage, pathParamName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		target := ctx.Param(pathParamName)
		res, err := (*storage).Find(target)
		if err != nil {
			ctx.String(http.StatusNotFound, err.Error())
			log.Println(err)
			return
		}
		ctx.String(http.StatusOK, res)
	}
}
package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"gitlab.com/danmory/web-hashing-server/http"
	"gitlab.com/danmory/web-hashing-server/storages"
)

type CLIParams struct {
	dFlag bool
}

func main() {
	params := parseCLI()
	storage := selectStorage(params.dFlag)
	r := gin.Default()
	r.POST("/", http.HandleSave(&storage))
	r.GET("/:searched", http.HandleRetrieve(&storage))
	r.Run()
}

func parseCLI() CLIParams {
	d := flag.Bool("d", false, "if -d present then data stored in DB")
	flag.Parse()
	return CLIParams{dFlag: *d}
}

func selectStorage(isDB bool) storages.Storage {
	var storageType storages.Type
	if isDB {
		storageType = storages.Database
	} else {
		storageType = storages.Memory
	}
	storage, err := storages.Get(storageType)
	if err != nil {
		panic(err)
	}
	return storage
}

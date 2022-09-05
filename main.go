package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/danmory/web-hashing-server/controllers"
	"gitlab.com/danmory/web-hashing-server/storages"
)

type CLIParams struct {
	dFlag bool
}

func main() {
	params := parseCLI()
	storage := selectStorage(params.dFlag)
	defer storage.Close()
	r := createServer(&storage)
	log.Fatal(r.Run())
}

func init() {
	godotenv.Load()
	log.SetOutput(gin.DefaultWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func createServer(storage *storages.Storage) *gin.Engine {
	r := gin.Default()
	r.POST("/", controllers.HandleSave(storage))
	r.GET("/:searched", controllers.HandleRetrieve(storage, "searched"))
	return r
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

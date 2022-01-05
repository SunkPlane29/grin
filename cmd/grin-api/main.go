package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SunkPlane29/grin/pkg/service/application"
	"github.com/SunkPlane29/grin/pkg/service/storage/memory"
	"github.com/SunkPlane29/grin/pkg/util"
)

var PORT = os.Getenv("PORT")

func main() {
	if PORT == "" {
		PORT = "8080"
	}
	log.Printf("Listening on :%s", PORT)

	userStorage := memory.NewUserStorage()
	postStorage := memory.NewPostStorage()
	grinStorage := application.NewGrinStorage(application.StorageOptions{
		UserStorage: userStorage,
		PostStorage: postStorage,
	})

	grinAPI := application.New(grinStorage)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", PORT),
		util.LoggerMiddleware("grin-api | ", util.RecoverMiddleware(grinAPI))),
	)
}

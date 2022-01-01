package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SunkPlane29/grin/pkg/application"
	"github.com/SunkPlane29/grin/pkg/storage/memory"
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
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), application.LoggerMiddleware(application.RecoverMiddleware(grinAPI))))
}

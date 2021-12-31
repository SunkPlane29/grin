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
	fmt.Printf("Listening on :%s\n", PORT)

	userStorage := memory.NewUserStorage()
	grinStorage := application.NewGrinStorage(userStorage)

	grinAPI := application.New(grinStorage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), application.LoggerMiddleware(application.RecoverMiddleware(grinAPI))))
}

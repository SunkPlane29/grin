package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SunkPlane29/grin/pkg/auth/token"
	"github.com/SunkPlane29/grin/pkg/service/application"
	"github.com/SunkPlane29/grin/pkg/service/storage/sqlite"
	"github.com/SunkPlane29/grin/pkg/util"
)

var PORT = os.Getenv("PORT")

func main() {
	if PORT == "" {
		PORT = "8080"
	}
	log.Printf("Listening on :%s", PORT)

	sqliteStorage, err := sqlite.NewScratch(context.Background(), "./grin.db")
	if err != nil {
		log.Fatal(err)
	}
	defer sqliteStorage.Close()

	// userStorage := memory.NewUserStorage()
	// postStorage := memory.NewPostStorage()

	grinStorage := application.NewGrinStorage(application.StorageOptions{
		UserStorage: sqliteStorage,
		PostStorage: sqliteStorage,
		// UserStorage: userStorage,
		// PostStorage: postStorage,
	})

	keys, err := token.NewKeysFromCertFiles("cert/id_rsa.pub", "cert/id_rsa")
	if err != nil {
		log.Fatal(err)
	}

	grinAPI := application.New(grinStorage, token.NewJWT(keys.PubKey, keys.PrivateKey))
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", PORT),
		util.LoggerMiddleware("grin-api | ", util.RecoverMiddleware(grinAPI))),
	)
}

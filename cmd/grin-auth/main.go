package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SunkPlane29/grin/pkg/auth/application"
	"github.com/SunkPlane29/grin/pkg/auth/core"
	"github.com/SunkPlane29/grin/pkg/auth/storage/sqlite"
	"github.com/SunkPlane29/grin/pkg/auth/token"
	"github.com/SunkPlane29/grin/pkg/util"
)

func newDB() core.AuthenticationStorage {
	if os.Getenv("SCRATCHDB") == "true" {
		s, err := sqlite.NewScratch(context.Background(), "./grin-auth.db")
		if err != nil {
			log.Fatal(err)
		}

		return s
	} else {
		s, err := sqlite.New("./grin-auth.db")
		if err != nil {
			log.Fatal(err)
		}

		return s
	}
}

func main() {
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "9090"
	}

	// s := memory.NewAuthorizationStorage()

	s := newDB()

	keys, err := token.NewKeysFromCertFiles("cert/id_rsa.pub", "cert/id_rsa")
	if err != nil {
		log.Fatal(err)
	}

	serv := core.NewAuthorizationService(s, keys)
	_ = serv

	authServer := application.NewAuthServer(serv)

	log.Printf("Listening on :%s\n", PORT)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", PORT),
		util.RecoverMiddleware(util.LoggerMiddleware("grin-auth | ", authServer))),
	)
}

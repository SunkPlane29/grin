package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SunkPlane29/grin/pkg/application"
	"github.com/SunkPlane29/grin/pkg/storage/memory"
)

const PORT = "8080"

func main() {
	fmt.Println("Hello, World!")

	userStorage := memory.NewUserStorage()
	grinStorage := application.NewGrinStorage(userStorage)

	grinAPI := application.New(grinStorage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), grinAPI))
}

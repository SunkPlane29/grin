package main

import (
	"fmt"

	"github.com/SunkPlane29/solitude/pkg/post"
	"github.com/SunkPlane29/solitude/pkg/user"
)

func main() {
	fmt.Println("Hello, World!")

	_ = user.NewService(nil)
	_ = post.NewService(nil)
}

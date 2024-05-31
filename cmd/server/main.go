package main

import (
	"github.com/tuanbui-n9/project-4/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8080")
}

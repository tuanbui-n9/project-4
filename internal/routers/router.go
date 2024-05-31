package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/tuanbui-n9/project-4/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping")
		v1.GET("/users/:id", c.NewUserController().GetUserById)
	}

	return r
}

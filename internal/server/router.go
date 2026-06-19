package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/maneul0498-netizen/unicomer_tech_challenge/docs"
	"github.com/maneul0498-netizen/unicomer_tech_challenge/internal/interfaces/http/handler"
)

type Router struct {
	Eng     *gin.Engine
	Handler *handler.Handler
}

func (r *Router) InitRouters() {
	r.initV1Routers()
}

func (r *Router) initV1Routers() {

	h := r.Eng.Group("/api/v1/holidays")
	{
		h.GET("", r.Handler.Get)
	}

	r.Eng.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

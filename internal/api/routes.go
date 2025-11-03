package api

import (
	"fmt"
	"net/http"
	"payme/internal/utils"

	_ "payme/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

type GinServer struct {
	eng *gin.Engine
	cache *utils.RedisCache
}

func NewGinServer() *GinServer {
	// cache, err := utils.NewRedisCache(time.Second * 5)
	// if err != nil {
	// 	log.Fatalf("failed to connect to redis due to %v\n", err.Error())
	// }
	return &GinServer{eng: gin.Default()}
}

func (s* GinServer)NewRouteSubGroup(version, name string) *gin.RouterGroup {
	routeGroup :=  s.eng.Group(fmt.Sprintf("/api/%s/%s", version, name))
	return routeGroup
}

func (s *GinServer) Run(addr string) error {
	s.eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.eng.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "welcome to new york city!",
			"docs": "you can access the documentation at `/docs/swagger.html`",
			"health": "go to `/health` to check the health status",
		})
	})

	return s.eng.Run(addr)
}

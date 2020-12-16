package main

import (
	"github.com/vcnt72/golang-learn/repository"
	"net/http"
	"github.com/vcnt72/golang-learn/middleware"
	"github.com/vcnt72/golang-learn/controller"
	"github.com/vcnt72/golang-learn/service"
	"github.com/gin-gonic/gin"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService service.VideoService = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
	jwtService service.JWTService = service.NewJwtService()
	loginService service.LoginService = service.NewLoginService()
	loginController controller.LoginController = controller.NewLoginController(loginService,jwtService)
)

func main() {
	server := gin.New()


	server.Use(gin.Recovery(),middleware.Logger())

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, nil)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"token" : token,
			})
		}
	})

	apiRoutes := server.Group("/api", middleware.AuthorizeJwt()) 
	{
		

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
	
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "success",
				})
			}
		})
	}
	

	server.Run(":8080")
}
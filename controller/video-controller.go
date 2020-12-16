package controller

import (
	"log"
	"github.com/vcnt72/golang-learn/entity"
	"github.com/gin-gonic/gin"
	"github.com/vcnt72/golang-learn/service"
)


type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}


func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	log.Println(err)

	if err != nil {
		return err
	}

	log.Println(video.Title)
	c.service.Save(video)
	return nil
}
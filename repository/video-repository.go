package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/vcnt72/golang-learn/entity"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	dsn := "host=localhost user=postgres password=mateup123 dbname=golang-learn port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database {
		connection: db,
	}
}


func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)	
	return videos
}


func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}
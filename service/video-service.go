package service
import (
	"github.com/vcnt72/golang-learn/repository"
	"github.com/vcnt72/golang-learn/entity"
)


type VideoService interface {
	Save(entity.Video ) error
	Update(entity.Video) error
	Delete(entity.Video) error
	FindAll() []entity.Video
}

type videoService struct {
	repository repository.VideoRepository
}

func New(videoRepository repository.VideoRepository) VideoService {
	return &videoService{
		repository: videoRepository,
	}
}


func (service *videoService) Update(video entity.Video) error {
	service.repository.Update(video)
	return nil
}

func (service *videoService) Delete(video entity.Video) error {
	service.repository.Delete(video)
	return nil
}


func (service *videoService) Save(video entity.Video) error {
	service.repository.Save(video)
	return nil
}

func (service *videoService) FindAll() []entity.Video {
	return service.repository.FindAll()
}
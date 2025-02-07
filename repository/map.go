package repository

import (
	"github.com/mahdi-cpp/api-go-emqx/cache"
	"github.com/mahdi-cpp/api-go-emqx/model"
	"github.com/mahdi-cpp/api-go-emqx/utils"
)

var mapDTO MapDTO

type MapDTO struct {
	Caption string            `json:"name"`
	Maps    []Map             `json:"maps"`
	Users   []model.PhotoBase `json:"users"`
}

type Map struct {
	Name  string          `json:"name"`
	Photo model.PhotoBase `json:"photo"`
}

func GetMaps(folder string) MapDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto MapDTO
	var avatarUrl = []string{"b2", "b5", "b19"}
	var index = 0
	var nameIndex = 0

	for i := 0; i < 3; i++ {
		var user = model.PhotoBase{}
		user.Name = avatarUrl[i]
		user.Key = -1
		user.ThumbSize = 270
		user.Circle = true
		user.PaintWidth = dp(100)
		user.PaintHeight = dp(100)
		dto.Users = append(dto.Users, user)
	}

	for i := 0; i < count; i++ {
		var mapData Map
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		mapData.Name = utils.MovieNames[nameIndex]

		mapData.Photo = photos[index]
		mapData.Photo.Key = -1
		mapData.Photo.ThumbSize = 270
		mapData.Photo.PaintWidth = dp(70)
		mapData.Photo.PaintHeight = dp(120)

		dto.Maps = append(dto.Maps, mapData)
		nameIndex++
		index++
	}

	index = 0

	return dto
}

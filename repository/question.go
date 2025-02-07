package repository

import (
	"github.com/mahdi-cpp/api-go-emqx/cache"
	"github.com/mahdi-cpp/api-go-emqx/model"
	"github.com/mahdi-cpp/api-go-emqx/utils"
)

var questionVoiceDTO QuestionVoiceDTO

type QuestionVoiceDTO struct {
	Caption        string          `json:"name"`
	QuestionVoices []QuestionVoice `json:"questionVoices"`
}

type QuestionVoice struct {
	Name  string          `json:"name"`
	Photo model.PhotoBase `json:"photo"`
}

func GetQuestionVoices(folder string) QuestionVoiceDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto QuestionVoiceDTO

	//if count > 50 {
	//	count = 50
	//}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var voice QuestionVoice
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		voice.Name = utils.MovieNames[nameIndex]

		voice.Photo = photos[index]
		voice.Photo.Key = -1
		voice.Photo.ThumbSize = 270
		voice.Photo.Crop = 2
		voice.Photo.Round = int(dp(10))
		voice.Photo.PaintWidth = dp(95)
		voice.Photo.PaintHeight = dp(95)

		dto.QuestionVoices = append(dto.QuestionVoices, voice)
		nameIndex++
		index++
	}

	index = 0

	return dto
}

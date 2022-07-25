package clients

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Nexters/myply/domain/entity"
	"github.com/Nexters/myply/infrastructure/configs"
)

type YoutubeV3SearchResponse struct {
	Items [](struct {
		Snippet struct {
			Title string `json:"title"`
			Thumb struct {
				Default struct {
					Url string `json:"url"`
				} `json:"default"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	}) `json:"items"`
}

type youtubeV3 struct {
	config configs.Config
}

func NewYoutubeApiV3(c *configs.Config) Youtube {
	return &youtubeV3{config: *c}
}

func (ya *youtubeV3) Search(query string) (*entity.Musics, error) {
	// TODO: research elegant http module
	url := "https://www.googleapis.com/youtube/v3/search" +
		"?key=" + ya.config.YoutubeApiKey +
		"&part=snippet" +
		"&q=" + query +
		"&videoCategoryId=10" +
		"&type=video" +
		"&regionCode=kr" +
		"&maxResults=50"

	res, httpError := http.Get(url)

	if httpError != nil {
		return nil, httpError
	}

	jsonBytes, _ := ioutil.ReadAll(res.Body)
	defer func() {
		res.Body.Close()
	}()

	var response YoutubeV3SearchResponse
	parseErr := json.Unmarshal(jsonBytes, &response)

	if parseErr != nil {
		return nil, parseErr
	}

	return ya.searchResponseToMusics(&response), nil
}

func (ya *youtubeV3) searchResponseToMusics(res *YoutubeV3SearchResponse) *entity.Musics {
	count := len(res.Items)
	musics := make(entity.Musics, count)

	for i := range res.Items {
		musics[i] = entity.Music{
			ThumbURL: res.Items[i].Snippet.Thumb.Default.Url,
			Title:    res.Items[i].Snippet.Title,
		}
	}

	return &musics
}

package musics

import "fmt"

type Music struct {
	YoutubeVideoID string
	ThumbnailURL   string
	Title          string
	YoutubeTags    []string
	// TODO: IsMemoed bool,
}

type Musics []Music

func NewMusic(youtubeVideoID, thumbnailURL, title string, tags []string) *Music {
	return &Music{
		YoutubeVideoID: youtubeVideoID,
		ThumbnailURL:   thumbnailURL,
		Title:          title,
		YoutubeTags:    tags,
	}
}

func (m *Music) IsEqual(other *Music) bool {
	return other.YoutubeVideoID == m.YoutubeVideoID
}

func (m *Music) DeepLink() string {
	return fmt.Sprintf("vnd.youtube://%s", m.YoutubeVideoID)
}

package entity

type Music struct {
	ThumbURL string `json:"thumbURL"`
	Title    string `json:"title"`
}

type Musics []Music
